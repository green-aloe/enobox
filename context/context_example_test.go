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
	minFreq := ctx.Value(ctxKeyMinFreq{})

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

func ExampleContext_SetSampleRate() {
	ctx := context.NewContext()
	sampleRate1 := ctx.SampleRate()

	ctx.SetSampleRate(48_000)
	sampleRate2 := ctx.SampleRate()

	fmt.Println(sampleRate1, sampleRate2)

	// Output:
	// 44100 48000
}

func ExampleContext_NyqistFrequency() {
	ctx := context.NewContext()
	nyquistFreq1 := ctx.NyqistFrequency()

	ctx.SetSampleRate(48_000)
	nyquistFreq2 := ctx.NyqistFrequency()

	fmt.Println(nyquistFreq1, nyquistFreq2)

	// Output:
	// 22050 24000
}
