package domain

import "errors"

// Business logic errors.
type Error interface {
	error
}

var (
	ErrSymbolInvalid        Error = errors.New("symbol is invalid")
	ErrTickerTooManySymbols Error = errors.New("too many symbols")
	ErrTickerEmptySymbols   Error = errors.New("empty tickers symbols")
	ErrTickerPeriodInvalid  Error = errors.New("invalid tickers period")
	ErrIncorrectRequestParm Error = errors.New("incorrect request parameter")
)
