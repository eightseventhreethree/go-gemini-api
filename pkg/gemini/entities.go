package gemini

import (
	"github.com/shopspring/decimal"
)

type SymbolRequest struct {
	Name string `json:"name"`
}

type SymbolResponse struct {
	Symbols []string
	Count   int
}

type SymbolDetailResponse struct {
	Symbol         string           `json:"symbol"`
	BaseCurrency   string           `json:"base_currency"`
	QuoteCurrency  string           `json:"quote_currency"`
	TickSize       decimal.Decimal  `json:"tick_size"`
	QuoteIncrement decimal.Decimal  `json:"quote_increment"`
	MinOrderSize   string           `json:"min_order_size"`
	Status         ValidOrderStatus `json:"status"`
	WrapEnabled    bool             `json:"wrap_enabled"`
}

type CurrentAuctionResponse struct {
	ClosedUntilMS               int64           `json:"closed_until_ms"`
	LastAuctionEID              int32           `json:"last_auction_eid"`
	LastAuctionPrice            decimal.Decimal `json:"last_auction_price"`
	LastAuctionQuanity          decimal.Decimal `json:"last_auction_quantity"`
	LastHighestBidPrice         decimal.Decimal `json:"last_highest_bid_price"`
	LastLowestAskPrice          decimal.Decimal `json:"last_lowest_ask_price"`
	LastCollarPrice             decimal.Decimal `json:"last_collar_price"`
	MostRecentIndicativePrice   decimal.Decimal `json:"most_recent_indicative_price"`
	MostRecentIndicitiveQuanity decimal.Decimal `json:"most_recent_indicative_quantity"`
	MostRecentHighestBidPrice   decimal.Decimal `json:"most_recent_highest_bid_price"`
	MostRecentLowestAskPrice    decimal.Decimal `json:"most_recent_lowest_ask_price"`
	MostRecentCollarPrice       decimal.Decimal `json:"most_recent_collar_price"`
	NextUpdateMS                int64           `json:"next_update_ms"`
	NextAuctionMS               int64           `json:"next_auction_ms"`
}

type TickerV1Response struct {
	Bid    decimal.Decimal        `json:"bid"`
	Ask    decimal.Decimal        `json:"ask"`
	Last   decimal.Decimal        `json:"last"`
	Volume map[string]interface{} `json:"volume"`
}

type TickerV2Response struct {
	Symbol  string            `json:"symbol"`
	Open    decimal.Decimal   `json:"open"`
	High    decimal.Decimal   `json:"high"`
	Low     decimal.Decimal   `json:"low"`
	Close   decimal.Decimal   `json:"close"`
	Changes []decimal.Decimal `json:"changes"`
	Bid     decimal.Decimal   `json:"bid"`
	Ask     decimal.Decimal   `json:"ask"`
}

type CandlesRequest struct {
	Symbol    string       `json:"symbol"`
	TimeFrame TimeNotation `json:"time_frame"`
}

type CandlesResponse struct {
	Candles [][]decimal.Decimal
}

type OrderBookRequest struct {
	Symbol    string `json:"symbol"`
	LimitBids int32  `json:"limit_bids"`
	LimitAsks int32  `json:"limit_asks"`
}

type OrderBookResponse struct {
	Bids []OrderBookFields
	Asks []OrderBookFields
}

type PriceFeedResponse struct {
	Pair             string `json:"pair"`
	Price            string `json:"price"`
	PercentChange24h string `json:"percentChange24h"`
}

type OrderBookFields struct {
	Price  decimal.Decimal `json:"price"`
	Amount decimal.Decimal `json:"amount"`
	// timestamp docs note not to use, unimplemented
}

type ValidOrderStatus string

const (
	Open       ValidOrderStatus = "open"
	Closed     ValidOrderStatus = "closed"
	CancelOnly ValidOrderStatus = "cancel_only"
	PostOnly   ValidOrderStatus = "post_only"
	LimitOnly  ValidOrderStatus = "limit_only"
)

type TimeNotation string

const (
	Min1  TimeNotation = "1m"
	Min5  TimeNotation = "5m"
	Min15 TimeNotation = "15m"
	Min30 TimeNotation = "30m"
	Hour1 TimeNotation = "1hr"
	Hour6 TimeNotation = "6hr"
	Day1  TimeNotation = "1day"
)
