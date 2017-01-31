package finance

import (
	"time"

	"github.com/shopspring/decimal"
)

// quoteFields are the requested quote fields.
var quoteFields = []string{
	"s",  // Symbol
	"a",  // Ask
	"a2", // AverageDailyVolume
	"a5", // Ask Size
	"b",  // Bid
	"b4", // BookValue
	"b6", // Bid Size
	"c1", // Change
	"c4", // Currency
	"g",  // DaysLow
	"h",  // DaysHigh
	"j",  // YearLow
	"k",  // YearHigh
	"k3", // Last Trade Size
	"j1", // MarketCapitalization
	"l1", // LastTradePriceOnly
	"m3", // FiftydayMovingAverage
	"m4", // TwoHundreddayMovingAverage
	"n",  // Name
	"o",  // Open
	"p",  // PreviousClose
	"p2", // ChangeinPercent
	"d1", // LastTradeDate
	"t1", // LastTradeTime
	"v",  // Volume
	"x",  // StockExchange
	"y",  // DividendYield
	"d",  // DividendShare
	"e",  // EarningsShare
	"e7", // EPSEstimateCurrentYear
	"e8", // EPSEstimateNextYear
	"e9", // EPSEstimateNextQuarter
	"j4", // EBITDA
	"p5", // PriceSales
	"p6", // PriceBook
	"q",  // ExDividendDate
	"r",  // PERatio
	"r1", // DividendPayDate
	"r5", // PEGRatio
	"s7", // ShortRatio
	"t8", // OneyrTargetPrice
}

// Quote is the object that is returned for a quote inquiry.
type Quote struct {
	Symbol             string
	Name               string
	LastTradeTime      time.Time
	LastTradePrice     decimal.Decimal
	LastTradeSize      int
	Ask                decimal.Decimal
	AskSize            int
	Bid                decimal.Decimal
	BidSize            int
	Volume             int
	ChangeNominal      decimal.Decimal
	ChangePercent      decimal.Decimal
	Open               decimal.Decimal
	PreviousClose      decimal.Decimal
	Exchange           string
	DayLow             decimal.Decimal
	DayHigh            decimal.Decimal
	FiftyTwoWeekLow    decimal.Decimal
	FiftyTwoWeekHigh   decimal.Decimal
	Currency           string
	MarketCap          string
	FiftyDayMA         decimal.Decimal
	TwoHundredDayMA    decimal.Decimal
	AvgDailyVolume     int
	FiftyTwoWeekTarget decimal.Decimal
	ShortRatio         decimal.Decimal
	BookValue          decimal.Decimal
	EBITDA             string
	PriceSales         decimal.Decimal
	PriceBook          decimal.Decimal
	PERatio            decimal.Decimal
	PEGRatio           decimal.Decimal
	DivYield           decimal.Decimal
	DivPerShare        decimal.Decimal
	DivExDate          time.Time
	DivPayDate         time.Time
	EPS                decimal.Decimal
	EPSEstCurrentYear  decimal.Decimal
	EPSEstNextYear     decimal.Decimal
	EPSEstNextQuarter  decimal.Decimal
}

// newQuote creates a new instance of a quote.
func newQuote(row []string) *Quote {

	fields := make(map[string]string)
	for idx, v := range quoteFields {
		fields[v] = row[idx]
	}

	return &Quote{
		Symbol:             fields["s"],
		Name:               fields["n"],
		LastTradeTime:      parseDateAndTime(fields["d1"], fields["t1"]),
		LastTradePrice:     toDecimal(fields["l1"]),
		LastTradeSize:      toInt(fields["k3"]),
		Ask:                toDecimal(fields["a"]),
		AskSize:            toInt(fields["a5"]),
		Bid:                toDecimal(fields["b"]),
		BidSize:            toInt(fields["b6"]),
		Volume:             toInt(fields["v"]),
		ChangeNominal:      toDecimal(fields["c1"]),
		ChangePercent:      toDecimal(fields["p2"]),
		Open:               toDecimal(fields["o"]),
		PreviousClose:      toDecimal(fields["p"]),
		Exchange:           fields["x"],
		DayLow:             toDecimal(fields["g"]),
		DayHigh:            toDecimal(fields["h"]),
		FiftyTwoWeekLow:    toDecimal(fields["j"]),
		FiftyTwoWeekHigh:   toDecimal(fields["k"]),
		Currency:           fields["c4"],
		MarketCap:          fields["j1"],
		FiftyDayMA:         toDecimal(fields["m3"]),
		TwoHundredDayMA:    toDecimal(fields["m4"]),
		AvgDailyVolume:     toInt(fields["a2"]),
		FiftyTwoWeekTarget: toDecimal(fields["t8"]),
		ShortRatio:         toDecimal(fields["s7"]),
		BookValue:          toDecimal(fields["b4"]),
		EBITDA:             fields["j4"],
		PriceSales:         toDecimal(fields["p5"]),
		PriceBook:          toDecimal(fields["p6"]),
		PERatio:            toDecimal(fields["r"]),
		PEGRatio:           toDecimal(fields["r5"]),
		DivYield:           toDecimal(fields["y"]),
		DivPerShare:        toDecimal(fields["d"]),
		DivExDate:          parseDate(fields["q"]),
		DivPayDate:         parseDate(fields["r1"]),
		EPS:                toDecimal(fields["e"]),
		EPSEstCurrentYear:  toDecimal(fields["e7"]),
		EPSEstNextYear:     toDecimal(fields["e8"]),
		EPSEstNextQuarter:  toDecimal(fields["e9"]),
	}
}
