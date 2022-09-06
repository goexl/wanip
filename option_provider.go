package wanip

var (
	_        = Ipv4
	_        = Ipv6
	_ option = (*optionProvider)(nil)
)

type optionProvider struct {
	ipv4 ipv4
	ipv6 ipv6
}

// Ipv4 Ipv4提供商
func Ipv4(ipv4 ipv4) *optionProvider {
	return &optionProvider{
		ipv4: ipv4,
	}
}

// Ipv6 Ipv6提供商
func Ipv6(ipv6 ipv6) *optionProvider {
	return &optionProvider{
		ipv6: ipv6,
	}
}

func (c *optionProvider) apply(options *options) {
	if nil != c.ipv4 {
		options.provider.ipv4 = c.ipv4
	}

	if nil != c.ipv6 {
		options.provider.ipv6 = c.ipv6
	}
}
