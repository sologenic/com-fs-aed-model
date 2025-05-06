package ticker

import (
	"context"
	"errors"
	"fmt"
	sync "sync"
	"time"

	aedgrpc "github.com/sologenic/com-fs-aed-model"
	aedclient "github.com/sologenic/com-fs-aed-model/client"
	aedgrpcdmn "github.com/sologenic/com-fs-aed-model/domain"
	assetgrpc "github.com/sologenic/com-fs-asset-model"
	assetdmn "github.com/sologenic/com-fs-asset-model/domain"
	assetdmnsymbol "github.com/sologenic/com-fs-asset-model/domain/symbol"
	utilcache "github.com/sologenic/com-fs-utils-lib/go/cache"

	"github.com/sologenic/com-fs-utils-lib/go/logger"
	"github.com/sologenic/com-fs-utils-lib/models/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrTickerEmptySymbols   = errors.New("no symbols provided")
	ErrTickerTooManySymbols = errors.New("too many symbols requested")
	ErrSymbolInvalid        = errors.New("invalid symbol format")
)

const (
	MaxTickerSymbolsNumber = 40
	DefaultTickerPeriod    = 24 * time.Hour
	DefaultCacheDuration   = 1 * time.Minute
)

type TickerReadOptions struct {
	Symbols []string
	To      time.Time
	Period  time.Duration
	Network metadata.Network
}

type Tickers map[string]*TickerPoint

type Cache struct {
	Mutex *sync.RWMutex
	Data  map[string]*utilcache.LockableCache
}

func NewTickerReadOptions(symbols []string, to time.Time, period time.Duration) *TickerReadOptions {
	return &TickerReadOptions{
		Symbols: uniqueSymbols(symbols),
		To:      to,
		Period:  period,
	}
}

func (opt *TickerReadOptions) Validate() error {
	if len(opt.Symbols) == 0 {
		return ErrTickerEmptySymbols
	} else if len(opt.Symbols) > MaxTickerSymbolsNumber {
		return ErrTickerTooManySymbols
	}

	for _, symb := range opt.Symbols {
		if !validSymbol(symb) {
			return ErrSymbolInvalid
		}
	}
	return nil
}

func uniqueSymbols(symbols []string) []string {
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

func validSymbol(symbol string) bool {
	_, err := assetdmnsymbol.NewSymbolFromString(symbol)
	return err == nil
}

func GetTickers(
	ctx context.Context,
	aedClient aedgrpc.AEDServiceClient,
	assetClient assetgrpc.AssetListServiceClient,
	opt *TickerReadOptions,
	organizationID string,
	c *Cache,
) *TickerResponse {
	retvals := getTickers(ctx, aedClient, assetClient, opt, organizationID, c)
	tickers := tickersToHTTP(retvals, opt)
	return tickers.ToResponse()
}

// Tickers to http evaluates the symbols to be either non inverted or inverted and switches the volume and invertedVolume accordingly
func tickersToHTTP(tickers *Tickers, opt *TickerReadOptions) *Tickers {
	retvals := make(Tickers)
	for _, symbol := range opt.Symbols {
		if _, ok := (*tickers)[symbol]; ok {
			retvals[symbol] = (*tickers)[symbol]
			continue
		}
	}

	return &retvals
}

// Tickers are once calculated valid for up to refreshInterval seconds, however requests might come in in parallel and the cache might be empty, leading to multiple retrieval requests and subsequent caching of the same data.
// We do however do not want to serialize the requests, since that could be slow or blocking.
// The alternative used here, is that the actual cache is 15 seconds, and synchronized with the clock at refreshInterval to allow refreshes.
// As long as there is demand for the data, somewhere in the 5 second interval beyond the wanted caching period, the data will be refreshed by placing a request a refresh channel.
// This refresh channel is connect to blocking go routines per request type (symbol or group of symbols).
//
// Note: Since multiple instances of the service are running, the refresh is not 100% deduplicated: Another instances might attempt to refresh the cache at about the same time.
func getTickers(
	ctx context.Context,
	aedClient aedgrpc.AEDServiceClient,
	assetClient assetgrpc.AssetListServiceClient,
	opt *TickerReadOptions,
	organizationID string,
	c *Cache,
) *Tickers {
	// Retrieve the tickers from the AED service:
	tickerPoints := make(map[string]*TickerPoint)

	AEDs := make([]*aedgrpc.AEDs, 0, len(opt.Symbols))
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, symbol := range opt.Symbols {
		// Cache check for the symbol:
		c.Mutex.RLock()
		if cache, ok := c.Data[symbol]; ok {
			v := cache.Value.(*TickerPoint)
			tickerPoints[symbol] = v
			c.Mutex.RUnlock()
			continue
		}
		c.Mutex.RUnlock()
		wg.Add(1)
		go func(symbol string) {
			baseAEDs, err := getAED(ctx, aedClient, assetClient, symbol, opt, organizationID)
			if err != nil {
				logger.Errorf("(no cache) Error getting aed data for %s: %s", symbol, err.Error())
				wg.Done()
				return
			}
			mutex.Lock()
			AEDs = append(AEDs, baseAEDs)
			mutex.Unlock()
			wg.Done()
		}(symbol)
	}
	wg.Wait()
	tickersResp := AEDsToTickers(AEDs, opt, tickerPoints, c)
	return (*Tickers)(&tickersResp)
}

// Retrieve the AED data from the source
func getAED(
	ctx context.Context,
	aedClient aedgrpc.AEDServiceClient,
	assetClient assetgrpc.AssetListServiceClient,
	symbol string,
	opt *TickerReadOptions,
	organizationID string,
) (*aedgrpc.AEDs, error) {
	loadSymbol := &aedgrpc.AEDFilter{
		Symbol:         symbol,
		Network:        opt.Network,
		Period:         &aedgrpc.Period{Type: aedgrpc.PeriodType_PERIOD_TYPE_HOUR, Duration: 1},
		To:             timestamppb.New(opt.To),
		From:           timestamppb.New(opt.To.Add(-opt.Period)),
		Backfill:       true,
		AllowCache:     true,
		OrganizationID: organizationID,
	}
	baseAEDs, err := aedClient.Get(aedclient.AuthCtx(ctx), loadSymbol)
	if err != nil {
		return nil, fmt.Errorf("error getting aed data for %s: %w", symbol, err)
	}
	// Prevent downstream failures on empty arrays.
	if len(baseAEDs.AEDs) == 0 {
		return nil, fmt.Errorf("no aed data found for %s", symbol)
	}
	// Normalize the AED data
	for _, aed := range baseAEDs.AEDs {
		var err error
		aed, err = normalizeAED(ctx, assetClient, aed, organizationID)
		if err != nil {
			logger.Errorf("Error normalizing AED %v: %v", aed, err)
			continue
		}
	}
	return baseAEDs, nil
}

func normalizeAED(
	ctx context.Context,
	assetClient assetgrpc.AssetListServiceClient,
	aed *aedgrpc.AED,
	organizationID string,
) (*aedgrpc.AED, error) {
	symbol, err := assetdmnsymbol.NewSymbolFromString(aed.Symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to parse symbol: %w", err)
	}
	// Get base asset precision
	baseAsset, err := assetClient.GetAsset(ctx, &assetgrpc.AssetKey{
		Key: assetdmn.CreateAssetKeyStr(symbol.Base, organizationID),
	})
	if err != nil || baseAsset.AssetDetails == nil || baseAsset.AssetDetails.Denom == nil {
		return nil, fmt.Errorf("failed to fetch base precision for %v: %w", symbol.Base, err)
	}
	// Get quote asset precision
	quoteAsset, err := assetClient.GetAsset(ctx, &assetgrpc.AssetKey{
		Key: assetdmn.CreateAssetKeyStr(symbol.Quote, organizationID),
	})
	if err != nil || quoteAsset.AssetDetails == nil || quoteAsset.AssetDetails.Denom == nil {
		return nil, fmt.Errorf("failed to fetch quote precision for %v: %w", symbol.Quote, err)
	}

	return aedgrpcdmn.NormalizeAED(
		ctx,
		aed,
		organizationID,
		baseAsset.AssetDetails.Denom.Precision,
		quoteAsset.AssetDetails.Denom.Precision,
	)
}

// AEDsToTickers converts the aed data to ticker data
// The values calculated are cached for refreshInterval max (clock rounded to refreshInterval intervals) with an allowed stale period of 5s, giving an always retrieval of values
// from the cache for performance and data cost reasons (we can refresh in a go blocking routine while the data is still being served quickly)
func AEDsToTickers(AEDs []*aedgrpc.AEDs, domainOptions *TickerReadOptions, tickerPoints map[string]*TickerPoint, c *Cache) map[string]*TickerPoint {
	for _, aed := range AEDs {
		tickerPoint := calculateTickerAED(aed, domainOptions)
		tickerPoints[aed.AEDs[0].Symbol] = tickerPoint
		c.Mutex.Lock()
		c.Data[aed.AEDs[0].Symbol] = &utilcache.LockableCache{
			Value:       tickerPoint,
			LastUpdated: time.Now(),
		}
		c.Mutex.Unlock()
	}
	return tickerPoints
}

// calculate the open, high, low, close, volume and invertedVolume
// Returns a single AED with the calculated values.
func calculateTickerAED(AEDs *aedgrpc.AEDs, options *TickerReadOptions) *TickerPoint {
	if len(AEDs.AEDs) == 0 {
		return nil
	}
	fromTime := options.To.Add(-options.Period)
	// get the base AEDs for the requested period calculation.
	// The input data contains the base data to calculate the requested period over.
	// Calculate the volume:
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
