package opganalysis

import (
	"fmt"
	"log"

	"github.com/jimzord12/learning-go-fantan/cmd/opg-analysis/models"
	"github.com/jimzord12/learning-go-fantan/cmd/opg-analysis/opgfuncs"
	"github.com/jimzord12/learning-go-fantan/cmd/opg-analysis/test/api"
)

type Stocks = []models.Stock
type Stock = models.Stock
type Position = models.Position
type Selection = models.Selection

func main() {
	// 1. Load the Stocks
	stocks, err := opgfuncs.LoadStockData("./opg.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2. Filter the Stocks
	stocks = opgfuncs.FilterStockData(stocks)
	fmt.Println("Stocks:", stocks)

	// 3. Calculate the Stocks
	var selections []Selection
	for _, stock := range stocks {
		position := opgfuncs.Calculate(stock.Gap, stock.OpeningPrice)

		articles, err := api.FetchNews(stock.Ticker)
		if err != nil {
			log.Println(err)
			continue
		} else {
			log.Printf("Fetched %d articles about %s", len(articles.Articles), stock.Ticker)
		}

		sel := Selection{
			Ticker:   stock.Ticker,
			Position: position,
			Articles: articles.Articles,
		}

		selections = append(selections, sel)
	}

	// 4. Fetch the News from the Internet

}
