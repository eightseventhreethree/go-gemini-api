package publicapi

import (
	"fmt"

	"github.com/eightseventhreethree/gemini-go2rest/pkg/gemini"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	*gemini.Client
}

// GetSymbols https://docs.gemini.com/rest-api/#symbols
func (c *Client) GetSymbols() (*resty.Response, error) {
	return c.R().Get(Endpoints["v1"]["symbols"])
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#symbol-details
func (c *Client) GetSymbolDetails(symbol string) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v1"]["symbol-details"], symbol)
	return c.R().Get(e)
}

// GetTickerV1 https://docs.gemini.com/rest-api/#ticker
func (c *Client) GetTickerV1(symbol string) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v1"]["ticker"], symbol)
	return c.R().Get(e)
}

// GetTickerV2 https://docs.gemini.com/rest-api/#ticker-v2
func (c *Client) GetTickerV2(symbol string) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v2"]["ticker"], symbol)
	return c.R().Get(e)
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#candles
func (c *Client) GetCandles(symbol string, time_frame string) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v2"]["candles"], symbol, time_frame)
	return c.R().Get(e)
}

// GetSymbolDetails https://docs.gemini.com/rest-api/#current-order-book
func (c *Client) GetCurrentOrderBook(symbol string, optionalquery string) (*resty.Response, error) {
	e := fmt.Sprintf(Endpoints["v1"]["current-order-book"], symbol)
	if optionalquery != "" {
		// i.e limit_bids=50&limit_asks=50
		return c.R().SetQueryString(optionalquery).Get(e)
	} else {
		return c.R().Get(e)
	}

}
