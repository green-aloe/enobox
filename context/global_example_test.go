package context_test

import (
	"fmt"

	"github.com/green-aloe/enobox/context"
)

func ExampleSampleRate() {
	sampleRate := context.SampleRate()

	fmt.Println(sampleRate)

	// Output:
	// 44100
}

func ExampleSetSampleRate() {
	defer context.SetSampleRate(context.DefaultSampleRate)

	sampleRate1 := context.SampleRate()
	context.SetSampleRate(48_000)
	sampleRate2 := context.SampleRate()

	fmt.Println(sampleRate1, sampleRate2)

	// Output:
	// 44100 48000
}
