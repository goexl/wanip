package wanip

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

var _ = newAkamai

type akamai struct{}

func newAkamai() *akamai {
	return new(akamai)
}

func (c *akamai) ipv4(client *resty.Client) (ip string, err error) {
	if rsp, getErr := client.R().Get(`https://whatismyip.akamai.com`); nil != getErr {
		err = getErr
	} else if nil != rsp && rsp.IsSuccess() {
		ip = rsp.String()
	} else {
		err = exc.NewFields(`获取外网地址出错`, field.Int(`code`, rsp.StatusCode()))
	}

	return
}
