package tone_test

import (
	"fmt"

	"github.com/green-aloe/enobox/context"
	"github.com/green-aloe/enobox/note"
	"github.com/green-aloe/enobox/tone"
)

func ExampleNewTone() {
	ctx := context.NewContext()

	tone := tone.NewTone(ctx)

	fmt.Println(tone.Frequency, tone.Gain, tone.HarmonicGains)

	// Output:
	// 0 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExampleNewToneAt() {
	ctx := context.NewContext()

	tone := tone.NewToneAt(ctx, 523.2511)

	fmt.Println(tone.Frequency, tone.Gain, tone.HarmonicGains)

	// Output:
	// 523.2511 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExampleNewToneFrom() {
	ctx := context.NewContext()

	tone := tone.NewToneFrom(ctx, note.C, 5)

	fmt.Println(tone.Frequency, tone.Gain, tone.HarmonicGains)

	// Output:
	// 523.2511 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExampleNewToneWith() {
	ctx := context.NewContext()

	for _, harmGains := range [][]float32{
		{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0},
		{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0},
		{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9, 3.0},
	} {
		tone := tone.NewToneWith(ctx, 440, 1, harmGains)

		fmt.Println(tone.Frequency, tone.Gain, tone.HarmonicGains)
	}

	// Output:
	// 440 1 [0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9 1 0 0 0 0 0 0 0 0 0 0]
	// 440 1 [0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9 1 1.1 1.2 1.3 1.4 1.5 1.6 1.7 1.8 1.9 2]
	// 440 1 [0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9 1 1.1 1.2 1.3 1.4 1.5 1.6 1.7 1.8 1.9 2]

}

func ExampleNewSquareTone() {
	ctx := context.NewContext()

	sqrTone := tone.NewSquareTone(ctx, 440)

	fmt.Println(sqrTone.Frequency, sqrTone.Gain, sqrTone.HarmonicGains)

	// Show the calculation for each harmonic gain.
	// gain = 1 / order if order is odd, 0 if order is even
	gains := make([]float32, len(sqrTone.HarmonicGains))
	for i := range gains {
		if i%2 > 0 {
			gains[i] = tone.Trunc(1/float32(i+2), tone.MaxSigFigs)
		}
	}

	fmt.Println(gains)

	// Output:
	// 440 0 [0 0.333333 0 0.2 0 0.142857 0 0.111111 0 0.090909 0 0.076923 0 0.0666666 0 0.0588235 0 0.0526315 0 0.047619]
	// [0 0.333333 0 0.2 0 0.142857 0 0.111111 0 0.090909 0 0.076923 0 0.0666666 0 0.0588235 0 0.0526315 0 0.047619]
}

func ExampleNewTriangleTone() {
	ctx := context.NewContext()

	triTone := tone.NewTriangleTone(ctx, 440)

	fmt.Println(triTone.Frequency, triTone.Gain, triTone.HarmonicGains)

	// Show the calculation for each harmonic gain.
	// gain = 1 / order^2 if order is odd, 0 if order is even
	gains := make([]float32, len(triTone.HarmonicGains))
	for i := range gains {
		if i%2 > 0 {
			gains[i] = tone.Trunc(1/float32((i+2)*(i+2)), tone.MaxSigFigs)
		}
	}

	fmt.Println(gains)

	// Output:
	// 440 0 [0 0.111111 0 0.04 0 0.0204081 0 0.0123456 0 0.00826446 0 0.00591716 0 0.00444444 0 0.0034602 0 0.00277008 0 0.00226757]
	// [0 0.111111 0 0.04 0 0.0204081 0 0.0123456 0 0.00826446 0 0.00591716 0 0.00444444 0 0.0034602 0 0.00277008 0 0.00226757]
}

func ExampleNewSawtoothTone() {
	ctx := context.NewContext()

	sawTone := tone.NewSawtoothTone(ctx, 440)

	fmt.Println(sawTone.Frequency, sawTone.Gain, sawTone.HarmonicGains)

	// Show the calculation for each harmonic gain.
	// gain = 1 / order
	gains := make([]float32, len(sawTone.HarmonicGains))
	for i := range gains {
		gains[i] = tone.Trunc(1/float32(i+2), tone.MaxSigFigs)
	}

	fmt.Println(gains)

	// Output:
	// 440 0 [0.5 0.333333 0.25 0.2 0.166666 0.142857 0.125 0.111111 0.1 0.090909 0.0833333 0.076923 0.0714285 0.0666666 0.0625 0.0588235 0.0555555 0.0526315 0.05 0.047619]
	// [0.5 0.333333 0.25 0.2 0.166666 0.142857 0.125 0.111111 0.1 0.090909 0.0833333 0.076923 0.0714285 0.0666666 0.0625 0.0588235 0.0555555 0.0526315 0.05 0.047619]
}

func ExampleTone_HarmonicFreq() {
	ctx := context.NewContext()

	tone := tone.NewTone(ctx)
	harmFreq1 := tone.HarmonicFreq(2)

	tone.Frequency = 440

	harmFreq2 := tone.HarmonicFreq(1)
	harmFreq3 := tone.HarmonicFreq(2)
	harmFreq4 := tone.HarmonicFreq(3)

	fmt.Println(harmFreq1, harmFreq2, harmFreq3, harmFreq4)

	// Output:
	// 0 440 880 1320
}

func ExampleTone_Clone() {
	ctx := context.NewContext()

	tone := tone.NewToneWith(ctx, 523.2511, 0.5, []float32{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0})
	clone := tone.Clone()

	tone.Frequency = 440
	tone.Gain = 1
	tone.HarmonicGains[0] = 100
	tone.HarmonicGains[1] = 200
	tone.HarmonicGains[2] = 300

	fmt.Println(tone.Frequency, tone.Gain, tone.HarmonicGains)
	fmt.Println(clone.Frequency, clone.Gain, clone.HarmonicGains)

	// Output:
	// 440 1 [100 200 300 0.4 0.5 0.6 0.7 0.8 0.9 1 1.1 1.2 1.3 1.4 1.5 1.6 1.7 1.8 1.9 2]
	// 523.2511 0.5 [0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9 1 1.1 1.2 1.3 1.4 1.5 1.6 1.7 1.8 1.9 2]
}

func ExampleTone_Empty() {
	ctx := context.NewContext()

	tone := tone.NewTone(ctx)
	isEmpty1 := tone.Empty()

	tone.Frequency = 440
	isEmpty2 := tone.Empty()

	tone.Frequency, tone.Gain = 0, 1
	isEmpty3 := tone.Empty()

	fmt.Println(isEmpty1, isEmpty2, isEmpty3)

	// Output:
	// true false false
}

func ExampleTone_Reset() {
	ctx := context.NewContext()

	tone := tone.NewSquareTone(ctx, 440)

	fmt.Println(tone.Frequency, tone.Gain, tone.HarmonicGains)

	tone.Reset()

	fmt.Println(tone.Frequency, tone.Gain, tone.HarmonicGains)

	// Output:
	// 440 0 [0 0.333333 0 0.2 0 0.142857 0 0.111111 0 0.090909 0 0.076923 0 0.0666666 0 0.0588235 0 0.0526315 0 0.047619]
	// 0 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExampleTrunc() {
	frequencies := []float32{0.123456789, 0.987654321, 99.99}
	for _, f1 := range frequencies {
		f2 := tone.Trunc(f1, 3)
		f3 := tone.Trunc(f1, tone.MaxSigFigs)

		fmt.Println(f1, f2, f3)
	}

	// Output:
	// 0.12345679 0.123 0.123456
	// 0.9876543 0.987 0.987654
	// 99.99 99.9 99.99
}
