package domain

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	aedgrpc "github.com/sologenic/com-fs-aed-model"
	"google.golang.org/protobuf/testing/protocmp"
)

// Test the periods lookup
func Test_AssociatedPeriods(t *testing.T) {
	// Test the lookup for 1 minute:
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period)], (*aedgrpc.Period)(nil), protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period3m := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 3}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period3m)], period, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period5m := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 5}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period5m)], period, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period15m := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 15}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period15m)], period5m, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period30m := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 30}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period30m)], period15m, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period1h := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period1h)], period30m, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period3h := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 3}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period3h)], period1h, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period6h := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 6}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period6h)], period3h, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period12h := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 12}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period12h)], period6h, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period1d := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period1d)], period12h, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period3d := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 3}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period3d)], period1d, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period1w := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	if diff := cmp.Diff(AssociatedPeriods[ToString(period1w)], period1d, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
}

func Test_ToString(t *testing.T) {
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	if s := ToString(period); s != "1_PERIOD_TYPE_MINUTE" {
		t.Errorf("Mismatch: expected '1_PERIOD_TYPE_MINUTE', got '%s'", s)
	}
}

func Test_offset(t *testing.T) {
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	if o := Offset(period); o != 0 {
		t.Errorf("Mismatch: expected 0, got %v", o)
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1}
	if o := Offset(period); o != 0 {
		t.Errorf("Mismatch: expected 0, got %v", o)
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1}
	if o := Offset(period); o != 0 {
		t.Errorf("Mismatch: expected 0, got %v", o)
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	i := int64(4 * 24 * 60 * 60)
	if o := Offset(period); o != i {
		t.Errorf("Mismatch: expected %v, got %v", i, o)
	}
}

func Test_ToAEDKeyTimestamp(t *testing.T) {
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	var tInput int64 = 1677081282
	var tMinute int64 = 1677081240
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tMinute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tMinute))
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 3}
	var t3Minute int64 = 1677081240
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t3Minute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t3Minute))
	}
	var t5Minute int64 = 1677081000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 5}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t5Minute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t5Minute))
	}
	var t15Minute int64 = 1677080700
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 15}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t15Minute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t15Minute))
	}
	var t30Minute int64 = 1677079800
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 30}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t30Minute) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t30Minute))
	}
	var tHour int64 = 1677078000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tHour) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tHour))
	}

	var t3Hour int64 = 1677078000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 3}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t3Hour) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t3Hour))
	}
	var t6Hour int64 = 1677067200
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 6}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t6Hour) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t6Hour))
	}
	var t12Hour int64 = 1677067200
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 12}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t12Hour) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t12Hour))
	}
	var tDay int64 = 1677024000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tDay) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tDay))
	}
	var t3Day int64 = 1677024000
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 3}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), t3Day) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), t3Day))
	}
	var tWeek int64 = 1676851200
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tWeek) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tWeek))
	}
	// week has an offset, check the offset to be correct by checking 2 specific values in the week: 1 before the offset and 1 after the offset
	tInput = tMinute + int64(4*24*60*60)
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	if !cmp.Equal(ToAEDKeyTimestamp(period, tInput), tWeek) {
		t.Error(cmp.Diff(ToAEDKeyTimestamp(period, tInput), tWeek))
	}
}

func Test_ToMinute(t *testing.T) {
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	if diff := cmp.Diff(ToMinute(period), period, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1}
	if diff := cmp.Diff(ToMinute(period), &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 60}, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1}
	if diff := cmp.Diff(ToMinute(period), &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 60 * 24}, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
	period = &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_WEEK, Duration: 1}
	if diff := cmp.Diff(ToMinute(period), &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 60 * 24 * 7}, protocmp.Transform()); diff != "" {
		t.Errorf("Mismatch: %s", diff)
	}
}
