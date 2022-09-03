package wanip

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

var _ = NewAgent

// Agent 执行机器人
type Agent struct {
	client  *resty.Client
	logger  simaqian.Logger
	options *options
}

// NewAgent 创建执行机器人
func NewAgent(opts ...option) (agent *Agent, err error) {
	agent = new(Agent)
	agent.options = defaultOptions()
	for _, opt := range opts {
		opt.apply(agent.options)
	}
	if nil == agent.options.logger {
		agent.logger, err = simaqian.New()
	} else {
		agent.logger = agent.options.logger
	}
	if nil == agent.options.client {
		agent.client = resty.New()
	} else {
		agent.client = agent.options.client
	}

	return
}

func (a *Agent) Get(opts ...getOption) (ip string, err error) {
	_options := defaultGetOptions()
	for _, opt := range opts {
		opt.applyGet(_options)
	}

	switch _options.mode {
	case modeIpv4:
		ip, err = a.options.provider.ipv4.ipv4(a.client)
	}

	return
}
