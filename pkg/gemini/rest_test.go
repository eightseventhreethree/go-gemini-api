package gemini

import (
	"fmt"
	"testing"
)

// global client for test
var CLIENT = NewClient(&ClientOpts{
	BaseURL: "https://api.sandbox.gemini.com",
})

func TestGetSymbols(t *testing.T) {
	resp, err := CLIENT.GetSymbols()
	if err != nil {
		t.Error(err)
	}
	if len(resp.Symbols) != resp.Count {
		t.Errorf("Failed to return sufficient symbols: %v error: %ss", resp.Symbols, err)
	}
	fmt.Printf("%#v", resp)
}

func TestGetSymbolDetails(t *testing.T) {
	symbol := "BTCUSD"
	resp, err := CLIENT.GetSymbolDetails(&SymbolRequest{Name: symbol})
	if err != nil {
		t.Error(err)
	}
	if resp.Symbol != symbol {
		t.Errorf("Failed to find %s in resp", symbol)
	}
	fmt.Printf("%#v", resp)
}

func TestGetTickerV1(t *testing.T) {
	resp, err := CLIENT.GetTickerV1(&SymbolRequest{Name: "BTCUSD"})
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode() != 200 {
		t.Error("Failed to get 200")
	}
	fmt.Println(resp)
}

func TestGetCandles(t *testing.T) {
	resp, err := CLIENT.GetCandles(&CandlesRequest{Symbol: "BTCUSD", TimeFrame: Min1})
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode() != 200 {
		t.Error("Failed to get 200")
	}
	//fmt.Println(resp)
}

func TestGetCurrentOrderBook(t *testing.T) {
	resp, err := CLIENT.GetCurrentOrderBook(&OrderBookRequest{Symbol: "btcusd", LimitBids: 10, LimitAsks: 10})
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode() != 200 {
		t.Error("Failed to get 200")
	}
	fmt.Println(resp)
}
