package context

import (
	"context"
)

// These types each represent one key for storing a value in a context.
//
// They're all empty structs to avoid any allocations. Because one empty struct is indistinguishable
// from another, we have to use a different type for each key.
type contextKeyTest struct{}
type contextKeyTime struct{}
type contextKeySampleRate struct{}

// Context holds the context for a single sample of audio.
type Context struct {
	context.Context
}

// NewContext sets up and returns a context for one sample in time.
func NewContext() Context {
	ctx := context.Background()

	// Initialize some default values.
	ctx = context.WithValue(ctx, contextKeyTime{}, NewTime())
	ctx = context.WithValue(ctx, contextKeySampleRate{}, SampleRate())

	return Context{
		Context: ctx,
	}
}

// value returns a hard-typed value from the context. If the value is not found, it returns the zero
// value of the type.
func value[T any](ctx *Context, key any) T {
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
	return value[Time](ctx, contextKeyTime{})
}

// SetTime sets the context's internal timestamp.
func (ctx *Context) SetTime(time Time) {
	if ctx == nil || ctx.Context == nil {
		return
	}

	ctx.Context = context.WithValue(ctx, contextKeyTime{}, time)
}

// SampleRate returns the sample rate for this context.
func (ctx *Context) SampleRate() int {
	return value[int](ctx, contextKeySampleRate{})
}

// SetSampleRate sets the sample rate for this context.
func (ctx *Context) SetSampleRate(rate int) {
	if ctx == nil || ctx.Context == nil {
		return
	}

	ctx.Context = context.WithValue(ctx, contextKeySampleRate{}, rate)
}

// NyqistFrequency returns the maximum frequency that should be used with this context's sample rate.
func (ctx *Context) NyqistFrequency() float32 {
	return float32(ctx.SampleRate() / 2)
}
