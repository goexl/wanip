package wanip

import (
	"github.com/goexl/simaqian"
)

var (
	_        = Logger
	_ option = (*optionLogger)(nil)
)

type optionLogger struct {
	logger simaqian.Logger
}

// Logger 客户端
func Logger(logger simaqian.Logger) *optionLogger {
	return &optionLogger{
		logger: logger,
	}
}

func (l *optionLogger) apply(options *options) {
	options.logger = l.logger
}
