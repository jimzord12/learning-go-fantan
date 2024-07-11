package opgfuncs

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jimzord12/learning-go-fantan/cmd/opg-analysis/global"
	"github.com/jimzord12/learning-go-fantan/cmd/opg-analysis/models"
	"github.com/jimzord12/learning-go-fantan/cmd/opg-analysis/opgfuncs/helpers"
)

type Stocks = []models.Stock
type Stock = models.Stock
type Position = models.Position

func FilterStockData(stocks Stocks) Stocks {
	return slices.DeleteFunc(stocks, func(s Stock) bool {
		return math.Abs(s.Gap) < 0.1
	})
}

func LoadStockData(path string) (Stocks, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return Stocks{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
		return Stocks{}, err
	}

	rows = slices.Delete(rows, 0, 1)

	var stocks Stocks

	for _, row := range rows {
		ticker := row[0]
		gap, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			fmt.Println(err)
			continue
		}

		openingPrice, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Println(err)
			continue
		}

		stocks = append(stocks, models.Stock{
			Ticker:       ticker,
			Gap:          gap,
			OpeningPrice: openingPrice,
		})
	}

	return stocks, nil
}

func Calculate(gapPercent, openingPrice float64) Position {
	closingPrice := openingPrice / (1 + gapPercent)
	gapValue := closingPrice - openingPrice
	profitFromGap := global.ProfitPercent * gapValue

	stopLoss := openingPrice - profitFromGap
	takeProfit := openingPrice + profitFromGap

	shares := int(global.MaxLossPerTrade / math.Abs(stopLoss-openingPrice))

	profit := math.Abs(openingPrice-takeProfit) * float64(shares)
	profit = math.Round(profit*100) / 100

	return Position{
		EntryPrice:      helpers.Round2Dec(openingPrice),
		Shares:          shares,
		TakeProfitPrice: helpers.Round2Dec(takeProfit),
		StopLossPrice:   helpers.Round2Dec(stopLoss),
		Profit:          helpers.Round2Dec(profit),
	}
}

// Name only the file path, not the extention.
func DeliverJSON(filePath string, selections []models.Selection) error {
	if strings.Contains(filePath, ".") {
		return fmt.Errorf("please do not add an extention to the file's path. (ex. myfile.json), only enter the path.")
	}

	// 1. Creating the file
	file, err := os.Create(filePath + ".json")
	if err != nil {
		return fmt.Errorf("[ERROR]: failed to create file: %w", err)
	}
	defer file.Close()

	// 2. Creating the JSON Encoder
	encoder := json.NewEncoder(file)
	err = encoder.Encode(selections)
	if err != nil {
		return fmt.Errorf("[ERROR]: failed to encode the file into json: %w", err)
	}
	log.Printf("File successfully created at: %s", filePath)

	return nil
}
