package context_test

import (
	gocontext "context"
	"fmt"

	"github.com/green-aloe/enobox/context"
)

func ExampleAddDecorator() {
	type ctxKey struct{}
	var minFreqKey ctxKey

	context.AddDecorator(func(ctx context.Context) gocontext.Context {
		return gocontext.WithValue(ctx, minFreqKey, 22.22)
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

	fmt.Println(time)
	fmt.Println(sampleRate)

	ctx = context.NewContextWith(context.ContextOptions{
		Time:       context.NewTimeWith(35_000).ShiftBy(400),
		SampleRate: 35_000,
	})
	time = ctx.Time()
	sampleRate = ctx.SampleRate()

	fmt.Println(time)
	fmt.Println(sampleRate)

	// Output:
	// 0 seconds, sample 1/44100
	// 44100
	// 0 seconds, sample 401/35000
	// 35000
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
