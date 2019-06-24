package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CurrencyModel struct {
	Responses []Response
}

type Response struct {
	Code                        string `json:"code"`
	Name                        string `json:"name"`
	IsCrypto                    bool   `json:"is_crypto"`
	MinimalAmount               string `json:"minimal_amount"`
	MaximalAmount               string `json:"maximal_amount"`
	IsBaseOfEnabledPair         bool   `json:"is_base_of_enabled_pair"`
	IsQuoteOfEnabledPair        bool   `json:"is_quote_of_enabled_pair"`
	HasEnabledPairs             bool   `json:"has_enabled_pairs"`
	IsBaseOfEnabledPairForTest  bool   `json:"is_base_of_enabled_pair_for_test"`
	IsQuoteOfEnabledPairForTest bool   `json:"is_quote_of_enabled_pair_for_test"`
	HasEnabledPairsForTest      bool   `json:"has_enabled_pairs_for_test"`
	WithdrawalFee               string `json:"withdrawal_fee"`
}

func (currencyModel *CurrencyModel) deserialize(data []byte) {
	err := json.Unmarshal(data, &currencyModel.Responses)
	errorOutput(err)
}

func homeRouterHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "My first API v1")
	errorOutput(err)
}

func (currencyModel *CurrencyModel) listRouterHandler(writer http.ResponseWriter, request *http.Request) {
	response, err := json.Marshal(currencyModel)
	errorOutput(err)
	_, err = writer.Write(response)
	errorOutput(err)
}

func (currencyModel *CurrencyModel) firstElementHandler(writer http.ResponseWriter, request *http.Request) {
	response, err := json.Marshal(currencyModel.Responses[0])
	errorOutput(err)
	_, err = writer.Write(response)
	errorOutput(err)
}

func (currencyModel *CurrencyModel) latestElementHandler(writer http.ResponseWriter, request *http.Request) {
	response, err := json.Marshal(currencyModel.Responses[len(currencyModel.Responses)-1])
	errorOutput(err)
	_, err = writer.Write(response)
	errorOutput(err)
}

func (currencyModel *CurrencyModel) handlers() {
	http.HandleFunc("/", homeRouterHandler)
	http.HandleFunc("/list", currencyModel.listRouterHandler)
	http.HandleFunc("/list/first", currencyModel.firstElementHandler)
	http.HandleFunc("/list/latest", currencyModel.latestElementHandler)
	err := http.ListenAndServe(":8080", nil)
	errorOutput(err)
}

func main() {
	var currencyModel = CurrencyModel{}
	data, err := http.Get("https://api.nexchange.io/en/api/v1/currency/")
	errorOutput(err)
	bytes, err := ioutil.ReadAll(data.Body)
	errorOutput(err)
	currencyModel.deserialize(bytes)
	currencyModel.handlers()
}

func errorOutput(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
