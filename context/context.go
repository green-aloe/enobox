package context

import (
	"context"
	"sync"
)

// Context holds the context for a single sample of audio.
type Context struct {
	context.Context

	time       Time
	sampleRate int
}

// A Decorator modifies a context.
type Decorator func(Context) context.Context

var (
	decorators      []Decorator
	decoratorsMutex sync.RWMutex
)

// AddDecorator adds a decorator to the list of decorators. These are run in the order they are
// added when setting up a new context. This allows for adding additional data to a context as it's
// created.
func AddDecorator(decorator Decorator) {
	if decorator == nil {
		return
	}

	decoratorsMutex.Lock()
	defer decoratorsMutex.Unlock()

	decorators = append(decorators, decorator)
}

// NewContext sets up and returns a context for a single sample of audio.
func NewContext() Context {
	ctx := Context{
		Context:    context.Background(),
		time:       NewTime(),
		sampleRate: SampleRate(),
	}

	if len(decorators) > 0 {
		decoratorsMutex.RLock()
		defer decoratorsMutex.RUnlock()

		for _, decorator := range decorators {
			if output := decorator(ctx); output != nil {
				if ectx, ok := output.(Context); ok {
					ctx = ectx
				} else {
					ctx.Context = output
				}
			}
		}
	}

	return ctx
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
