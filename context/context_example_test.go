package context_test

import (
	gocontext "context"
	"fmt"

	"github.com/green-aloe/enobox/context"
)

func ExampleAddDecorator() {
	type ctxKeyMinFreq struct{}
	context.AddDecorator(func(ctx context.Context) gocontext.Context {
		return gocontext.WithValue(ctx, ctxKeyMinFreq{}, 22.22)
	})

	ctx := context.NewContext()
	fmt.Println(ctx.Value(ctxKeyMinFreq{}))

	// Output:
	// 22.22
}

func ExampleNewContext() {
	ctx := context.NewContext()

	fmt.Println(ctx.Time())
	fmt.Println(ctx.SampleRate())

	// Output:
	// 0 seconds, sample 1/44100
	// 44100
}

func ExampleContext_Time() {
	ctx := context.NewContext()

	fmt.Println(ctx.Time())

	// Output:
	// 0 seconds, sample 1/44100
}

func ExampleContext_SetTime() {
	ctx := context.NewContext()

	fmt.Println(ctx.Time())
	ctx.SetTime(ctx.Time().Increment())
	fmt.Println(ctx.Time())

	// Output:
	// 0 seconds, sample 1/44100
	// 0 seconds, sample 2/44100
}

func ExampleContext_SampleRate() {
	ctx := context.NewContext()

	fmt.Println(ctx.SampleRate())
	// Output:
	// 44100
}

func ExampleContext_SetSampleRate() {
	ctx := context.NewContext()

	fmt.Println(ctx.SampleRate())
	ctx.SetSampleRate(48_000)
	fmt.Println(ctx.SampleRate())

	// Output:
	// 44100
	// 48000
}

func ExampleContext_NyqistFrequency() {
	ctx := context.NewContext()

	fmt.Println(ctx.NyqistFrequency())
	ctx.SetSampleRate(48_000)
	fmt.Println(ctx.NyqistFrequency())

	// Output:
	// 22050
	// 24000
}
