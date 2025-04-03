package main

import (
	"log"
	"os"
	"time"

	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/generate/spot/spotpublic"
	"github.com/FuYuanDe2024/kucoin-universal-sdk/sdk/golang/pkg/types"
)

func onKlineData(topic, subject string, tradeEvent *spotpublic.KlinesEvent) error {

	log.Printf("KLine for %s @ %s: %+v", tradeEvent.Symbol, time.Unix(tradeEvent.Time, 0), tradeEvent.Candles)
	return nil
}

func KlineExample() {
	// Use the default logger or supply your custom logger
	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	// Retrieve API secret information from environment variables
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	// Set specific options, others will fall back to default values
	wsOption := types.NewWebSocketClientOptionBuilder().Build()

	// Create a client using the specified options
	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithWebSocketClientOption(wsOption).
		Build()
	client := api.NewClient(option)

	// Get the websocket service
	wsService := client.WsService()

	spotPublicWs := wsService.NewSpotPublicWS()
	err := spotPublicWs.Start()
	if err != nil {
		panic(err)
	}
	defer spotPublicWs.Stop()

	subId, err := spotPublicWs.Klines("BTC-USDT,ETH-USDT", "1min", onKlineData)
	if err != nil {
		panic(err)
	}

	time.Sleep(3 * time.Minute)

	err = spotPublicWs.UnSubscribe(subId)
	if err != nil {
		panic(err)
	}

	printKlineData()

	log.Println("Program finished.")
}
