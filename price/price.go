package price


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Define structs to match the JSON response structure
type Response struct {
	Data map[string]CryptoData `json:"data"`
}

type CryptoData struct {
	Name  string `json:"name"`
	Quote Quote  `json:"quote"`
}

type Quote struct {
	USD Price `json:"USD"`
}

type Price struct {
	Price float64 `json:"price"`
}

var currentPrice float64

func CheckPrice(token string) float64 {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("symbol", token) // Specify the cryptocurrency symbol (e.g., BTC for Bitcoin)
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "8249c2f9-84cd-4ea0-8c75-35aa21c63083")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// Parse the JSON response
	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// Extract and print the price data
	for _, crypto := range response.Data {
		currentPrice = crypto.Quote.USD.Price
	}


	return currentPrice
}
