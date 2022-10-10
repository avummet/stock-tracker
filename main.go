package main

import (
	"github.com/avummet/stock-tracker/requester"
    "github.com/avummet/stock-tracker/server"
	"fmt"
	"os"
	"github.com/avummet/stock-tracker/helper"
)

const (
	appName = "stock-tracker"
)

func main(){
	// get env vars
	ticker := os.Getenv(SYMBOL)
	ndays := os.Getenv(NDAYS)

	//make request to api and get stock data
	resp := requester.MakeRequest(ticker, ndays)

	//
	data := helper.

	// webservice
	server.RunHTTPServer()
	

}