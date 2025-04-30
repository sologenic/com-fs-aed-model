package domain

import (
	"fmt"
	"sort"
	"time"

	aedgrpc "github.com/sologenic/com-fs-aed-model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Time unit constants
const (
	secondsInMinute = 60
	secondsInHour   = 60 * 60          // 3,600
	secondsInDay    = 24 * 60 * 60     // 86,400
	secondsInWeek   = 7 * 24 * 60 * 60 // 604,800
)

// The list of periods that we want to calculate
var (
	PeriodsList       []*aedgrpc.Period
	AssociatedPeriods = make(map[string]*aedgrpc.Period)
	PeriodsMap        = map[string]*aedgrpc.Period{}
)

func init() {
	initPeriodsList()
	initLookupPeriods()
	initPeriodsMap()
}

func initPeriodsList() {
	PeriodsList = make([]*aedgrpc.Period, 0)
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 3})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 5})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 15})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 30})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 3})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 6})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 12})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 3})
	PeriodsList = append(PeriodsList, &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1})
}

func initPeriodsMap() {
	// String to period lookup map (for faster lookup than recalculating every time)
	for _, p := range PeriodsList {
		PeriodsMap[ToString(p)] = p
	}
}

// Lookup periods is a map of the period types and their corresponding period duration in minutes for use in the AED calculation
// The map is calculated to contain the largest possible modulus value on which to base the next calculation
// By storing the values in minutes in this map, they do not need to be parsed every time the AED is calculated
func initLookupPeriods() {
	type MinutePeriod struct {
		duration int
		key      string
		period   *aedgrpc.Period
	}
	// Initialize the period lookup map:
	// Translate the values all to Duration in minute:
	convertedPeriods := make([]*MinutePeriod, 0)
	for _, period := range PeriodsList {
		convertedPeriods = append(convertedPeriods,
			&MinutePeriod{
				duration: int(ToMinute(period).Duration),
				key:      ToString(period),
				period:   period,
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

func ToString(p *aedgrpc.Period) string {
	return fmt.Sprintf("%d_%s", p.Duration, p.Type)
}

// Modulus calculations require a stable duration. Minute is the choosen base unit.
// This function translates the period to a period in minutes
func ToMinute(p *aedgrpc.Period) *aedgrpc.Period {
	// Dereference the pointer to get a copy of the object (Do not act on the pointer else all periods change and lookups depending on the structure will fail)
	pt := aedgrpc.PeriodType_PERIOD_TYPE_MINUTE
	d := p.Duration
	switch p.Type {
	case aedgrpc.PeriodType_PERIOD_TYPE_HOUR:
		d = p.Duration * 60
	case aedgrpc.PeriodType_PERIOD_TYPE_DAY:
		d = p.Duration * 60 * 24
	case aedgrpc.PeriodType_PERIOD_TYPE_WEEK:
		d = p.Duration * 60 * 24 * 7
	}
	return &aedgrpc.Period{
		Type:     pt,
		Duration: d,
	}
}

func Offset(p *aedgrpc.Period) int64 {
	if p.Type == aedgrpc.PeriodType_PERIOD_TYPE_WEEK {
		return 4 * secondsInDay // Start the week on mondays
	}
	return 0
}

// Returns the key timestamp for any given period by calculating the minute minus the modulus for the given duration
func ToAEDKeyTimestamp(p *aedgrpc.Period, timestamp int64) int64 {
	t := ToMinute(p)
	ts := timestamp - timestamp%(int64(t.Duration)*secondsInMinute)
	ts = ts + Offset(p)
	if ts > timestamp {
		// Correct the week calculation if the period is beyond the calculated value
		ts = ts - secondsInWeek
	}
	return ts
}

func ToAEDKeyTimestamppb(p *aedgrpc.Period, timestamp *timestamppb.Timestamp) *timestamppb.Timestamp {
	t := timestamp.AsTime().Unix()
	ts := ToAEDKeyTimestamp(p, t)
	return timestamppb.New(time.Unix(0, ts))
}

// Returns the key timestamp for any given period by calculating the minute minus the modulus for the given duration
func ToAEDKeyTimestampFrom(p *aedgrpc.Period, timestamp int64) int64 {
	return ToAEDKeyTimestamp(p, timestamp)
}

// Returns the end of the timestamp window for the given period and timestamp
func ToAEDKeyTimestampTo(p *aedgrpc.Period, timestamp int64) int64 {
	t := ToMinute(p)
	ts := ToAEDKeyTimestamp(p, timestamp)
	// Set to end of period:
	ts = ts + int64(t.Duration)*secondsInMinute
	return ts
}

// Translates a string (1minute, 1hour, 3day, 1week, etc) to a period
func StringToPeriod(period string) *aedgrpc.Period {
	return PeriodsMap[period]
}
