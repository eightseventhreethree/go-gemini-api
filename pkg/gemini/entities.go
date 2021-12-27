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
