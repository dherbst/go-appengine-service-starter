package starter

import (
	"golang.org/x/net/context"
)

const (
	configKey = "starter config"
)

// NewContextWithConfig creates a context with the config and returns the wrapped context.
func NewContextWithConfig(c Config) context.Context {
	return context.WithValue(context.Background(), configKey, c)
}

// ConfigFromContext returns the config struct from the wrapping context.
func ConfigFromContext(ctx context.Context) (Config, bool) {
	config, ok := ctx.Value(configKey).(Config)
	return config, ok
}
