package limiter

import "time"

// Options keeps the settings to set up redis limiter.
type Options func(*RateLimiterService)

// WithGlobalLimiterCustom is a function to set global value to the RouterOption.
func WithGlobalLimiterCustom(globalLimit int, globalDuration time.Duration) Options {
	return func(r *RateLimiterService) {
		r.globalLimit = globalLimit
		r.globalDuration = globalDuration
	}
}

// WithIdentifierLimiterCustom is a function to set identifier value to the RouterOption.
func WithIdentifierLimiterCustom(identifierLimit int, identifierDuration time.Duration) Options {
	return func(r *RateLimiterService) {
		r.identifierLimit = identifierLimit
		r.identifierDuration = identifierDuration
	}
}
