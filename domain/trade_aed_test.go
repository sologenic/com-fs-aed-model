package domain

import (
	"math"
	"testing"
	"time"

	aedgrpc "github.com/sologenic/com-fs-aed-model"
	"github.com/sologenic/com-fs-utils-lib/models/metadata"
	"google.golang.org/genproto/googleapis/type/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewAEDFromTrade(t *testing.T) {
	period := &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}
	bucketTs := timestamppb.New(time.Unix(1_700_000_000, 0))
	got := NewAEDFromTrade(
		"97ffe1fa-e1e4-4b98-b30e-1c84c2cf64bd",
		"AAPL:USD",
		period,
		bucketTs,
		metadata.Network_TESTNET,
		100.5,
		2.0,
	)

	if got.Series != aedgrpc.Series_INTERNAL_TRADES {
		t.Fatalf("expected Series_INTERNAL_TRADES, got %v", got.Series)
	}
	if GetFloatValue(got, aedgrpc.Field_OPEN) != 100.5 {
		t.Fatalf("open=%v", GetFloatValue(got, aedgrpc.Field_OPEN))
	}
	if GetFloatValue(got, aedgrpc.Field_VOLUME) != 2.0 {
		t.Fatalf("volume=%v", GetFloatValue(got, aedgrpc.Field_VOLUME))
	}
}

func TestMergeTradeIntoAED(t *testing.T) {
	existing := NewAEDFromTrade(
		"org", "AAPL:USD",
		&aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1},
		timestamppb.Now(),
		metadata.Network_TESTNET,
		100, 1,
	)

	MergeTradeIntoAED(existing, 110, 3)
	if GetFloatValue(existing, aedgrpc.Field_OPEN) != 100 {
		t.Fatalf("open=%v", GetFloatValue(existing, aedgrpc.Field_OPEN))
	}
	if GetFloatValue(existing, aedgrpc.Field_HIGH) != 110 {
		t.Fatalf("high=%v", GetFloatValue(existing, aedgrpc.Field_HIGH))
	}
	if GetFloatValue(existing, aedgrpc.Field_VOLUME) != 4 {
		t.Fatalf("volume=%v", GetFloatValue(existing, aedgrpc.Field_VOLUME))
	}
	if GetIntValue(existing, aedgrpc.Field_NUMBER_OF_TRADES) != 2 {
		t.Fatalf("trades=%d", GetIntValue(existing, aedgrpc.Field_NUMBER_OF_TRADES))
	}
}

func TestDecimalToFloat(t *testing.T) {
	got, err := DecimalToFloat(&decimal.Decimal{Value: "12.34"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(got-12.34) > 1e-9 {
		t.Fatalf("got %v", got)
	}
	// Domain rules (positive / finite) are enforced by protovalidate, not conversion.
	if _, err := DecimalToFloat(&decimal.Decimal{Value: "0"}); err != nil {
		t.Fatalf("zero should convert: %v", err)
	}
	if _, err := DecimalToFloat(nil); err == nil {
		t.Fatal("expected error for nil")
	}
}
