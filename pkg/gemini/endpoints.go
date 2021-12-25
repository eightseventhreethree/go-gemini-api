package gemini

var Endpoints = map[string]map[string]string{
	"v1": {
		"symbols":            "/v1/symbols",
		"symbol-details":     "/v1/symbols/details/%s",
		"ticker":             "/v1/pubticker/%s",
		"current-order-book": "/v1/book/%s",
		"trade-history":      "/v1/trades/%s",
		"current-auction":    "/v1/auction/%s",
		"auction-history":    "/v1/auction/%s/history",
		"price-feed":         "/v1/pricefeed",
	},
	"v2": {
		"ticker-v2": "/v2/ticker/%s",
		"candles":   "/v2/candles/%s/%s",
	},
}
