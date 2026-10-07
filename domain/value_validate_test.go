package domain

import (
	"math"
	"testing"

	"buf.build/go/protovalidate"
	aedgrpc "github.com/sologenic/com-fs-aed-model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func dayPeriod() *aedgrpc.Period {
	return &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DAY, Duration: 1}
}

func aedWithOpen(series aedgrpc.Series, open float64) *aedgrpc.AED {
	return &aedgrpc.AED{
		OrganizationID: "org-1",
		Symbol:         "org-1:utestcore:balance",
		Timestamp:      timestamppb.Now(),
		Period:         dayPeriod(),
		Series:         series,
		Value: []*aedgrpc.Value{
			{Field: aedgrpc.Field_OPEN, Float64Val: &open},
		},
	}
}

func TestValueProtovalidate(t *testing.T) {
	v, err := protovalidate.New()
	if err != nil {
		t.Fatalf("protovalidate.New: %v", err)
	}

	pos := 10.0
	zero := 0.0
	neg := -1.0
	nan := math.NaN()
	inf := math.Inf(1)
	trades := int64(1)
	negTrades := int64(-1)

	cases := []struct {
		name    string
		value   *aedgrpc.Value
		wantErr bool
	}{
		{
			name:    "open positive finite",
			value:   &aedgrpc.Value{Field: aedgrpc.Field_OPEN, Float64Val: &pos},
			wantErr: false,
		},
		{
			name:    "volume zero allowed",
			value:   &aedgrpc.Value{Field: aedgrpc.Field_VOLUME, Float64Val: &zero},
			wantErr: false,
		},
		{
			name:    "open zero allowed on Value alone",
			value:   &aedgrpc.Value{Field: aedgrpc.Field_OPEN, Float64Val: &zero},
			wantErr: false,
		},
		{
			name:    "high nan rejected",
			value:   &aedgrpc.Value{Field: aedgrpc.Field_HIGH, Float64Val: &nan},
			wantErr: true,
		},
		{
			name:    "low inf rejected",
			value:   &aedgrpc.Value{Field: aedgrpc.Field_LOW, Float64Val: &inf},
			wantErr: true,
		},
		{
			name:    "pe ratio negative finite allowed",
			value:   &aedgrpc.Value{Field: aedgrpc.Field_PE_RATIO, Float64Val: &neg},
			wantErr: false,
		},
		{
			name:    "number of trades non-negative",
			value:   &aedgrpc.Value{Field: aedgrpc.Field_NUMBER_OF_TRADES, Int64Val: &trades},
			wantErr: false,
		},
		{
			name:    "number of trades negative rejected",
			value:   &aedgrpc.Value{Field: aedgrpc.Field_NUMBER_OF_TRADES, Int64Val: &negTrades},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := v.Validate(tc.value)
			if tc.wantErr && err == nil {
				t.Fatalf("expected validation error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected validation error: %v", err)
			}
		})
	}
}

func TestAEDSeriesOHLCProtovalidate(t *testing.T) {
	v, err := protovalidate.New()
	if err != nil {
		t.Fatalf("protovalidate.New: %v", err)
	}

	cases := []struct {
		name    string
		aed     *aedgrpc.AED
		wantErr bool
	}{
		{
			name:    "trade open positive",
			aed:     aedWithOpen(aedgrpc.Series_INTERNAL_TRADES, 10.0),
			wantErr: false,
		},
		{
			name:    "trade open zero rejected",
			aed:     aedWithOpen(aedgrpc.Series_INTERNAL_TRADES, 0.0),
			wantErr: true,
		},
		{
			name:    "trade open negative rejected",
			aed:     aedWithOpen(aedgrpc.Series_INTERNAL_TRADES, -1.0),
			wantErr: true,
		},
		{
			name:    "billing open zero allowed",
			aed:     aedWithOpen(aedgrpc.Series_BILLING, 0.0),
			wantErr: false,
		},
		{
			name:    "billing open positive allowed",
			aed:     aedWithOpen(aedgrpc.Series_BILLING, 10.0),
			wantErr: false,
		},
		{
			name:    "billing open negative rejected",
			aed:     aedWithOpen(aedgrpc.Series_BILLING, -1.0),
			wantErr: true,
		},
		{
			name: "billing close zero allowed",
			aed: func() *aedgrpc.AED {
				open, close := 0.0, 0.0
				return &aedgrpc.AED{
					OrganizationID: "org-1",
					Symbol:         "org-1:utestcore:balance",
					Timestamp:      timestamppb.Now(),
					Period:         dayPeriod(),
					Series:         aedgrpc.Series_BILLING,
					Value: []*aedgrpc.Value{
						{Field: aedgrpc.Field_OPEN, Float64Val: &open},
						{Field: aedgrpc.Field_CLOSE, Float64Val: &close},
					},
				}
			}(),
			wantErr: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := v.Validate(tc.aed)
			if tc.wantErr && err == nil {
				t.Fatalf("expected validation error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected validation error: %v", err)
			}
		})
	}
}

func TestPeriodProtovalidate(t *testing.T) {
	v, err := protovalidate.New()
	if err != nil {
		t.Fatalf("protovalidate.New: %v", err)
	}

	if err := v.Validate(&aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 1}); err != nil {
		t.Fatalf("valid period: %v", err)
	}
	if err := v.Validate(&aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_DO_NOT_USE, Duration: 1}); err == nil {
		t.Fatalf("expected error for unused period type")
	}
	if err := v.Validate(&aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_MINUTE, Duration: 0}); err == nil {
		t.Fatalf("expected error for zero duration")
	}
}
