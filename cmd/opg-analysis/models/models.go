package models

import "github.com/jimzord12/learning-go-fantan/cmd/opg-analysis/test/api/apimodels"

type Article = apimodels.SeekingAlplaNews

type Stock struct {
	Ticker       string
	Gap          float64
	OpeningPrice float64
}

type Position struct {
	// The price at which to buy or sell
	EntryPrice float64
	// How many shares to buy or sell
	Shares int
	// The price at which to exit and take my profit
	TakeProfitPrice float64
	// The price at which to stop my loss if the stock doesn't go our way
	StopLossPrice float64
	// Expected final Profit
	Profit float64
}

type Selection struct {
	Ticker string
	Position
	Articles []Article
}
