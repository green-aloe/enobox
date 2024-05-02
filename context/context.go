package context

import (
	"context"
)

// Context holds the context for a single sample of audio.
type Context struct {
	context.Context

	time       Time
	sampleRate int
}

// NewContext sets up and returns a context for a single sample of audio.
func NewContext() Context {
	return Context{
		Context:    context.Background(),
		time:       NewTime(),
		sampleRate: SampleRate(),
	}
}

// Time returns the context's internal timestamp.
func (ctx *Context) Time() Time {
	if ctx == nil {
		return Time{}
	}

	return ctx.time
}

// SetTime sets the context's internal timestamp.
func (ctx *Context) SetTime(time Time) {
	if ctx == nil {
		return
	}

	ctx.time = time
}

// SampleRate returns the sample rate for this context.
func (ctx *Context) SampleRate() int {
	if ctx == nil {
		return 0
	}

	return ctx.sampleRate
}

// SetSampleRate sets the sample rate for this context.
func (ctx *Context) SetSampleRate(rate int) {
	if ctx == nil {
		return
	}

	ctx.sampleRate = rate
}

// NyqistFrequency returns the maximum frequency that should be used with this context's sample rate.
func (ctx *Context) NyqistFrequency() float32 {
	return float32(ctx.SampleRate() / 2)
}
