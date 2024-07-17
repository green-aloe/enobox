package context

import (
	gocontext "context"
	"reflect"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

type ctxKey int

// Test_AddDecorator tests that AddDecorator adds valid decorators to the list of decorators.
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
		AddDecorator(func(ctx Context) Context { return ctx })
		require.Len(t, decorators, 1)
	})

	t.Run("concurrent access", func(t *testing.T) {
		decorators = nil

		var key ctxKey

		var wg sync.WaitGroup
		for i := range 1_000 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				AddDecorator(func(ctx Context) Context {
					ctx.SetValue(key, i)
					return ctx
				})
			}()
		}

		wg.Wait()
		require.Len(t, decorators, 1_000)

		ctx := &context{Context: gocontext.Background()}
		ints := make([]int, len(decorators))
		for i, decorator := range decorators {
			ctx := decorator(ctx)
			require.NotNil(t, ctx)
			v := ctx.Value(key)
			ints[i] = v.(int)
		}
		sort.Ints(ints)
		for i, n := range ints {
			require.Equal(t, i, n)
		}
	})
}

// Test_NewContext tests that NewContext returns a Context that has been initialized correctly with
// the default values.
func Test_NewContext(t *testing.T) {
	t.Run("types", func(t *testing.T) {
		ctx := NewContext()
		require.NotEmpty(t, ctx)

		ectx, ok := ctx.(*context)
		require.True(t, ok)
		require.IsType(t, gocontext.Background(), ectx.Context)

		rt := reflect.TypeOf(ctx)
		ri := reflect.TypeOf((*gocontext.Context)(nil)).Elem()
		require.True(t, rt.Implements(ri))
	})

	t.Run("initial values", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, NewTime(), ctx.Time())
		require.Equal(t, DefaultSampleRate, ctx.SampleRate())
	})

	t.Run("decorators", func(t *testing.T) {
		decorators = nil
		defer func() {
			decorators = nil
		}()

		const (
			key1 ctxKey = iota + 1
			key2
			key3
			key4
		)

		AddDecorator(func(ctx Context) Context {
			ctx.SetValue(key1, "value1")
			return ctx
		})
		AddDecorator(func(ctx Context) Context {
			ctx.SetValue(key4, 4)
			return ctx
		})
		AddDecorator(func(ctx Context) Context {
			return nil
		})
		AddDecorator(func(ctx Context) Context {
			ctx.SetTime(NewTimeAt(10, 11, 123))
			return ctx
		})
		AddDecorator(func(ctx Context) Context {
			ctx.SetValue(key2, '😬')
			return ctx
		})

		ctx := NewContext()
		require.Equal(t, 10, ctx.Time().Second())
		require.Equal(t, 11, ctx.Time().Sample())
		require.Equal(t, 123, ctx.Time().SampleRate())
		require.Equal(t, DefaultSampleRate, ctx.SampleRate())
		require.Equal(t, "value1", ctx.Value(key1))
		require.Equal(t, '😬', ctx.Value(key2))
		require.Nil(t, ctx.Value(key3))
		require.Equal(t, 4, ctx.Value(key4))
	})
}

// Test_NewContextWith tests that NewContextWith returns a Context that has been initialized
// correctly with the provided configuration.
func Test_NewContextWith(t *testing.T) {
	t.Run("types", func(t *testing.T) {
		ctx := NewContextWith(ContextOptions{})
		require.NotEmpty(t, ctx)

		ectx, ok := ctx.(*context)
		require.True(t, ok)
		require.IsType(t, gocontext.Background(), ectx.Context)

		rt := reflect.TypeOf(ctx)
		ri := reflect.TypeOf((*gocontext.Context)(nil)).Elem()
		require.True(t, rt.Implements(ri))
	})

	t.Run("default values", func(t *testing.T) {
		ctx := NewContextWith(ContextOptions{})
		require.IsType(t, gocontext.Background(), ctx.(*context).Context)
		require.Equal(t, NewTime(), ctx.Time())
		require.Equal(t, DefaultSampleRate, ctx.SampleRate())
	})

	t.Run("configured values", func(t *testing.T) {
		ctx := NewContextWith(ContextOptions{
			Context:    gocontext.TODO(),
			Time:       NewTimeWith(100).ShiftBy(123),
			SampleRate: 999,
		})
		require.IsType(t, gocontext.TODO(), ctx.(*context).Context)
		require.Equal(t, 1, ctx.Time().Second())
		require.Equal(t, 24, ctx.Time().Sample())
		require.Equal(t, 100, ctx.Time().SampleRate())
		require.Equal(t, 999, ctx.SampleRate())
	})

	t.Run("decorators", func(t *testing.T) {
		decorators = nil
		defer func() {
			decorators = nil
		}()

		const (
			key1 ctxKey = iota + 1
			key2
			key3
			key4
			key5
			key6
		)

		AddDecorator(func(ctx Context) Context {
			ctx.SetValue(key1, uint64(1))
			return ctx
		})
		AddDecorator(func(ctx Context) Context {
			ctx.SetValue(key4, "value4")
			return ctx
		})
		AddDecorator(func(ctx Context) Context {
			return nil
		})
		AddDecorator(func(ctx Context) Context {
			ctx.SetTime(NewTimeAt(20, 21, 21_000))
			return ctx
		})
		AddDecorator(func(ctx Context) Context {
			ctx.SetValue(key2, "value2")
			return ctx
		})

		ctx := NewContextWith(ContextOptions{
			Decorators: []Decorator{
				func(ctx Context) Context {
					return nil
				},
				func(ctx Context) Context {
					ctx.SetValue(key4, "value4444")
					return ctx
				},
				nil,
				func(ctx Context) Context {
					ctx.SetTime(NewTimeWith(789))
					return ctx
				},
				func(ctx Context) Context {
					ctx.SetValue(key5, "value5")
					return ctx
				},
				func(ctx Context) Context {
					ctx.SetValue(key6, struct{ int }{10})
					return ctx
				},
				func(ctx Context) Context {
					ctx.SetTime(NewTimeAt(90, 99, 999))
					return ctx
				},
			},
		})
		require.Equal(t, 90, ctx.Time().Second())
		require.Equal(t, 99, ctx.Time().Sample())
		require.Equal(t, 999, ctx.Time().SampleRate())
		require.Equal(t, DefaultSampleRate, ctx.SampleRate())
		require.Equal(t, uint64(1), ctx.Value(key1))
		require.Equal(t, "value2", ctx.Value(key2))
		require.Nil(t, ctx.Value(key3))
		require.Equal(t, "value4444", ctx.Value(key4))
		require.Equal(t, "value5", ctx.Value(key5))
		require.Equal(t, struct{ int }{10}, ctx.Value(key6))
	})
}

// Test_Context_SetValue tests that Context's SetValue method sets the correct value in the context.
func Test_Context_SetValue(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var ctx *context
		require.NotPanics(t, func() { ctx.SetValue("key", "value") })
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx context
		require.NotPanics(t, func() { ctx.SetValue("key", "value") })
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		ctx.SetValue("key", "value")
		require.Equal(t, "value", ctx.Value("key"))
	})

	t.Run("overwrite", func(t *testing.T) {
		ctx := NewContext()
		ctx.SetValue("key", "value1")
		ctx.SetValue("key", "value2")
		require.Equal(t, "value2", ctx.Value("key"))
	})

	t.Run("nil key", func(t *testing.T) {
		ctx := NewContext()
		require.Panics(t, func() { ctx.SetValue(nil, "value") })
		require.Nil(t, ctx.Value("key"))
	})

	t.Run("nil value", func(t *testing.T) {
		ctx := NewContext()
		ctx.SetValue("key", nil)
		require.Nil(t, ctx.Value("key"))
	})
}

// Test_Context_Time tests that Context's Time method returns the correct timestamp.
func Test_Context_Time(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var ctx *context
		require.Zero(t, ctx.Time())
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx context
		require.Zero(t, ctx.Time())
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, 0, ctx.Time().Second())
		require.Equal(t, 1, ctx.Time().Sample())
		require.Equal(t, DefaultSampleRate, ctx.Time().SampleRate())
	})

	t.Run("updated", func(t *testing.T) {
		ctx := NewContext()
		ctx.SetTime(ctx.Time().ShiftBy(10))
		require.Equal(t, 0, ctx.Time().Second())
		require.Equal(t, 11, ctx.Time().Sample())
		require.Equal(t, DefaultSampleRate, ctx.Time().SampleRate())
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
		var ctx *context
		ctx.SetTime(NewTimeAt(10, 20, 44_100))
		require.Zero(t, ctx.Time())
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx context
		ctx.SetTime(NewTimeAt(10, 20, 44_100))
		require.Equal(t, 10, ctx.Time().Second())
		require.Equal(t, 20, ctx.Time().Sample())
		require.Equal(t, 44_100, ctx.Time().SampleRate())
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		ctx.SetTime(NewTimeAt(10, 20, 100))
		require.Equal(t, 10, ctx.Time().Second())
		require.Equal(t, 20, ctx.Time().Sample())
		require.Equal(t, 100, ctx.Time().SampleRate())
	})

	t.Run("overwrite", func(t *testing.T) {
		ctx := NewContext()
		ctx.SetTime(NewTimeAt(1, 2, 10))
		ctx.SetTime(NewTimeAt(10, 20, 100))
		require.Equal(t, 10, ctx.Time().Second())
		require.Equal(t, 20, ctx.Time().Sample())
		require.Equal(t, 100, ctx.Time().SampleRate())
	})
}

// Test_Context_SampleRate tests that Context's SampleRate method returns the correct sample rate.
func Test_Context_SampleRate(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	t.Run("nil pointer", func(t *testing.T) {
		var ctx *context
		require.Zero(t, ctx.SampleRate())
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx context
		require.Zero(t, ctx.SampleRate())
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, DefaultSampleRate, ctx.SampleRate())
	})

	t.Run("updated", func(t *testing.T) {
		ctx := NewContextWith(ContextOptions{
			SampleRate: DefaultSampleRate + 99,
		})
		require.Equal(t, 44199, ctx.SampleRate())
	})

	t.Run("immutable", func(t *testing.T) {
		ctx := NewContext()
		SetSampleRate(100)
		require.Equal(t, 44100, ctx.SampleRate())
		require.Equal(t, 100, SampleRate())
	})
}

// Test_Context_NyqistFrequency tests that Context's NyqistFrequency method returns the correct
// frequency.
func Test_Context_NyqistFrequency(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	t.Run("nil pointer", func(t *testing.T) {
		var ctx *context
		require.Zero(t, ctx.NyqistFrequency())
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx context
		require.Zero(t, ctx.NyqistFrequency())
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, float32(22_050), ctx.NyqistFrequency())
	})
}

// Test_Context_Value tests that Context's Value method returns the correct value from a context and
// correctly handles missing values and bad configurations.
func Test_Context_Value(t *testing.T) {
	var testKey ctxKey

	t.Run("nil pointer", func(t *testing.T) {
		var ctx *context
		require.Panics(t, func() { ctx.Value(testKey) })
	})

	t.Run("uninitialized", func(t *testing.T) {
		var ctx context
		require.Panics(t, func() { ctx.Value(testKey) })
	})

	t.Run("missing key", func(t *testing.T) {
		ctx := NewContext()
		require.Nil(t, ctx.Value(testKey))
	})

	t.Run("initialized", func(t *testing.T) {
		ctx := NewContext()
		ctx.SetValue(testKey, "test")
		require.Equal(t, "test", ctx.Value(testKey))
	})
}
