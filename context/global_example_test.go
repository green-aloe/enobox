package context_test

import (
	"fmt"

	"github.com/green-aloe/enobox/context"
)

func ExampleSampleRate() {
	fmt.Println(context.SampleRate())

	// Output:
	// 44100
}

func ExampleSetSampleRate() {
	defer context.SetSampleRate(context.DefaultSampleRate)

	fmt.Println(context.SampleRate())
	context.SetSampleRate(48_000)
	fmt.Println(context.SampleRate())

	// Output:
	// 44100
	// 48000
}
