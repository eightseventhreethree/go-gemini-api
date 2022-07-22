package gemini

import "fmt"

type Endpoints string

const (
	V1Symbols          Endpoints = "/v1/symbols"
	V1SymbolsDetails             = "/v1/symbols/details/%s"
	V1Ticker                     = "/v1/pubticker/%s"
	V1CurrentOrderBook           = "/v1/book/%s"
	V1TradeHistory               = "/v1/trades/%s"
	V1CurrentAuction             = "/v1/auction/%s"
	V1AuctionHistory             = "/v1/auction/%s/history"
	V1PriceFeed                  = "/v1/pricefeed"
	V2Ticker                     = "/v2/ticker/%s"
	V2Candles                    = "/v2/candles/%s/%s"
)

func (e Endpoints) toString() string {
	return fmt.Sprintf("%s", e)
}
