package wanip

type (
	getOption interface {
		applyGet(options *getOptions)
	}

	getOptions struct {
		mode mode
	}
)

func defaultGetOptions() *getOptions {
	return &getOptions{
		mode: modeIpv4,
	}
}
