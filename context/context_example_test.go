package context_test

import (
	gocontext "context"
	"fmt"

	"github.com/green-aloe/enobox/context"
)

func ExampleAddDecorator() {
	type ctxKey struct{}
	var minFreqKey ctxKey

	context.AddDecorator(func(ctx context.Context) context.Context {
		ctx.SetValue(minFreqKey, 22.22)
		return ctx
	})

	ctx := context.NewContext()
	minFreq := ctx.Value(minFreqKey)

	fmt.Println(minFreq)

	// Output:
	// 22.22
}

func ExampleNewContext() {
	ctx := context.NewContext()
	time := ctx.Time()
	sampleRate := ctx.SampleRate()

	fmt.Println(time)
	fmt.Println(sampleRate)

	// Output:
	// 0 seconds, sample 1/44100
	// 44100
}

func ExampleNewContextWith() {
	ctx := context.NewContextWith(context.ContextOptions{})
	time := ctx.Time()
	sampleRate := ctx.SampleRate()
	value := ctx.Value("key")

	fmt.Println(time)
	fmt.Println(sampleRate)
	fmt.Println(value)

	ctx = context.NewContextWith(context.ContextOptions{
		Context:    gocontext.WithoutCancel(gocontext.Background()),
		Time:       context.NewTimeWith(35_000).ShiftBy(400),
		SampleRate: 35_000,
		Decorators: []context.Decorator{
			func(ctx context.Context) context.Context {
				ctx.SetValue("key", "value")
				return ctx
			},
		},
	})
	time = ctx.Time()
	sampleRate = ctx.SampleRate()
	value = ctx.Value("key")

	fmt.Println(time)
	fmt.Println(sampleRate)
	fmt.Println(value)

	// Output:
	// 0 seconds, sample 1/44100
	// 44100
	// <nil>
	// 0 seconds, sample 401/35000
	// 35000
	// value
}

func ExampleContext_SetValue() {
	ctx := context.NewContext()
	value := ctx.Value("key")

	fmt.Println(value)

	ctx.SetValue("key", "value")
	value = ctx.Value("key")

	fmt.Println(value)

	// Output:
	// <nil>
	// value
}

func ExampleContext_Time() {
	ctx := context.NewContext()
	time := ctx.Time()

	fmt.Println(time)

	// Output:
	// 0 seconds, sample 1/44100
}

func ExampleContext_SetTime() {
	ctx := context.NewContext()
	time1 := ctx.Time()

	time2 := time1.Increment()
	ctx.SetTime(time2)
	time3 := ctx.Time()

	fmt.Println(time1)
	fmt.Println(time2)
	fmt.Println(time3)

	// Output:
	// 0 seconds, sample 1/44100
	// 0 seconds, sample 2/44100
	// 0 seconds, sample 2/44100
}

func ExampleContext_SampleRate() {
	ctx := context.NewContext()
	sampleRate := ctx.SampleRate()

	fmt.Println(sampleRate)

	// Output:
	// 44100
}

func ExampleContext_NyqistFrequency() {
	ctx := context.NewContext()
	nyquistFreq := ctx.NyqistFrequency()

	fmt.Println(nyquistFreq)

	// Output:
	// 22050
}
