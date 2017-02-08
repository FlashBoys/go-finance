// Package finance

package finance

import (
	"strings"

	"github.com/shopspring/decimal"
)

const quoteURL = "http://download.finance.yahoo.com/d/quotes.csv"

// Quote is the object that is returned for a quote inquiry.
type Quote struct {
	Symbol             string          `yfin:"s"`
	Name               string          `yfin:"n"`
	LastTradeTime      Timestamp       `yfin:"t1"`
	LastTradeDate      Timestamp       `yfin:"d1"`
	LastTradePrice     decimal.Decimal `yfin:"l1"`
	Ask                decimal.Decimal `yfin:"a"`
	Bid                decimal.Decimal `yfin:"b"`
	Volume             int             `yfin:"v"`
	ChangeNominal      decimal.Decimal `yfin:"c1"`
	ChangePercent      decimal.Decimal `yfin:"p2"`
	Open               decimal.Decimal `yfin:"o"`
	PreviousClose      decimal.Decimal `yfin:"p"`
	Exchange           string          `yfin:"x"`
	DayLow             decimal.Decimal `yfin:"g"`
	DayHigh            decimal.Decimal `yfin:"h"`
	FiftyTwoWeekLow    decimal.Decimal `yfin:"j"`
	FiftyTwoWeekHigh   decimal.Decimal `yfin:"k"`
	Currency           string          `yfin:"c4"`
	MarketCap          string          `yfin:"j1"`
	FiftyDayMA         decimal.Decimal `yfin:"m3"`
	TwoHundredDayMA    decimal.Decimal `yfin:"m4"`
	AvgDailyVolume     int             `yfin:"a2"`
	FiftyTwoWeekTarget decimal.Decimal `yfin:"t8"`
	ShortRatio         decimal.Decimal `yfin:"s7"`
	BookValue          decimal.Decimal `yfin:"b4"`
	EBITDA             string          `yfin:"j4"`
	PriceSales         decimal.Decimal `yfin:"p5"`
	PriceBook          decimal.Decimal `yfin:"p6"`
	PERatio            decimal.Decimal `yfin:"r"`
	PEGRatio           decimal.Decimal `yfin:"r5"`
	DivYield           decimal.Decimal `yfin:"y"`
	DivPerShare        decimal.Decimal `yfin:"d"`
	DivExDate          Timestamp       `yfin:"q"`
	DivPayDate         Timestamp       `yfin:"r1"`
	EPS                decimal.Decimal `yfin:"e"`
	EPSEstCurrentYear  decimal.Decimal `yfin:"e7"`
	EPSEstNextYear     decimal.Decimal `yfin:"e8"`
	EPSEstNextQuarter  decimal.Decimal `yfin:"e9"`
}

// GetQuote fetches a single symbol's quote from Yahoo Finance.
func GetQuote(symbol string) (q Quote, err error) {

	params := map[string]string{
		"s": symbol,
		"f": constructFields(q),
		"e": ".csv",
	}

	t, err := fetchCSV(buildURL(quoteURL, params))
	mapFields(t[0], &q)

	return
}

// GetQuotes fetches multiple symbol's quotes from Yahoo Finance.
func GetQuotes(symbols []string) (q []Quote, err error) {

	var sq Quote
	params := map[string]string{
		"s": strings.Join(symbols[:], ","),
		"f": constructFields(sq),
		"e": ".csv",
	}

	t, err := fetchCSV(buildURL(quoteURL, params))
	for _, row := range t {
		mapFields(row, &sq)
		q = append(q, sq)
	}

	return
}
