package main

import (
	"github.com/stock-tracker/config"
	helper "github.com/stock-tracker/helpers"
	"github.com/stock-tracker/requester"
	"github.com/stock-tracker/server"
)

const (
	appName = "stock-tracker"
)

func main() {
	// get env vars
	cfg, err := config.InitConfig(appName)
	if err != nil {
		panic(err)
	}

	//make request to api and get stock data
	requester.MakeRequest(cfg.Ticker, cfg.Numdays)

	// Parse data struct to extract NUmdays of prices and calc avg.price
	helper.Parsestruct()

	// webservice to expose resultset
	serverConfig, err := config.InitServerConfig(cfg)

	if err != nil {
		panic(err)
	}
	server.RunHTTPServer(serverConfig)

}
