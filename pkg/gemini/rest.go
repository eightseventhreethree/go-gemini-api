package gemini

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/eightseventhreethree/gemini-go2rest/pkg/handlers"
	"github.com/go-resty/resty/v2"
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
	GetTickerV1(symbol *SymbolRequest) (*resty.Response, error)
	GetTickerV2(symbol *SymbolRequest) (*resty.Response, error)
	GetCandles(candles *CandlesRequest) (*resty.Response, error)
	GetCurrentOrderBook(order *OrderBookRequest) (*resty.Response, error)
}

type client struct {
	*resty.Client
}

func NewClient(opts *ClientOpts) Client {
	c := resty.New()
	c.SetBaseURL(opts.BaseURL)
	c.SetRetryCount(opts.RetryLimit)
	c.SetRetryMaxWaitTime(opts.Timeout)
	c.SetRetryWaitTime(opts.RetryDelay)
	return &client{c}
}

func (c *client) getUmarshalAndStore(endpoint string, s interface{}) error {
	resp, err := c.R().Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(resp.Body(), &s); err != nil {
		return err
	}
	return err
}

// GetSymbols https://docs.gemini.com/rest-api/#symbols
func (c *client) GetSymbols() (*SymbolResponse, error) {
	symbols := SymbolResponse{}
	err := c.getUmarshalAndStore(Endpoints["v1"]["symbols"], &symbols.Symbols)
	handlers.CheckErrLog(err, "failed GetSymbols()")
	symbols.Count = len(symbols.Symbols)
	return &symbols, nil
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#symbol-details
func (c *client) GetSymbolDetails(symbol *SymbolRequest) (*SymbolDetailResponse, error) {
	e := fmt.Sprintf(Endpoints["v1"]["symbol-details"], symbol.Name)
	symbolDetails := SymbolDetailResponse{}
	err := c.getUmarshalAndStore(e, &symbolDetails)
	handlers.CheckErrLog(err, "failed GetSymbolDetails()")
	return &symbolDetails, nil
}

// GetTickerV1 https://docs.gemini.com/rest-api/#ticker
func (c *client) GetTickerV1(symbol *SymbolRequest) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v1"]["ticker"], symbol.Name)
	return c.R().Get(e)
}

// GetTickerV2 https://docs.gemini.com/rest-api/#ticker-v2
func (c *client) GetTickerV2(symbol *SymbolRequest) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v2"]["ticker"], symbol.Name)
	return c.R().Get(e)
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#candles
func (c *client) GetCandles(candles *CandlesRequest) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v2"]["candles"], candles.Symbol, candles.TimeFrame)
	return c.R().Get(e)
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#current-order-book
func (c *client) GetCurrentOrderBook(order *OrderBookRequest) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v1"]["current-order-book"], order.Symbol)
	return c.R().SetQueryParams(map[string]string{
		"limit_asks": fmt.Sprintf("%v", order.LimitAsks),
		"limit_bids": fmt.Sprintf("%v", order.LimitBids),
	}).Get(e)
}

// GetTradeHistory https://docs.gemini.com/rest-api/#trade-history
func (c *client) GetTradeHistory(order *OrderBookRequest) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v1"]["current-order-book"], order.Symbol)
	return c.R().SetQueryParams(map[string]string{
		"limit_asks": fmt.Sprintf("%v", order.LimitAsks),
		"limit_bids": fmt.Sprintf("%v", order.LimitBids),
	}).Get(e)
}
