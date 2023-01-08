package gemini

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/eightseventhreethree/go-gemini-api/pkg/handlers"
	"github.com/hashicorp/go-retryablehttp"
)

type ClientOpts struct {
	ApiKey     string // unused atm
	ApiSecret  string // unused atm
	BaseURL    string
	Timeout    time.Duration
	RetryLimit int
	RetryDelay time.Duration
}

type Client interface {
	GetSymbols() (SymbolResponse, error)
	GetSymbolDetails(symbol *SymbolRequest) (SymbolDetailResponse, error)
	GetTickerV1(symbol *SymbolRequest) (TickerV1Response, error)
	GetTickerV2(symbol *SymbolRequest) (TickerV2Response, error)
	GetCandles(candles *CandlesRequest) (CandlesResponse, error)
	GetCurrentOrderBook(order *OrderBookRequest) (OrderBookResponse, error)
	GetTradeHistory(history *TradeHistoryRequest) (TradeHistoryResponse, error)
	GetCurrentAuction(symbol *SymbolRequest) (CurrentAuctionResponse, error)
	GetPriceFeed() ([]PriceFeedResponse, error)
}

type client struct {
	ApiKey     string
	ApiSecret  string
	BaseURL    string
	HTTPClient *retryablehttp.Client
}

func NewClient(opts *ClientOpts) Client {
	c := retryablehttp.NewClient()
	c.RetryMax = opts.RetryLimit
	c.RetryWaitMin = opts.RetryDelay
	return &client{
		BaseURL:    opts.BaseURL,
		ApiKey:     opts.ApiKey,    // unused atm
		ApiSecret:  opts.ApiSecret, // unused atm
		HTTPClient: c,
	}
}

func (c *client) httpRequest(req *retryablehttp.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return res, err
}

func (c *client) getUmarshalAndStore(u string, s interface{}) error {
	endpoint := fmt.Sprintf("%s%s", c.BaseURL, u)
	req, err := retryablehttp.NewRequest("GET", endpoint, nil)
	handlers.CheckErrLog(err, "Failed to create NewRequest")

	resp, err := c.httpRequest(req)
	handlers.CheckErrPanic(err)
	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	handlers.CheckErrLog(err, "failed to read resp.Body")

	if err := json.Unmarshal(responseBytes, &s); err != nil {
		return err
	}
	handlers.CheckErrLog(err, "failed to Unmarshal responseBytes")

	return err
}

/*
GetSymbols https://docs.gemini.com/rest-api/#symbols
Returns []string of symbols and count of symbols
*/
func (c *client) GetSymbols() (SymbolResponse, error) {
	symbols := SymbolResponse{}
	err := c.getUmarshalAndStore(V1Symbols.toString(), &symbols.Symbols)
	handlers.CheckErrLog(err, "failed GetSymbols()")
	symbols.Count = len(symbols.Symbols)
	return symbols, err
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#symbol-details
func (c *client) GetSymbolDetails(symbol *SymbolRequest) (SymbolDetailResponse, error) {
	e := fmt.Sprintf(V1SymbolsDetails, symbol.Name)
	symbolDetails := SymbolDetailResponse{}
	err := c.getUmarshalAndStore(e, &symbolDetails)
	handlers.CheckErrLog(err, "failed GetSymbolDetails()")
	return symbolDetails, err
}

// GetTickerV1 https://docs.gemini.com/rest-api/#ticker
func (c *client) GetTickerV1(symbol *SymbolRequest) (TickerV1Response, error) {
	e := fmt.Sprintf(V1Ticker, symbol.Name)
	tickerResp := TickerV1Response{}
	err := c.getUmarshalAndStore(e, &tickerResp)
	handlers.CheckErrLog(err, "failed GetTickerV1()")
	return tickerResp, err
}

// GetTickerV2 https://docs.gemini.com/rest-api/#ticker-v2
func (c *client) GetTickerV2(symbol *SymbolRequest) (TickerV2Response, error) {
	e := fmt.Sprintf(V2Ticker, symbol.Name)
	tickerResp := TickerV2Response{}
	err := c.getUmarshalAndStore(e, &tickerResp)
	handlers.CheckErrLog(err, "failed GetTickerV2()")
	return tickerResp, err
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#candles
func (c *client) GetCandles(candles *CandlesRequest) (CandlesResponse, error) {
	e := fmt.Sprintf(V2Candles, candles.Symbol, candles.TimeFrame)
	candlesResp := CandlesResponse{}
	err := c.getUmarshalAndStore(e, &candlesResp.Candles)
	handlers.CheckErrLog(err, "failed GetCandles()")
	return candlesResp, err
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#current-order-book
func (c *client) GetCurrentOrderBook(order *OrderBookRequest) (OrderBookResponse, error) {
	e := fmt.Sprintf(V1CurrentOrderBook, order.Symbol)
	e = fmt.Sprintf("%s?limit_asks=%v&limit_bids=%v", e, order.LimitAsks, order.LimitBids)
	orderBookResp := OrderBookResponse{}
	err := c.getUmarshalAndStore(e, &orderBookResp)
	handlers.CheckErrLog(err, "failed GetCurrentOrderBook()")
	return orderBookResp, err
}

// GetTradeHistory https://docs.gemini.com/rest-api/#trade-history
func (c *client) GetTradeHistory(history *TradeHistoryRequest) (TradeHistoryResponse, error) {
	e := fmt.Sprintf(V1TradeHistory, history.Symbol)
	historyResp := TradeHistoryResponse{}
	err := c.getUmarshalAndStore(e, &historyResp)
	handlers.CheckErrLog(err, "failed GetTradeHistory()")
	return historyResp, err
}

// GetCurrentAuction https://docs.gemini.com/rest-api/#current-auction
func (c *client) GetCurrentAuction(symbol *SymbolRequest) (CurrentAuctionResponse, error) {
	e := fmt.Sprintf(V1CurrentAuction, symbol.Name)
	currentAuctionResp := CurrentAuctionResponse{}
	err := c.getUmarshalAndStore(e, &currentAuctionResp)
	handlers.CheckErrLog(err, "failed GetCurrentAuction()")
	return currentAuctionResp, err
}

func (c *client) GetPriceFeed() ([]PriceFeedResponse, error) {
	pricefeedResp := []PriceFeedResponse{}
	err := c.getUmarshalAndStore(V1PriceFeed, &pricefeedResp)
	handlers.CheckErrLog(err, "failed GetPriceFeed()")
	return pricefeedResp, err
}
