package buffer

import (
	"sort"
	"sync"
	"testing"

	"github.com/green-aloe/enobox/context"
	"github.com/green-aloe/enobox/tone"
	"github.com/green-aloe/utilities/pool"
	"github.com/stretchr/testify/require"
)

// Test_init tests that a context decorator that uses the configured number of harmonic gains is
// added on package initialization.
func Test_init(t *testing.T) {
	t.Run("default pool", func(t *testing.T) {
		ctx := context.NewContext()

		v := ctx.Value(bufferPoolKey)
		require.NotNil(t, v)

		bufferPool, ok := v.(*pool.Pool[Buffer])
		require.True(t, ok)
		require.NotNil(t, bufferPool)

		buffer := bufferPool.Get()
		require.Len(t, buffer.Tones, context.DefaultSampleRate)
		for i := range buffer.Tones {
			require.Len(t, buffer.Tones[i].HarmonicGains, tone.DefaultNumHarmGains)
		}
	})

	t.Run("multiple pools", func(t *testing.T) {
		defer tone.SetNumHarmGains(tone.DefaultNumHarmGains)

		for i := range 2 {
			for _, config := range []struct {
				sampleRate   int
				numHarmGains int
			}{
				{context.DefaultSampleRate, tone.DefaultNumHarmGains},
				{48_000, 10},
				{96_000, 100},
			} {
				tone.SetNumHarmGains(config.numHarmGains)
				ctx := context.NewContextWith(context.ContextOptions{
					SampleRate: config.sampleRate,
				})

				v := ctx.Value(bufferPoolKey)
				require.NotNil(t, v)

				bufferPool, ok := v.(*pool.Pool[Buffer])
				require.True(t, ok)
				require.NotNil(t, bufferPool)
				require.Equal(t, i, bufferPool.Count())

				buffer := bufferPool.Get()
				require.Len(t, buffer.Tones, config.sampleRate)
				for j := range buffer.Tones {
					require.Len(t, buffer.Tones[j].HarmonicGains, config.numHarmGains)

					if i == 0 {
						buffer.Tones[j].Frequency = float32(i + 1)
						require.NotZero(t, buffer.Tones[j].Frequency)
					} else {
						require.Zero(t, buffer.Tones[j].Frequency)
					}
				}

				bufferPool.Store(buffer)
				require.Equal(t, 1, bufferPool.Count())
			}
		}
	})

	t.Run("concurrency", func(t *testing.T) {
		bufferPools = make(map[bufferPoolsKey]*pool.Pool[Buffer])

		var wg sync.WaitGroup
		for i := range 1_000 {
			wg.Add(1)

			go func() {
				defer wg.Done()

				_ = context.NewContextWith(context.ContextOptions{
					SampleRate: i + 1,
				})
			}()
		}

		wg.Wait()

		require.Len(t, bufferPools, 1_000)
		keys := make([]bufferPoolsKey, 0, 1_000)
		for k, v := range bufferPools {
			keys = append(keys, k)

			require.NotNil(t, v)
			require.Zero(t, v.Count())

			buffer := v.Get()
			require.Len(t, buffer.Tones, k.sampleRate)
			for j := range buffer.Tones {
				require.Len(t, buffer.Tones[j].HarmonicGains, k.numHarmGains)
			}
		}

		sort.Slice(keys, func(i, j int) bool { return keys[i].sampleRate < keys[j].sampleRate })
		require.Len(t, keys, 1_000)
		for i, k := range keys {
			require.Equal(t, i+1, k.sampleRate)
			require.Equal(t, tone.DefaultNumHarmGains, k.numHarmGains)
		}
	})
}

// Test_BufferPool tests that BufferPool returns the correct buffer pool for the given context.
func Test_BufferPool(t *testing.T) {
	t.Run("nil context", func(t *testing.T) {
		require.Nil(t, BufferPool(nil))
	})

	t.Run("no value set", func(t *testing.T) {
		ctx := context.NewTestContext()
		require.Nil(t, BufferPool(ctx))
	})

	t.Run("non-pool value", func(t *testing.T) {
		for _, v := range []any{uint64(20), true, "20", 20.0, int8(20), byte(20), rune(20)} {
			ctx := context.NewContextWith(context.ContextOptions{
				Decorators: []context.Decorator{
					func(ctx context.Context) context.Context {
						return ctx.WithValue(bufferPoolKey, v)
					},
				},
			})
			require.Nil(t, BufferPool(ctx))
		}
	})

	t.Run("pool value", func(t *testing.T) {
		ctx := context.NewContext()
		require.NotNil(t, BufferPool(ctx))
	})
}
