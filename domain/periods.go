package domain

import (
	"fmt"
	"sort"
	"time"
)

type Period struct {
	PeriodType string
	Duration   int
}

// The list of periods that we want to calculate
var (
	PeriodsList       []*Period
	AssociatedPeriods = make(map[string]*Period)
	PeriodsMap        = map[string]*Period{}
)

func init() {
	initPeriodsList()
	initLookupPeriods()
	initPeriodsMap()
}

func initPeriodsList() {
	PeriodsList = make([]*Period, 0)
	PeriodsList = append(PeriodsList, &Period{PeriodType: "minute", Duration: 1})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "minute", Duration: 3})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "minute", Duration: 5})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "minute", Duration: 15})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "minute", Duration: 30})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "hour", Duration: 1})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "hour", Duration: 3})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "hour", Duration: 6})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "hour", Duration: 12})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "day", Duration: 1})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "day", Duration: 3})
	PeriodsList = append(PeriodsList, &Period{PeriodType: "week", Duration: 1})
}

func initPeriodsMap() {
	// String to period lookup map (for faster lookup than recalculating every time)
	for _, p := range PeriodsList {
		PeriodsMap[p.ToString()] = p
	}
}

// Lookup periods is a map of the period types and their corresponding period duration in minutes for use in the AED calculation
// The map is calculated to contain the largest possible modulus value on which to base the next calculation
// By storing the values in minutes in this map, they do not need to be parsed every time the AED is calculated
func initLookupPeriods() {
	type MinutePeriod struct {
		duration int
		key      string
		period   *Period
	}
	// Initialize the period lookup map:
	// Translate the values all to Duration in minute:
	convertedPeriods := make([]*MinutePeriod, 0)
	for _, period := range PeriodsList {
		convertedPeriods = append(convertedPeriods,
			&MinutePeriod{duration: period.ToMinute().Duration, key: period.ToString(),
				period: period,
			})
	}
	// Order the converted periods by duration descending: (More efficient to find the next best matching modulus)
	sort.Slice(convertedPeriods, func(i, j int) bool {
		return convertedPeriods[i].duration > convertedPeriods[j].duration
	})
	// Setup a map with the now new period duration as the period to use as lookup key
	// When looking this up we only care about the us4ed key for the data, the actual durations do not matter any longer (they have been used to calculate this map)
	for _, period := range convertedPeriods {
		// Find the next best matching modulus
		for _, nextPeriod := range convertedPeriods {
			if period.duration%nextPeriod.duration == 0 && period.duration != nextPeriod.duration {
				AssociatedPeriods[period.key] = nextPeriod.period
				break
			}
		}
	}
}

func (s *Period) ToString() string {
	return fmt.Sprintf("%d%s", s.Duration, s.PeriodType)
}

// Modulus calculations require a stable duration. Minute is the choosen base unit.
// This function translates the period to a period in minutes
func (s *Period) ToMinute() *Period {
	// Dereference the pointer to get a copy of the object (Do not act on the pointer else all periods change and lookups depending on the structure will fail)
	d := *s
	switch s.PeriodType {
	case "hour":
		d.PeriodType = "minute"
		d.Duration = s.Duration * 60
	case "day":
		d.PeriodType = "minute"
		d.Duration = s.Duration * 60 * 24
	case "week":
		d.PeriodType = "minute"
		d.Duration = s.Duration * 60 * 24 * 7
	}
	return &d
}

func (s *Period) offset() int64 {
	if s.PeriodType == "week" {
		return 4 * int64(time.Minute) * 24 * 60 // Start the week on mondays
	}
	return 0
}

// Returns the key timestamp for any given period by calculating the minute minus the modulus for the given duration
func (s *Period) ToAEDKeyTimestamp(timestamp int64) int64 {
	t := s.ToMinute()
	ts := timestamp - timestamp%(int64(t.Duration)*int64(time.Minute))
	ts = ts + s.offset()
	if ts > timestamp {
		// Correct the week calculation if the period is beyond the calculated value
		ts = ts - int64(time.Minute*24*60*7)
	}
	return ts
}

// Returns the key timestamp for any given period by calculating the minute minus the modulus for the given duration
func (s *Period) ToAEDKeyTimestampFrom(timestamp int64) int64 {
	return s.ToAEDKeyTimestamp(timestamp)
}

// Returns the end of the timestamp window for the given period and timestamp
func (s *Period) ToAEDKeyTimestampTo(timestamp int64) int64 {
	t := s.ToMinute()
	ts := s.ToAEDKeyTimestamp(timestamp)
	// Set to end of period:
	ts = ts + int64(t.Duration)*int64(time.Minute)
	return ts
}

// Translates a string (1minute, 1hour, 3day, 1week, etc) to a period
func StringToPeriod(period string) *Period {
	return PeriodsMap[period]
}
