package context

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AddDecorator(t *testing.T) {
	defer func() {
		decorators = nil
	}()

	t.Run("nil decorator", func(t *testing.T) {
		require.Len(t, decorators, 0)
		AddDecorator(nil)
		require.Len(t, decorators, 0)
	})

	t.Run("non-nil decorator", func(t *testing.T) {
		require.Len(t, decorators, 0)
		AddDecorator(func(ctx Context) context.Context { return ctx })
		require.Len(t, decorators, 1)
	})

	t.Run("concurrent access", func(t *testing.T) {
		decorators = nil

		type key struct{}

		var wg sync.WaitGroup
		for i := 0; i < 1_000; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				AddDecorator(func(Context) context.Context {
					return context.WithValue(context.Background(), key{}, i)
				})
			}(i)
		}

		wg.Wait()
		require.Len(t, decorators, 1_000)

		ints := make([]int, 1_000)
		for i, decorator := range decorators {
			ctx := decorator(Context{})
			require.NotNil(t, ctx)
			v := ctx.Value(key{})
			ints[i] = v.(int)
		}
		sort.Ints(ints)
		for i, v := range ints {
			require.Equal(t, i, v)
		}
	})
}

// Test_NewContext tests that NewContext returns a Context that has been initialized correctly.
func Test_NewContext(t *testing.T) {
	t.Run("types", func(t *testing.T) {
		ctx := NewContext()
		require.NotEmpty(t, ctx)

		require.IsType(t, Context{}, ctx)

		rt := reflect.TypeOf(&ctx)
		ri := reflect.TypeOf((*context.Context)(nil)).Elem()
		require.True(t, rt.Implements(ri))
	})

	t.Run("initial values", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, NewTime(), ctx.Time())
		require.Equal(t, DefaultSampleRate, ctx.SampleRate())
		require.IsType(t, context.Background(), ctx.Context)
	})

	t.Run("decorators", func(t *testing.T) {
		decorators = nil
		defer func() {
			decorators = nil
		}()

		type (
			key1 struct{}
			key2 struct{}
			key3 struct{}
			key4 struct{}
		)

		AddDecorator(func(ctx Context) context.Context {
			return context.WithValue(ctx, key1{}, "value1")
		})
		AddDecorator(func(ctx Context) context.Context {
			return context.WithValue(ctx, key4{}, "value4")
		})
		AddDecorator(func(ctx Context) context.Context {
			return nil
		})
		AddDecorator(func(ctx Context) context.Context {
			ctx.SetSampleRate(123)
			return ctx
		})
		AddDecorator(func(ctx Context) context.Context {
			return context.WithValue(ctx, key2{}, "value2")
		})

		ctx := NewContext()
		require.Equal(t, NewTime(), ctx.Time())
		require.Equal(t, 123, ctx.SampleRate())
		require.Equal(t, "value1", ctx.Value(key1{}))
		require.Equal(t, "value2", ctx.Value(key2{}))
		require.Nil(t, ctx.Value(key3{}))
		require.Equal(t, "value4", ctx.Value(key4{}))
	})
}

// Test_Context_Time tests that Context's Time method returns the correct timestamp.
func Test_Context_Time(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var ctx *Context
		require.Zero(t, ctx.Time())
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx Context
		require.Zero(t, ctx.Time())
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, 0, ctx.Time().Second())
		require.Equal(t, 1, ctx.Time().Sample())
	})

	t.Run("updated", func(t *testing.T) {
		ctx := NewContext()
		ctx.time = ctx.Time().ShiftBy(10)
		require.Equal(t, 0, ctx.Time().Second())
		require.Equal(t, 11, ctx.Time().Sample())
	})

	t.Run("immutable", func(t *testing.T) {
		ctx := NewContext()
		time := ctx.Time()
		time.second, time.sample = 10, 11
		require.Equal(t, 0, ctx.Time().Second())
		require.Equal(t, 1, ctx.Time().Sample())
		require.Equal(t, 10, time.Second())
		require.Equal(t, 11, time.Sample())
	})
}

// Test_Context_SetTime tests that Context's SetTime method sets the correct timestamp in the
// context.
func Test_Context_SetTime(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var ctx *Context
		ctx.SetTime(Time{10, 20})
		require.Zero(t, ctx.Time())
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx Context
		ctx.SetTime(Time{10, 20})
		require.Equal(t, 10, ctx.Time().Second())
		require.Equal(t, 20, ctx.Time().Sample())
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		ctx.SetTime(Time{10, 20})
		require.Equal(t, 10, ctx.Time().Second())
		require.Equal(t, 20, ctx.Time().Sample())
	})
}

// Test_Context_SampleRate tests that Context's SampleRate method returns the correct sample rate.
func Test_Context_SampleRate(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	t.Run("nil pointer", func(t *testing.T) {
		var ctx *Context
		require.Zero(t, ctx.SampleRate())
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx Context
		require.Zero(t, ctx.SampleRate())
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, DefaultSampleRate, ctx.SampleRate())
	})

	t.Run("updated", func(t *testing.T) {
		ctx := NewContext()
		ctx.sampleRate = ctx.SampleRate() + 99
		require.Equal(t, 44199, ctx.SampleRate())
	})

	t.Run("immutable", func(t *testing.T) {
		ctx := NewContext()
		SetSampleRate(100)
		require.Equal(t, 44100, ctx.SampleRate())
		require.Equal(t, 100, SampleRate())
	})
}

// Test_Context_SetSampleRate tests that Context's SetSampleRate method sets the correct sample rate
// in the context.
func Test_Context_SetSampleRate(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	t.Run("nil pointer", func(t *testing.T) {
		var ctx *Context
		ctx.SetSampleRate(10)
		require.Zero(t, ctx.SampleRate())
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx Context
		ctx.SetSampleRate(10)
		require.Equal(t, 10, ctx.SampleRate())
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		ctx.SetSampleRate(10)
		require.Equal(t, 10, ctx.SampleRate())
	})
}

// Test_Context_NyqistFrequency tests that Context's NyqistFrequency method returns the correct
// frequency.
func Test_Context_NyqistFrequency(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	t.Run("nil pointer", func(t *testing.T) {
		var ctx *Context
		require.Zero(t, ctx.NyqistFrequency())
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx Context
		require.Zero(t, ctx.NyqistFrequency())
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, float32(22_050), ctx.NyqistFrequency())
	})
}

// Test_Context_Value tests that Context's Value method returns the correct value from a context and
// handles missing values and bad configurations correctly.
func Test_Context_Value(t *testing.T) {
	type contextKeyTest struct{}

	t.Run("nil pointer", func(t *testing.T) {
		var ctx *Context
		require.Panics(t, func() { ctx.Value(contextKeyTest{}) })
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx Context
		require.Panics(t, func() { ctx.Value(contextKeyTest{}) })
	})

	t.Run("missing key", func(t *testing.T) {
		ctx := NewContext()
		require.Nil(t, ctx.Value(contextKeyTest{}))
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		ctx.Context = context.WithValue(ctx, contextKeyTest{}, "test")
		require.Equal(t, "test", ctx.Value(contextKeyTest{}))
	})
}
