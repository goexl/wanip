package wanip

import (
	"github.com/go-resty/resty/v2"
)

var (
	_        = Client
	_ option = (*optionClient)(nil)
)

type optionClient struct {
	client *resty.Client
}

// Client 客户端
func Client(client *resty.Client) *optionClient {
	return &optionClient{
		client: client,
	}
}

func (c *optionClient) apply(options *options) {
	options.client = c.client
}
