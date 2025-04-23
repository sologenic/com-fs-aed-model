package domain

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	aedgrpc "github.com/sologenic/com-fs-aed-model"
)

// Test the periods lookup
func Test_AssociatedPeriods(t *testing.T) {
	// Test the lookup for 1 minute:
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	cmp.Equal(AssociatedPeriods[ToString(period)], period)
	period3m := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 3}
	cmp.Equal(AssociatedPeriods[ToString(period3m)], period)
	period5m := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 5}
	cmp.Equal(AssociatedPeriods[ToString(period5m)], period)
	period15m := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 15}
	cmp.Equal(AssociatedPeriods[ToString(period15m)], period5m)
	period30m := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 30}
	cmp.Equal(AssociatedPeriods[ToString(period30m)], period15m)
	period1h := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1}
	cmp.Equal(AssociatedPeriods[ToString(period1h)], period30m)
	period3h := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 3}
	cmp.Equal(AssociatedPeriods[ToString(period3h)], period1h)
	period6h := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 6}
	cmp.Equal(AssociatedPeriods[ToString(period6h)], period3h)
	period12h := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 12}
	cmp.Equal(AssociatedPeriods[ToString(period12h)], period6h)
	period1d := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1}
	cmp.Equal(AssociatedPeriods[ToString(period1d)], period12h)
	period3d := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 3}
	cmp.Equal(AssociatedPeriods[ToString(period3d)], period1d)
	period1w := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	cmp.Equal(AssociatedPeriods[ToString(period1w)], period1d)
}

func Test_ToString(t *testing.T) {
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	cmp.Equal(ToString(period), "1minute")
}

func Test_offset(t *testing.T) {
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	cmp.Equal(Offset(period), 0)
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1}
	cmp.Equal(Offset(period), 0)
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1}
	cmp.Equal(Offset(period), 0)
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	i := 4 * int64(time.Minute) * 24 * 60
	cmp.Equal(Offset(period), i)
}

func Test_ToAEDKeyTimestamp(t *testing.T) {
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	var tInput int64 = 1677081282173787000
	var tMinute int64 = 1677081240000000000
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tMinute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tMinute))
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 3}
	var t3Minute int64 = 1677081240000000000
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t3Minute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t3Minute))
	}
	var t5Minute int64 = 1677081000000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 5}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t5Minute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t5Minute))
	}
	var t15Minute int64 = 1677080700000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 15}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t15Minute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t15Minute))
	}
	var t30Minute int64 = 1677079800000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 30}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t30Minute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t30Minute))
	}
	var tHour int64 = 1677078000000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tHour) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tHour))
	}

	var t3Hour int64 = 1677078000000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 3}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t3Hour) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t3Hour))
	}
	var t6Hour int64 = 1677067200000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 6}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t6Hour) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t6Hour))
	}
	var t12Hour int64 = 1677067200000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 12}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t12Hour) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t12Hour))
	}
	var tDay int64 = 1677024000000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tDay) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tDay))
	}
	var t3Day int64 = 1677024000000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 3}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t3Day) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t3Day))
	}
	var tWeek int64 = 1676851200000000000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tWeek) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tWeek))
	}
	// week has an offset, check the offset to be correct by checking 2 specific values in the week: 1 before the offset and 1 after the offset
	tInput = tMinute + 4*int64(time.Minute)*24*60
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tWeek) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tWeek))
	}
}

func Test_ToMinute(t *testing.T) {
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	if !cmp.Equal(ToMinute(period), period) {
		t.Error(cmp.Diff(ToMinute(period), 1))
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1}
	if !cmp.Equal(ToMinute(period), aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 60}) {
		t.Error(cmp.Diff(ToMinute(period), 60))
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1}
	if !cmp.Equal(ToMinute(period), aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 60 * 24}) {
		t.Error(cmp.Diff(ToMinute(period), 60*24))
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	if !cmp.Equal(ToMinute(period), aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 60 * 24 * 7}) {
		t.Error(cmp.Diff(ToMinute(period), 60*24*7))
	}
}
