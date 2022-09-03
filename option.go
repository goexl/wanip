package wanip

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		provider *provider
		client   *resty.Client
		logger   simaqian.Logger
	}
)

func defaultOptions() *options {
	return &options{
		provider: &provider{
			ipv4: newClang(),
		},
	}
}
