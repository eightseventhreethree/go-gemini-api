package gemini

import (
	"fmt"
	"testing"

	"github.com/eightseventhreethree/gemini-go2rest/pkg/handlers"
	"github.com/stretchr/testify/require"
)

// global client for test
var CLIENT = NewClient(&ClientOpts{
	BaseURL: "https://api.sandbox.gemini.com",
})

// global symbol
const SYMBOL = "BTCUSD"

func TestGetSymbols(t *testing.T) {
	resp, err := CLIENT.GetSymbols()
	handlers.CheckErrLogT(t, err, "TestGetSymbols")
	count := len(resp.Symbols)
	require.Equal(t, resp.Count, count, "count should be equal")
	if err != nil {
		fmt.Printf("%+v\n", resp)
	}
}

func TestGetSymbolDetails(t *testing.T) {
	resp, err := CLIENT.GetSymbolDetails(&SymbolRequest{Name: SYMBOL})
	handlers.CheckErrLogT(t, err, "TestGetSymbolDetails")
	require.Equal(t, resp.Symbol, SYMBOL, "these symbols should be equal")
	if err != nil {
		fmt.Printf("%+v\n", resp)
	}
}

func TestGetTickerV1(t *testing.T) {
	resp, err := CLIENT.GetTickerV1(&SymbolRequest{Name: SYMBOL})
	handlers.CheckErrLogT(t, err, "TestGetTickerV1")
	require.NotZero(t, resp.Volume)
	if err != nil {
		fmt.Printf("%+v\n", resp)
	}
}

func TestGetTickerV2(t *testing.T) {
	resp, err := CLIENT.GetTickerV2(&SymbolRequest{Name: SYMBOL})
	handlers.CheckErrLogT(t, err, "TestGetTickerV2")
	require.Equal(t, resp.Symbol, SYMBOL, "these symbols should be equal")
	if err != nil {
		fmt.Printf("%+v\n", resp)
	}
}

func TestGetCandles(t *testing.T) {
	resp, err := CLIENT.GetCandles(&CandlesRequest{Symbol: SYMBOL, TimeFrame: Min1})
	handlers.CheckErrLogT(t, err, "TestGetCandles")
	if err != nil {
		fmt.Printf("%+v\n", resp)
	}
}

func TestGetCurrentOrderBook(t *testing.T) {
	resp, err := CLIENT.GetCurrentOrderBook(&OrderBookRequest{Symbol: SYMBOL, LimitBids: 10, LimitAsks: 10})
	handlers.CheckErrLogT(t, err, "TestGetCurrentOrderBook")
	if err != nil {
		fmt.Printf("%+v\n", resp)
	}
}
