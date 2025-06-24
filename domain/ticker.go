package domain

import (
	"context"
	"fmt"
	sync "sync"
	"time"

	organizationgrpc "github.com/sologenic/com-fs-admin-organization-model"
	aedgrpc "github.com/sologenic/com-fs-aed-model"
	aedclient "github.com/sologenic/com-fs-aed-model/client"
	assetgrpc "github.com/sologenic/com-fs-asset-model"
	assetdmnsymbol "github.com/sologenic/com-fs-asset-model/domain/symbol"
	utilcache "github.com/sologenic/com-fs-utils-lib/go/cache"

	"github.com/sologenic/com-fs-utils-lib/go/logger"
	"github.com/sologenic/com-fs-utils-lib/models/metadata"
)

const (
	MaxTickerSymbolsNumber = 40
	DefaultTickerPeriod    = 24 * time.Hour
	DefaultCacheDuration   = 1 * time.Minute
)

type Tickers map[string]*aedgrpc.AED

type TickerResponse struct {
	Tickers []*aedgrpc.AED `json:"tickers"`
}

// Currently only 24h tickers are used.
type TickerReadOptions struct {
	Symbols []string
	Period  time.Duration
	Network metadata.Network
	Series  aedgrpc.Series
}

func NewTickerReadOptions(symbols []string, series aedgrpc.Series) *TickerReadOptions {
	return &TickerReadOptions{
		Symbols: uniqueSymbols(symbols),
		Series:  series,
		Period:  DefaultTickerPeriod,
	}
}

func (opt *TickerReadOptions) Validate() Error {
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

func validSymbol(strSymbol string) bool {
	_, err := assetdmnsymbol.NewSymbolFromString(strSymbol)
	return err == nil
}

// GetTickers retrieves tickers for the specified options
func GetTickers(
	ctx context.Context,
	aedClient aedgrpc.AEDServiceClient,
	assetClient assetgrpc.AssetListServiceClient,
	orgClient organizationgrpc.OrganizationServiceClient,
	opt *TickerReadOptions,
	organizationID string,
	tickerCache *utilcache.Cache,
	assetCache *utilcache.Cache,
) *TickerResponse {
	retvals := getTickers(ctx, aedClient, assetClient, orgClient, opt, organizationID, tickerCache, assetCache)
	tickers := tickersToHTTP(retvals, opt)
	return tickers.ToResponse()
}

func (t *Tickers) ToResponse() *TickerResponse {
	tickers := make([]*aedgrpc.AED, 0, len(*t))
	for _, ticker := range *t {
		tickers = append(tickers, ticker)
	}
	return &TickerResponse{Tickers: tickers}
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
	orgClient organizationgrpc.OrganizationServiceClient,
	opt *TickerReadOptions,
	organizationID string,
	tickerCache *utilcache.Cache,
	assetCache *utilcache.Cache,
) *Tickers {
	// Retrieve the tickers from the AED service:
	tickerAEDs := make(map[string]*aedgrpc.AED)

	AEDs := make([]*aedgrpc.AEDs, 0, len(opt.Symbols))
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, symbol := range opt.Symbols {
		// Cache check for the symbol:
		tickerCache.Mutex.RLock()
		if cache, ok := tickerCache.Data[symbol]; ok {
			v := cache.Value.(*aedgrpc.AED) // Changed from *TickerPoint to *aedgrpc.AED
			tickerAEDs[symbol] = v
			tickerCache.Mutex.RUnlock()
			continue
		}
		tickerCache.Mutex.RUnlock()
		wg.Add(1)
		go func(symbol string) {
			baseAEDs, err := getAED(ctx, aedClient, assetClient, orgClient, symbol, opt, organizationID, assetCache)
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
	tickersResp := AEDsToTickers(AEDs, opt, tickerAEDs, tickerCache)
	return (*Tickers)(&tickersResp)
}

// getAED retrieves AED data from the source
func getAED(
	ctx context.Context,
	aedClient aedgrpc.AEDServiceClient,
	assetClient assetgrpc.AssetListServiceClient,
	orgClient organizationgrpc.OrganizationServiceClient,
	symbol string,
	opt *TickerReadOptions,
	organizationID string,
	assetCache *utilcache.Cache,
) (*aedgrpc.AEDs, error) {
	latestRequest := &aedgrpc.LatestRequest{
		Symbol:         symbol,
		Network:        opt.Network,
		Series:         opt.Series,
		OrganizationID: organizationID,
	}
	latestAED, err := aedClient.GetLatest(aedclient.AuthCtx(ctx), latestRequest)
	if err != nil {
		return nil, fmt.Errorf("error getting latest aed data for %s: %w", symbol, err)
	}
	if latestAED == nil {
		return nil, fmt.Errorf("no latest aed data found for %s", symbol)
	}
	// Normalize the AED data
	normalizedAED, err := NormalizeAED(ctx, assetClient, orgClient, latestAED, organizationID, assetCache)
	if err != nil {
		logger.Errorf("Error normalizing AED %v: %v", latestAED, err)
		return nil, fmt.Errorf("error normalizing aed data for %s: %w", symbol, err)
	}
	return &aedgrpc.AEDs{AEDs: []*aedgrpc.AED{normalizedAED}}, nil
}

// AEDsToTickers converts the aed data to ticker data
// The values calculated are cached for refreshInterval max (clock rounded to refreshInterval intervals) with an allowed stale period of 5s, giving an always retrieval of values
// from the cache for performance and data cost reasons (we can refresh in a go blocking routine while the data is still being served quickly)
func AEDsToTickers(AEDs []*aedgrpc.AEDs, domainOptions *TickerReadOptions, tickerAEDs map[string]*aedgrpc.AED, tickerCache *utilcache.Cache) map[string]*aedgrpc.AED {
	for _, aed := range AEDs {
		tickerAED := calculateTickerAED(aed, domainOptions)
		tickerAEDs[aed.AEDs[0].Symbol] = tickerAED
		tickerCache.Mutex.Lock()
		tickerCache.Data[aed.AEDs[0].Symbol] = &utilcache.LockableCache{
			Value:       tickerAED,
			LastUpdated: time.Now(),
		}
		tickerCache.Mutex.Unlock()
	}
	return tickerAEDs
}

// calculate the open, high, low, close, volume and invertedVolume
// Returns a single AED with the calculated values.
func calculateTickerAED(AEDs *aedgrpc.AEDs, options *TickerReadOptions) *aedgrpc.AED {
	if len(AEDs.AEDs) == 0 {
		return nil
	}
	// Use the latest AED data directly since we're getting the most recent data
	latestAED := AEDs.AEDs[0]

	// Extract values using helper functions
	open := GetFloatValue(latestAED, aedgrpc.Field_OPEN)
	high := GetFloatValue(latestAED, aedgrpc.Field_HIGH)
	low := GetFloatValue(latestAED, aedgrpc.Field_LOW)
	close := GetFloatValue(latestAED, aedgrpc.Field_CLOSE)
	volume := GetFloatValue(latestAED, aedgrpc.Field_VOLUME)
	invertedVolume := GetFloatValue(latestAED, aedgrpc.Field_INVERTED_VOLUME)
	marketCap := GetFloatValue(latestAED, aedgrpc.Field_MARKET_CAP)
	eps := GetFloatValue(latestAED, aedgrpc.Field_EPS)
	per := GetFloatValue(latestAED, aedgrpc.Field_PE_RATIO)
	yield := GetFloatValue(latestAED, aedgrpc.Field_YIELD)

	// Use the actual timestamp from the latest data
	actualTimestamp := latestAED.Timestamp.AsTime()

	tickerAED := &aedgrpc.AED{
		Symbol:         latestAED.Symbol,
		OrganizationID: latestAED.OrganizationID,
		Timestamp:      latestAED.Timestamp,
		Period: &aedgrpc.Period{
			Type:     aedgrpc.PeriodType_PERIOD_TYPE_DAY,
			Duration: int32(options.Period.Hours() / 24),
		},
		MetaData: latestAED.MetaData,
		Series:   latestAED.Series,
		Value:    []*aedgrpc.Value{},
	}

	// Set values equivalent to TickerPoint fields
	SetFloatValue(tickerAED, aedgrpc.Field_OPEN, open)
	SetFloatValue(tickerAED, aedgrpc.Field_HIGH, high)
	SetFloatValue(tickerAED, aedgrpc.Field_LOW, low)
	SetFloatValue(tickerAED, aedgrpc.Field_CLOSE, close)
	SetFloatValue(tickerAED, aedgrpc.Field_LAST_PRICE, close)
	SetFloatValue(tickerAED, aedgrpc.Field_FIRST_PRICE, open)
	SetFloatValue(tickerAED, aedgrpc.Field_VOLUME, volume)
	SetFloatValue(tickerAED, aedgrpc.Field_INVERTED_VOLUME, invertedVolume)
	SetIntValue(tickerAED, aedgrpc.Field_OPEN_TIME, actualTimestamp.Unix())
	SetIntValue(tickerAED, aedgrpc.Field_CLOSE_TIME, actualTimestamp.Unix())
	SetFloatValue(tickerAED, aedgrpc.Field_MARKET_CAP, marketCap)
	SetFloatValue(tickerAED, aedgrpc.Field_EPS, eps)
	SetFloatValue(tickerAED, aedgrpc.Field_PE_RATIO, per)
	SetFloatValue(tickerAED, aedgrpc.Field_YIELD, yield)
	return tickerAED
}
