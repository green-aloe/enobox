package tone

import (
	"testing"

	"github.com/green-aloe/enobox/context"
	"github.com/stretchr/testify/require"
)

// Test_Consts tests any package-level constant values.
func Test_Consts(t *testing.T) {
	require.Equal(t, 20, DefaultNumHarmGains)

	require.Equal(t, 6, MaxSigFigs)
}

// Test_init tests that a context decorator is added on package initialization with the default
// number of harmonic gains.
func Test_init(t *testing.T) {
	ctx := context.NewContext()
	require.Equal(t, 20, NumHarmGains(ctx))
}

// Test_NumHarmGains tests that NumHarmGains returns the correct number of harmonic gains in a tone
// for the given context.
func Test_NumHarmGains(t *testing.T) {
	t.Run("nil context", func(t *testing.T) {
		require.Zero(t, NumHarmGains(nil))
	})

	t.Run("no value set", func(t *testing.T) {
		ctx := context.NewTestContext()
		require.Zero(t, NumHarmGains(ctx))
	})

	t.Run("non-integer value", func(t *testing.T) {
		for _, v := range []any{uint64(20), true, "20", 20.0, int8(20), byte(20), rune(20)} {
			ctx := context.NewContextWith(context.ContextOptions{
				Decorators: []context.Decorator{
					func(ctx context.Context) context.Context {
						return ctx.WithValue(numHarmGainsKey, v)
					},
				},
			})
			require.Zero(t, NumHarmGains(ctx))
		}
	})

	t.Run("invalid values", func(t *testing.T) {
		for _, n := range []int{0, -1, -20} {
			ctx := context.NewContextWith(context.ContextOptions{
				Decorators: []context.Decorator{
					func(ctx context.Context) context.Context {
						return ctx.WithValue(numHarmGainsKey, n)
					},
				},
			})
			require.Zero(t, NumHarmGains(ctx))
		}
	})

	t.Run("valid values", func(t *testing.T) {
		for n := range 1_000 {
			ctx := context.NewContextWith(context.ContextOptions{
				Decorators: []context.Decorator{
					func(ctx context.Context) context.Context {
						return ctx.WithValue(numHarmGainsKey, n+1)
					},
				},
			})
			require.Equal(t, n+1, NumHarmGains(ctx))
		}
	})
}
