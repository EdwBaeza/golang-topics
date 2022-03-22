package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const API_URL = "https://api.coindesk.com/v1/bpi/currentprice.json"

type BitcoinClient struct {
	Data map[string]map[string]string `json:"bpi"`
}

func NewBitcoinClient() *BitcoinClient {
	return &BitcoinClient{}
}

func (client *BitcoinClient) GetPrice() (float64, error) {
	httpclient := &http.Client{}
	httpRequest, _ := http.NewRequest("GET", API_URL, nil)
	response, errorResponse := httpclient.Do(httpRequest)

	if errorResponse != nil {
		log.Printf("Error Client: %s", errorResponse.Error())
		return 0.0, errorResponse
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(bodyBytes, &client)
	price := strings.Replace(client.Data["USD"]["rate"], ",", "", -1)
	return strconv.ParseFloat(price, 64)
}

func main() {
	bitcoinClient := NewBitcoinClient()
	price, _ := bitcoinClient.GetPrice()
	log.Printf("Bitcoin Price: %f", price)
}
