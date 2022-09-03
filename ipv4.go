package wanip

import (
	"github.com/go-resty/resty/v2"
)

type ipv4 interface {
	ipv4(client *resty.Client) (ip string, err error)
}
