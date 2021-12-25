package gemini

import (
	"time"

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

type Client struct {
	*resty.Client
}

func New(opts *ClientOpts) *Client {
	client := resty.New()
	client.SetBaseURL(opts.BaseURL)
	client.SetRetryCount(opts.RetryLimit)
	client.SetRetryMaxWaitTime(opts.Timeout)
	client.SetRetryWaitTime(opts.RetryDelay)
	return &Client{client}
}
