package gemini

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/eightseventhreethree/gemini-go2rest/pkg/handlers"
)

type ClientOpts struct {
	ApiKey     string
	ApiSecret  string
	BaseURL    string
	Timeout    time.Duration
	RetryLimit int
	RetryDelay time.Duration
}

type Client interface {
	GetSymbols() (*SymbolResponse, error)
	GetSymbolDetails(symbol *SymbolRequest) (*SymbolDetailResponse, error)
	GetTickerV1(symbol *SymbolRequest) (*TickerV1Response, error)
	GetTickerV2(symbol *SymbolRequest) (*TickerV2Response, error)
	GetCandles(candles *CandlesRequest) (*CandlesResponse, error)
	GetCurrentOrderBook(order *OrderBookRequest) (*OrderBookResponse, error)
	GetTradeHistory(order *OrderBookRequest) (*OrderBookResponse, error)
}

type client struct {
	ApiKey     string
	ApiSecret  string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(opts *ClientOpts) Client {
	return &client{
		BaseURL:   opts.BaseURL,
		ApiKey:    opts.ApiKey,
		ApiSecret: opts.ApiSecret,
		HTTPClient: &http.Client{
			Timeout: opts.Timeout,
		},
	}
}

func (c *client) httpRequest(req *http.Request) (*http.Response, error) {
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

func (c *client) getUmarshalAndStore(endpoint string, s interface{}) error {

	endpoint = fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	req, err := http.NewRequest("GET", endpoint, nil)
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
func (c *client) GetSymbols() (*SymbolResponse, error) {
	symbols := SymbolResponse{}
	err := c.getUmarshalAndStore(Endpoints["v1"]["symbols"], &symbols.Symbols)
	handlers.CheckErrLog(err, "failed GetSymbols()")
	symbols.Count = len(symbols.Symbols)
	return &symbols, err
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#symbol-details
func (c *client) GetSymbolDetails(symbol *SymbolRequest) (*SymbolDetailResponse, error) {
	e := fmt.Sprintf(Endpoints["v1"]["symbol-details"], symbol.Name)
	symbolDetails := SymbolDetailResponse{}
	err := c.getUmarshalAndStore(e, &symbolDetails)
	handlers.CheckErrLog(err, "failed GetSymbolDetails()")
	return &symbolDetails, err
}

// GetTickerV1 https://docs.gemini.com/rest-api/#ticker
func (c *client) GetTickerV1(symbol *SymbolRequest) (*TickerV1Response, error) {
	e := fmt.Sprintf(Endpoints["v1"]["ticker"], symbol.Name)
	tickerResp := TickerV1Response{}
	err := c.getUmarshalAndStore(e, &tickerResp)
	handlers.CheckErrLog(err, "failed GetTickerV1()")
	return &tickerResp, err
}

// GetTickerV2 https://docs.gemini.com/rest-api/#ticker-v2
func (c *client) GetTickerV2(symbol *SymbolRequest) (*TickerV2Response, error) {
	e := fmt.Sprintf(Endpoints["v2"]["ticker"], symbol.Name)
	tickerResp := TickerV2Response{}
	err := c.getUmarshalAndStore(e, &tickerResp)
	handlers.CheckErrLog(err, "failed GetTickerV2()")
	return &tickerResp, err
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#candles
func (c *client) GetCandles(candles *CandlesRequest) (*CandlesResponse, error) {
	e := fmt.Sprintf(Endpoints["v2"]["candles"], candles.Symbol, candles.TimeFrame)
	candlesResp := CandlesResponse{}
	err := c.getUmarshalAndStore(e, &candlesResp.Candles)
	handlers.CheckErrLog(err, "failed GetCandles()")
	return &candlesResp, err
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#current-order-book
func (c *client) GetCurrentOrderBook(order *OrderBookRequest) (*OrderBookResponse, error) {
	e := fmt.Sprintf(Endpoints["v1"]["current-order-book"], order.Symbol)
	e = fmt.Sprintf("%s?limit_asks=%v&limit_bids=%v", e, order.LimitAsks, order.LimitBids)
	orderBookResp := OrderBookResponse{}
	err := c.getUmarshalAndStore(e, &orderBookResp)
	handlers.CheckErrLog(err, "failed GetCurrentOrderBook()")
	return &orderBookResp, err
}

// GetTradeHistory https://docs.gemini.com/rest-api/#trade-history
func (c *client) GetTradeHistory(order *OrderBookRequest) (*OrderBookResponse, error) {
	e := fmt.Sprintf(Endpoints["v1"]["current-order-book"], order.Symbol)
	e = fmt.Sprintf("%s?limit_asks=%v&limit_bids=%v", e, order.LimitAsks, order.LimitBids)
	orderBookResp := OrderBookResponse{}
	err := c.getUmarshalAndStore(e, &orderBookResp)
	handlers.CheckErrLog(err, "failed GetTradeHistory()")
	return &orderBookResp, err
}
