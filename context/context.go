package context

// Context holds the context for a single sample of audio.
type Context struct {
	time       Time
	sampleRate int
}

// NewContext sets up and returns a context for one sample in time.
func NewContext() Context {
	return Context{
		time:       NewTime(),
		sampleRate: SampleRate(),
	}
}

// Time returns the context's internal timestamp.
func (ctx Context) Time() Time {
	return ctx.time
}

// SampleRate returns the sample rate for this context.
func (ctx Context) SampleRate() int {
	return ctx.sampleRate
}

// NyqistFrequency returns the maximum frequency that should be used with this sample rate.
func (ctx Context) NyqistFrequency() float32 {
	return float32(ctx.SampleRate() / 2)
}
