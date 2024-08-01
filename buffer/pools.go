package buffer

import (
	"sync"

	"github.com/green-aloe/enobox/context"
	"github.com/green-aloe/enobox/tone"
	"github.com/green-aloe/utilities/pool"
)

var (
	// We must use a custom pool implementation instead of the more standard sync.Pool because the
	// latter can drop buffers at any time. That wouldn't be good for us because we might need to
	// close the buffer before releasing it (so that non-go garbage collectors can also release
	// their references to the memory).
	bufferPools      = make(map[bufferPoolsKey]*pool.Pool[Buffer])
	bufferPoolsMutex sync.Mutex
	bufferPoolKey    bufferPoolCtxKey
)

type bufferPoolsKey struct {
	sampleRate   int
	numHarmGains int
}

type bufferPoolCtxKey struct{}

func init() {
	// Add a context decorator that sets a buffer pool in each new context depending on the sample
	// rate and number of harmonic gains in a tone configured for the context.
	//
	// We need to add this decorator after the decorator that sets the number of harmonic gains,
	// since this one uses the other one. This loose ordering is guaranteed because this decorator
	// calls the package for the other decorator, which means that package will finish its
	// initialization first and add its decorator the context builder before returning back here.
	context.AddDecorator(func(ctx context.Context) context.Context {
		bufferPoolsMutex.Lock()
		defer bufferPoolsMutex.Unlock()

		key := bufferPoolsKey{
			sampleRate:   ctx.SampleRate(),
			numHarmGains: tone.NumHarmGains(ctx),
		}

		bufferPool, ok := bufferPools[key]
		if !ok || bufferPool == nil {
			bufferPool = &pool.Pool[Buffer]{
				NewItem:  func() Buffer { return NewBuffer(ctx) },
				PreStore: func(buffer Buffer) Buffer { buffer.Reset(); return buffer },
			}
			bufferPools[key] = bufferPool
		}

		return ctx.WithValue(bufferPoolKey, bufferPool)
	})
}

// BufferPool returns the buffer pool for this context, or nil if no pool is set.
func BufferPool(ctx context.Context) *pool.Pool[Buffer] {
	if ctx == nil {
		return nil
	}

	if v := ctx.Value(bufferPoolKey); v != nil {
		if p, ok := v.(*pool.Pool[Buffer]); ok {
			return p
		}
	}

	return nil
}
