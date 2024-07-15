package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jimzord12/learning-go-fantan/cmd/opg-analysis/opgglobal"
	"github.com/jimzord12/learning-go-fantan/cmd/opg-analysis/test/api/apimodels"
)

type Attributes = apimodels.Attributes
type SeekingAlplaNews = apimodels.SeekingAlplaNews
type ResponseData = apimodels.SeekingAlplaNewsResponse

func FetchNews(ticker string) (ResponseData, error) {
	// 1. Creating the Request
	req, err := http.NewRequest(http.MethodGet, opgglobal.Url+ticker, nil)
	if err != nil {
		fmt.Println(err)
		return ResponseData{}, err
	}

	// 1.1 Creating the Request's Header
	req.Header.Add(opgglobal.ApiKeyHeader, opgglobal.ApiKey)

	// 2. Creating the HTTP Client (Like Postman, ThunderClient, Bruno, etc.)
	client := &http.Client{}

	// 3. Sending the Request using the HTTP Client
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ResponseData{}, err
	}

	// 4. Checking the Response's Status Code
	if res.StatusCode < 200 || res.StatusCode > 299 {
		fmt.Println("unsucessful status code (%d) received", res.StatusCode)
	}

	var resData ResponseData
	// 5. Decoding the received JSON into a Go value (struct)
	json.NewDecoder(res.Body).Decode(&resData)

	return resData, nil
}
