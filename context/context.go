package context

import (
	gocontext "context"
	"sync"
	"testing"
)

// A Context holds the context for a single sample of audio.
type Context interface {
	gocontext.Context
	WithValue(key, value any) Context
	Time() Time
	SetTime(time Time)
	SampleRate() int
	NyqistFrequency() float32
}

// A Decorator modifies a context.
type Decorator func(Context) Context

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

// context is the internal implementation of Context.
type context struct {
	gocontext.Context

	time       Time
	sampleRate int
}

// NewContext sets up and returns a context for a single sample of audio using default/global values.
func NewContext() Context {
	return NewContextWith(ContextOptions{})
}

// ContextOptions is the set of configurations that can be used when building a new custom context.
type ContextOptions struct {
	Context    gocontext.Context
	Time       Time
	SampleRate int
	Decorators []Decorator
}

// NewContextWith sets up and returns a context for a single sample of audio using the options
// provided. If a necessary option is not set, the default/global value is used. Any decorators
// provided are run after the global decorators set with AddDecorator.
func NewContextWith(options ContextOptions) Context {
	if options.Context == nil {
		options.Context = gocontext.Background()
	}
	if options.SampleRate <= 0 {
		options.SampleRate = SampleRate()
	}
	if options.Time == (Time{}) {
		options.Time = NewTimeWith(options.SampleRate)
	}

	var ctx Context = &context{
		Context:    options.Context,
		time:       options.Time,
		sampleRate: options.SampleRate,
	}

	for _, decorators := range [][]Decorator{decorators, options.Decorators} {
		for _, decorator := range decorators {
			if decorator != nil {
				if output := decorator(ctx); output != nil {
					ctx = output
				}
			}
		}
	}

	return ctx
}

// NewTestContext returns an empty, non-nil context that does not have any default values set. It
// returns nil if not called from a test. This is meant for testing purposes only.
func NewTestContext() Context {
	if !testing.Testing() {
		return nil
	}

	return &context{
		Context: gocontext.Background(),
	}
}

// WithValue sets an arbitrary value in the context under the provided key and returns the context.
// The key should follow the same general guidelines for the standard library's context.WithValue.
func (ctx *context) WithValue(key, value any) Context {
	if ctx == nil || ctx.Context == nil {
		return ctx
	}

	ctx.Context = gocontext.WithValue(ctx.Context, key, value)

	return ctx
}

// Time returns the context's internal timestamp.
func (ctx *context) Time() Time {
	if ctx == nil {
		return Time{}
	}

	return ctx.time
}

// SetTime sets the context's internal timestamp.
func (ctx *context) SetTime(time Time) {
	if ctx == nil {
		return
	}

	ctx.time = time
}

// SampleRate returns the sample rate for this context.
func (ctx *context) SampleRate() int {
	if ctx == nil {
		return 0
	}

	return ctx.sampleRate
}

// NyqistFrequency returns the maximum frequency that should be used with this context's sample rate.
func (ctx *context) NyqistFrequency() float32 {
	return float32(ctx.SampleRate() / 2)
}
