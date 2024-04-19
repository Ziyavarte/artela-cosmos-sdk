package log

import (
	sdklog "cosmossdk.io/log"
	"github.com/cometbft/cometbft/libs/log"
)

type sdkLogWrapper struct {
	log.Logger
}

var _ sdklog.Logger = (*sdkLogWrapper)(nil)

// Info takes a message and a set of key/value pairs and logs with level INFO.
// The key of the tuple must be a string.
func (l sdkLogWrapper) Warn(msg string, keyVals ...interface{}) {
	l.Error(msg, keyVals...)
}

// With returns a new wrapped logger with additional context provided by a set.
func (l sdkLogWrapper) With(keyVals ...any) sdklog.Logger {
	return sdkLogWrapper{l.Logger.With(keyVals...)}
}

// Impl returns the underlying zerolog logger.
// It can be used to used zerolog structured API directly instead of the wrapper.
func (l sdkLogWrapper) Impl() interface{} {
	return l.Logger
}

func WrapSDKLogger(logger log.Logger) sdklog.Logger {
	return sdkLogWrapper{logger}
}
