package ticker

import (
	"errors"
	"time"

	aedgrpc "github.com/sologenic/com-fs-aed-model"
	aedgrpcdmn "github.com/sologenic/com-fs-aed-model/domain"
	assetdmnsymbol "github.com/sologenic/com-fs-asset-model/domain/symbol"
	"github.com/sologenic/com-fs-utils-lib/models/metadata"
)

var (
	ErrTickerEmptySymbols   = errors.New("no symbols provided")
	ErrTickerTooManySymbols = errors.New("too many symbols requested")
	ErrSymbolInvalid        = errors.New("invalid symbol format")
)

const (
	MaxTickerSymbolsNumber = 40
	DefaultTickerPeriod    = 24 * time.Hour
)

type TickerReadOptions struct {
	Symbols []string
	To      time.Time
	Period  time.Duration
	Network metadata.Network
}

type Tickers map[string]*TickerPoint

func NewTickerReadOptions(symbols []string, to time.Time, period time.Duration) *TickerReadOptions {
	return &TickerReadOptions{
		Symbols: UniqueSymbols(symbols),
		To:      to,
		Period:  period,
	}
}

func UniqueSymbols(symbols []string) []string {
	keys := make(map[string]bool, len(symbols))
	var res []string

	for _, symbol := range symbols {
		if _, value := keys[symbol]; !value {
			keys[symbol] = true
			res = append(res, symbol)
		}
	}
	return res
}

func ValidateSymbols(symbols []string) error {
	if len(symbols) == 0 {
		return ErrTickerEmptySymbols
	} else if len(symbols) > MaxTickerSymbolsNumber {
		return ErrTickerTooManySymbols
	}

	for _, symb := range symbols {
		if !ValidSymbol(symb) {
			return ErrSymbolInvalid
		}
	}
	return nil
}

func ValidSymbol(symbol string) bool {
	_, err := assetdmnsymbol.NewSymbolFromString(symbol)
	return err == nil
}

// Calculate the open, high, low, close, volume and invertedVolume
// Returns a single AED with the calculated values.
func NewTickerPointFromAEDs(AEDs *aedgrpc.AEDs, options *TickerReadOptions) *TickerPoint {
	if len(AEDs.AEDs) == 0 {
		return nil
	}
	fromTime := options.To.Add(-options.Period)
	// get the base AEDs for the requested period calculation.
	// The input data contains the base data to calculate the requested period over.
	// Calculate the volume:
	// Calculate the aggregate values from the AEDs
	var open, close, low, high, volume, invertedVolume, firstPrice, marketCap, eps, per, yield float64
	// Assumption is that the data might not be ordered by time.
	var tStart, tEnd time.Time
	for _, baseAEDs := range AEDs.AEDs {
		// Extract values using helper functions
		lowVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_LOW)
		highVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_HIGH)
		volumeVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_VOLUME)
		invertedVolumeVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_INVERTED_VOLUME)
		openVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_OPEN)
		closeVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_CLOSE)
		marketCapVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_MARKET_CAP)
		epsVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_EPS)
		perVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_PE_RATIO)
		yieldVal := aedgrpcdmn.GetFloatValue(baseAEDs, aedgrpc.Field_YIELD)

		if low == 0.0 || low > lowVal {
			low = lowVal
		}
		if high < highVal {
			high = highVal
		}
		volume += volumeVal
		invertedVolume += invertedVolumeVal
		// calculate the open:
		if tStart.IsZero() || tStart.After(baseAEDs.Timestamp.AsTime()) {
			open = openVal
			tStart = baseAEDs.Timestamp.AsTime()
		}
		// Calculate the first price: This is the first price in the time period (so timestamp > From)
		if tStart.After(fromTime) && (firstPrice == 0.0 || firstPrice > openVal) {
			firstPrice = openVal
		}
		// calculate the close:
		if tEnd.IsZero() || tEnd.Before(baseAEDs.Timestamp.AsTime()) {
			close = closeVal
			tEnd = baseAEDs.Timestamp.AsTime()
			marketCap = marketCapVal
			eps = epsVal
			per = perVal
			yield = yieldVal
		}

	}
	// The calculated values might cover the requested time period or might be from before the requested time period:
	// If they are from before the requested time period, volume is 0, high, low and close are the same as the open.
	if tStart.Before(fromTime) {
		volume = 0
		high = open
		low = open
		close = open
		firstPrice = 0.0
		invertedVolume = 0.0
	}

	t := &TickerPoint{
		Symbol:         AEDs.AEDs[0].Symbol,
		OpenTime:       fromTime.Unix(),
		CloseTime:      options.To.Unix(),
		OpenPrice:      open,
		HighPrice:      high,
		LowPrice:       low,
		LastPrice:      close,
		FirstPrice:     firstPrice,
		Volume:         volume,
		InvertedVolume: invertedVolume,
		MarketCap:      marketCap,
		EPS:            eps,
		PERatio:        per,
		Yield:          yield,
	}
	return t
}

func (t *Tickers) ToResponse() *TickerResponse {
	resp := &TickerResponse{
		Tickers: make([]*TickerPoint, 0, len(*t)),
	}
	for _, ticker := range *t {
		resp.Tickers = append(resp.Tickers, ticker)
	}
	return resp
}

func (t *Tickers) FilterBySymbols(symbols []string) *Tickers {
	result := make(Tickers)
	for _, symbol := range symbols {
		if ticker, ok := (*t)[symbol]; ok {
			result[symbol] = ticker
		}
	}
	return &result
}
