package buffer_test

import (
	"fmt"

	"github.com/green-aloe/enobox/buffer"
	"github.com/green-aloe/enobox/context"
	"github.com/green-aloe/enobox/tone"
)

func ExampleBufferPool() {
	defer tone.SetNumHarmGains(tone.DefaultNumHarmGains)

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

		pool := buffer.BufferPool(ctx)
		buffer := pool.Get()

		fmt.Println(len(buffer.Tones), len(buffer.Tones[0].HarmonicGains))
	}

	// Output:
	// 44100 20
	// 48000 10
	// 96000 100
}
