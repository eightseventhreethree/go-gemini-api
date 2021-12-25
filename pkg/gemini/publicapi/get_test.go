package publicapi

import (
	"fmt"
	"testing"

	"github.com/eightseventhreethree/gemini-go2rest/pkg/gemini"
)

func TestGetSymbols(t *testing.T) {
	client := Client{Client: gemini.New(&gemini.ClientOpts{
		BaseURL: "https://api.sandbox.gemini.com",
	})}
	resp, err := client.GetSymbols()
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode() != 200 {
		t.Error("Failed to get 200")
	}
	fmt.Println(resp)
}

func TestGetSymbolDetails(t *testing.T) {
	client := Client{Client: gemini.New(&gemini.ClientOpts{
		BaseURL: "https://api.sandbox.gemini.com",
	})}
	resp, err := client.GetSymbolDetails("btcusd")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode() != 200 {
		t.Error("Failed to get 200")
	}
	fmt.Println(resp)
}

func TestGetTickerV1(t *testing.T) {
	client := Client{Client: gemini.New(&gemini.ClientOpts{
		BaseURL: "https://api.sandbox.gemini.com",
	})}
	resp, err := client.GetTickerV1("btcusd")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode() != 200 {
		t.Error("Failed to get 200")
	}
	fmt.Println(resp)
}

func TestGetCandles(t *testing.T) {
	client := Client{Client: gemini.New(&gemini.ClientOpts{
		BaseURL: "https://api.sandbox.gemini.com",
	})}
	resp, err := client.GetCandles("btcusd", "1day")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode() != 200 {
		t.Error("Failed to get 200")
	}
	//fmt.Println(resp)
}

func TestGetCurrentOrderBook(t *testing.T) {
	client := Client{Client: gemini.New(&gemini.ClientOpts{
		BaseURL: "https://api.sandbox.gemini.com",
	})}
	resp, err := client.GetCurrentOrderBook("btcusd", "")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode() != 200 {
		t.Error("Failed to get 200")
	}
	fmt.Println(resp)
}
