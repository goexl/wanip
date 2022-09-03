package wanip

import (
	"github.com/go-resty/resty/v2"
)

type ipv6 interface {
	ipv6(client *resty.Client) (ip string, err error)
}
