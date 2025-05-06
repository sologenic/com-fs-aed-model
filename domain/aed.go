package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	aedgrpc "github.com/sologenic/com-fs-aed-model"
)

type AEDPointResponse [6]interface{}

var allowedValues = map[string]bool{
	"1m":  true,
	"3m":  true,
	"5m":  true,
	"15m": true,
	"30m": true,
	"1h":  true,
	"3h":  true,
	"6h":  true,
	"12h": true,
	"1d":  true,
	"3d":  true,
	"1w":  true,
}

var (
	periodRegex = regexp.MustCompile(`^(\d+)([a-zA-Z]+)$`)
	// Errors
	ErrIncorrectRequestParm = errors.New("incorrect request parameter")
)

// Input: ["1m","3m","5m","15m","30m","1h","3h","6h","12h","1d","3d","1w"]
// Or invalid input
// Output:
// aedgrpc.Period
func HttpPeriodToPeriod(value string) (*aedgrpc.Period, error) {
	if !allowedValues[value] {
		return nil, ErrIncorrectRequestParm
	}
	period := &aedgrpc.Period{}
	matches := periodRegex.FindStringSubmatch(value)
	if len(matches) != 3 {
		return period, fmt.Errorf("invalid value: %s", value)
	}

	duration, _ := strconv.Atoi(matches[1])
	period.Duration = int32(duration)
	period.Type = mapStringToPeriodType(matches[2])
	return period, nil
}

func mapStringToPeriodType(s string) aedgrpc.PeriodType {
	switch s {
	case "m":
		return aedgrpc.PeriodType_PERIOD_TYPE_MINUTE
	case "h":
		return aedgrpc.PeriodType_PERIOD_TYPE_HOUR
	case "d":
		return aedgrpc.PeriodType_PERIOD_TYPE_DAY
	case "w":
		return aedgrpc.PeriodType_PERIOD_TYPE_WEEK
	default:
		return aedgrpc.PeriodType_PERIOD_TYPE_DO_NOT_USE
	}
}

// Outliers are correct and can occur due to ledger behaviour: A very small trade can occur at a very high price.
// This disturbs the graph and does not represent the reality of the pricing which occurs.
// Since such a transaction can occur as a single transaction in a single minute (so no other data to evaluate it against to be able to identify it as an outlier)
// the choice is made to smooth the outliers on retrieval.
// The smoothing is done by replacing the outlier with the average of the previous and next value if the current value deviates outside the norm and is not correctable within the current time interval
// Correction within the current time interval is done by replacing the outlier with another aed value from within the time interval.
// Outlier detection to see if the current minute with all respectable values are within range of the next interval, and for the last value to see if it is within range of the 1 to last interval.
func SmoothOutliers(series []*aedgrpc.AED, index int) *aedgrpc.AED {
	data := series[index]

	// We replace inline the values which are deviating.
	// Assumption for the very simple scenario is that the values provided will have a very small divider used, so will show a much too large value for the high price.
	// If the high price deviations outside of the norm, we will inspect open/close too and replace accordingly.
	// The second check this code does is to see if the previous or next value has a high price which is without the range of reasonable values.
	// If uses the previousaed by default except for the first value, for that it will use the next value.
	// Since this detection only works if there is more than 1 trade in the base data for calculate the AED, a single trade in a single aed will not show up and will not be corrected
	open := ParseFieldValue[float64](data, aedgrpc.Field_OPEN)
	high := ParseFieldValue[float64](data, aedgrpc.Field_HIGH)
	low := ParseFieldValue[float64](data, aedgrpc.Field_LOW)
	close := ParseFieldValue[float64](data, aedgrpc.Field_CLOSE)
	if high/low > 3 {
		// Create new AED with same metadata
		o := &aedgrpc.AED{
			OrganizationID: data.OrganizationID,
			Symbol:         data.Symbol,
			Timestamp:      data.Timestamp,
			MetaData:       data.MetaData,
			Series:         data.Series,
			Value:          make([]*aedgrpc.Value, 0),
		}

		// Set the low value
		o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_LOW, low))

		newHigh := low
		// Handle different outlier cases
		switch {
		case high/open > 3:
			// Open is correct
			newHigh = open
			o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_OPEN, open))
			fallthrough
		case high/close > 3:
			// Close is correct
			o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_CLOSE, close))
			if close > newHigh {
				newHigh = close
			}
			fallthrough
		case high/open < 2:
			// Open is incorrect (the high was the open)
			o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_OPEN, newHigh))
			fallthrough
		case high/close < 2:
			// Close is incorrect (the high was the close)
			o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_CLOSE, newHigh))
		}

		o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_HIGH, newHigh))

		// Copy other fields that might be present
		for _, v := range data.Value {
			if v.Field != aedgrpc.Field_OPEN &&
				v.Field != aedgrpc.Field_HIGH &&
				v.Field != aedgrpc.Field_LOW &&
				v.Field != aedgrpc.Field_CLOSE {
				o.Value = append(o.Value, v)
			}
		}
		return o
	}

	// Check for major deviations in the data
	// Missing scenario is when the value being retrieved is the current minute and that would have a graphing deviation.
	if len(series) > 1 {
		lookup := index - 1
		if index == 0 {
			lookup = 1
		}
		lookupLow := ParseFieldValue[float64](series[lookup], aedgrpc.Field_LOW)
		if high/lookupLow > 10 {
			// Severe deviation found: replace with reference values
			lookupClose := ParseFieldValue[float64](series[lookup], aedgrpc.Field_CLOSE)
			o := &aedgrpc.AED{
				OrganizationID: data.OrganizationID,
				Symbol:         data.Symbol,
				Timestamp:      series[lookup].Timestamp,
				MetaData:       data.MetaData,
				Series:         data.Series,
				Value:          make([]*aedgrpc.Value, 0),
			}

			// Use lookupClose for all values
			o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_OPEN, lookupClose))
			o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_HIGH, lookupClose))
			o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_LOW, lookupClose))
			o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_CLOSE, lookupClose))
			o.Value = append(o.Value, CreateFieldValue(aedgrpc.Field_VOLUME, 0.0))

			return o
		}
	}

	return data
}

func GetFloatValue(aed *aedgrpc.AED, field aedgrpc.Field) float64 {
	for _, v := range aed.Value {
		if v.Field == field && v.Float64Val != nil {
			return *v.Float64Val
		}
	}
	return 0.0
}

func GetIntValue(aed *aedgrpc.AED, field aedgrpc.Field) int64 {
	for _, v := range aed.Value {
		if v.Field == field && v.Int64Val != nil {
			return *v.Int64Val
		}
	}
	return 0
}

func GetStringValue(aed *aedgrpc.AED, field aedgrpc.Field) string {
	for _, v := range aed.Value {
		if v.Field == field && v.StringVal != nil {
			return *v.StringVal
		}
	}
	return ""
}

func SetFloatValue(aed *aedgrpc.AED, field aedgrpc.Field, value float64) {
	for _, v := range aed.Value {
		if v.Field == field {
			v.Float64Val = &value
			v.Int64Val = nil
			v.StringVal = nil
			return
		}
	}
	aed.Value = append(aed.Value, &aedgrpc.Value{
		Field:      field,
		Float64Val: &value,
	})
}

func SetIntValue(aed *aedgrpc.AED, field aedgrpc.Field, value int64) {
	for _, v := range aed.Value {
		if v.Field == field {
			v.Int64Val = &value
			v.Float64Val = nil
			v.StringVal = nil
			return
		}
	}
	aed.Value = append(aed.Value, &aedgrpc.Value{
		Field:    field,
		Int64Val: &value,
	})
}

func SetStringValue(aed *aedgrpc.AED, field aedgrpc.Field, value string) {
	for _, v := range aed.Value {
		if v.Field == field {
			v.StringVal = &value
			v.Float64Val = nil
			v.Int64Val = nil
			return
		}
	}
	aed.Value = append(aed.Value, &aedgrpc.Value{
		Field:     field,
		StringVal: &value,
	})
}
