// Package finance

package finance

import (
	"fmt"
	"strings"
)

const quoteURL = "http://download.finance.yahoo.com/d/quotes.csv"

// GetQuote fetches a single symbol's quote from Yahoo Finance.
func GetQuote(symbol string) (*Quote, error) {

	params := map[string]string{
		"s": symbol,
		"f": strings.Join(quoteFields[:], ""),
		"e": ".csv",
	}

	table, err := getQuotesTable(buildURL(quoteURL, params))
	if err != nil {
		return nil, err
	}
	
	quotes := generateQuotes(table)
	if len(quotes) == 0 {
		return nil, fmt.Errorf("The symbol does not exist")
	}
	
	return quotes[0], nil
}

// GetQuotes fetches multiple symbol's quotes from Yahoo Finance.
func GetQuotes(symbols []string) ([]*Quote, error) {

	params := map[string]string{
		"s": strings.Join(symbols[:], ","),
		"f": strings.Join(quoteFields[:], ""),
		"e": ".csv",
	}

	table, err := getQuotesTable(buildURL(quoteURL, params))
	if err != nil {
		return nil, err
	}
	return generateQuotes(table), nil
}

// getQuotesTable fetches the quotes data table from the endpoint.
func getQuotesTable(url string) ([][]string, error) {

	table, err := requestCSV(url)
	if err != nil {
		return nil, fmt.Errorf("request table error:  (error was: %s)\n", err.Error())
	}
	return table, nil
}

// generateQuotes turns the raw table data of quotes into proper quote structs.
func generateQuotes(table [][]string) (quotes []*Quote) {

	for _, row := range table {
		quotes = append(quotes, newQuote(row))
	}
	return quotes
}
