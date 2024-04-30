package context

import (
	"context"
)

const (
	keyTime       ctxKey = "time"
	keySampleRate ctxKey = "sampleRate"
)

type ctxKey string

// Context holds the context for a single sample of audio.
type Context struct {
	context.Context
}

// NewContext sets up and returns a context for one sample in time.
func NewContext() Context {
	ctx := context.Background()

	// Initialize some default values.
	ctx = context.WithValue(ctx, keyTime, NewTime())
	ctx = context.WithValue(ctx, keySampleRate, SampleRate())

	return Context{
		Context: ctx,
	}
}

func value[T any](ctx *Context, key ctxKey) T {
	if ctx != nil && ctx.Context != nil {
		if v := ctx.Value(key); v != nil {
			if t, ok := v.(T); ok {
				return t
			}
		}
	}

	var t T
	return t
}

// Time returns the context's internal timestamp.
func (ctx *Context) Time() Time {
	return value[Time](ctx, keyTime)
}

// SampleRate returns the sample rate for this context.
func (ctx *Context) SampleRate() int {
	return value[int](ctx, keySampleRate)
}

// NyqistFrequency returns the maximum frequency that should be used with this context's sample rate.
func (ctx *Context) NyqistFrequency() float32 {
	return float32(ctx.SampleRate() / 2)
}
