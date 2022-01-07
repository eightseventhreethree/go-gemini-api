package main

import (
	"fmt"
	"time"

	"github.com/eightseventhreethree/go-gemini-api/pkg/gemini"
	"github.com/eightseventhreethree/go-gemini-api/pkg/handlers"
)

func main() {
	api := gemini.NewClient(&gemini.ClientOpts{
		BaseURL:    "https://api.sandbox.gemini.com",
		Timeout:    time.Second.Round(10),
		RetryLimit: 3,
		RetryDelay: time.Second.Round(3),
		//ApiKey:     "xxxxxx",
		//ApiSecret:  "xxxxx",
	})

	resp, err := api.GetSymbols()
	handlers.CheckErrLog(err, "failed to GetSymbols")
	for _, v := range resp.Symbols {
		sym := &gemini.SymbolRequest{Name: v}
		symDetails, err := api.GetSymbolDetails(sym)
		handlers.CheckErrLog(err, "failed to call GetSymbolDetails")
		if symDetails.Status != gemini.Closed {
			tickerResp, err := api.GetTickerV2(sym)
			handlers.CheckErrLog(err, "failed to call GetTickerV1")
			tickerv1Resp, _ := api.GetTickerV1(sym)
			fmt.Printf("Symbol: %s -> Changes: %v -> Volume: %v \n", v, tickerResp.Changes, tickerv1Resp.Volume)
		}
	}
}
