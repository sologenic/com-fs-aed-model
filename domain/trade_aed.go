package domain

import (
	"fmt"
	"math"
	"strconv"

	aedgrpc "github.com/sologenic/com-fs-aed-model"
	"github.com/sologenic/com-fs-utils-lib/models/metadata"
	typedecimal "google.golang.org/genproto/googleapis/type/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DecimalToPositiveFiniteFloat converts a google.type.Decimal to float64.
// Rejects nil/empty, non-numeric, non-positive, NaN, and Inf values.
func DecimalToPositiveFiniteFloat(d *typedecimal.Decimal) (float64, error) {
	if d == nil {
		return 0, fmt.Errorf("decimal is nil")
	}
	value := d.GetValue()
	if value == "" {
		return 0, fmt.Errorf("decimal value is empty")
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	if f <= 0 || math.IsNaN(f) || math.IsInf(f, 0) {
		return 0, fmt.Errorf("decimal value must be a positive finite number: %s", value)
	}
	return f, nil
}

// NewAEDFromTrade builds a new trade-series AED bucket from one fill.
func NewAEDFromTrade(
	organizationID, symbol string,
	period *aedgrpc.Period,
	bucketTs *timestamppb.Timestamp,
	network metadata.Network,
	price, qty float64,
) *aedgrpc.AED {
	return &aedgrpc.AED{
		OrganizationID: organizationID,
		Symbol:         symbol,
		Timestamp:      bucketTs,
		Period:         period,
		Series:         aedgrpc.Series_INTERNAL_TRADES,
		MetaData: &metadata.MetaData{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Network:   network,
		},
		Value: []*aedgrpc.Value{
			{Field: aedgrpc.Field_OPEN, Float64Val: new(price)},
			{Field: aedgrpc.Field_HIGH, Float64Val: new(price)},
			{Field: aedgrpc.Field_LOW, Float64Val: new(price)},
			{Field: aedgrpc.Field_CLOSE, Float64Val: new(price)},
			{Field: aedgrpc.Field_VOLUME, Float64Val: new(qty)},
			{Field: aedgrpc.Field_NUMBER_OF_TRADES, Int64Val: new(int64(1))},
		},
	}
}

// MergeTradeIntoAED folds one fill into an existing bucket.
// Open stays. High/low expand. Close becomes the latest price. Volume and trade count accumulate.
func MergeTradeIntoAED(existing *aedgrpc.AED, price, qty float64) {
	high := math.Max(GetFloatValue(existing, aedgrpc.Field_HIGH), price)
	low := GetFloatValue(existing, aedgrpc.Field_LOW)
	if low == 0 {
		low = price
	} else {
		low = math.Min(low, price)
	}
	volume := GetFloatValue(existing, aedgrpc.Field_VOLUME) + qty
	trades := GetIntValue(existing, aedgrpc.Field_NUMBER_OF_TRADES) + 1

	SetFloatValue(existing, aedgrpc.Field_HIGH, high)
	SetFloatValue(existing, aedgrpc.Field_LOW, low)
	SetFloatValue(existing, aedgrpc.Field_CLOSE, price)
	SetFloatValue(existing, aedgrpc.Field_VOLUME, volume)
	SetIntValue(existing, aedgrpc.Field_NUMBER_OF_TRADES, trades)
	if existing.MetaData == nil {
		existing.MetaData = &metadata.MetaData{}
	}
	existing.MetaData.UpdatedAt = timestamppb.Now()
}
