package tone_test

import (
	"fmt"

	"github.com/green-aloe/enobox/tone"
)

func ExampleNewTone() {
	tone := tone.NewTone()
	fmt.Println(tone.Frequency, tone.Gain, len(tone.HarmonicGains))
	fmt.Println(tone.HarmonicGains)

	// Output:
	// 0 0 20
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExampleNewToneAt() {
	tone := tone.NewToneAt(523.2511)
	fmt.Println(tone.Frequency, tone.Gain, len(tone.HarmonicGains))
	fmt.Println(tone.HarmonicGains)

	// Output:
	// 523.2511 0 20
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExampleNewToneFrom() {
	tone := tone.NewToneFrom(tone.C, 5)
	fmt.Println(tone.Frequency, tone.Gain, len(tone.HarmonicGains))
	fmt.Println(tone.HarmonicGains)

	// Output:
	// 523.2511 0 20
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExampleNewSquareTone() {
	tone := tone.NewSquareTone(440)
	fmt.Println(tone.Frequency, tone.Gain, len(tone.HarmonicGains))
	fmt.Println(tone.HarmonicGains)

	// Output:
	// 440 0 20
	// [0 0.333333 0 0.2 0 0.142857 0 0.111111 0 0.090909 0 0.076923 0 0.0666666 0 0.0588235 0 0.0526315 0 0.047619]
}

func ExampleNewTriangleTone() {
	tone := tone.NewTriangleTone(440)
	fmt.Println(tone.Frequency, tone.Gain, len(tone.HarmonicGains))
	fmt.Println(tone.HarmonicGains)

	// Output:
	// 440 0 20
	// [0 0.111111 0 0.04 0 0.0204081 0 0.0123456 0 0.00826446 0 0.00591716 0 0.00444444 0 0.0034602 0 0.00277008 0 0.00226757]
}

func ExampleNewSawtoothTone() {
	tone := tone.NewSawtoothTone(440)
	fmt.Println(tone.Frequency, tone.Gain, len(tone.HarmonicGains))
	fmt.Println(tone.HarmonicGains)

	// Output:
	// 440 0 20
	// [0.5 0.333333 0.25 0.2 0.166666 0.142857 0.125 0.111111 0.1 0.090909 0.0833333 0.076923 0.0714285 0.0666666 0.0625 0.0588235 0.0555555 0.0526315 0.05 0.047619]
}

func ExampleTone_HarmonicFreq() {
	tone := tone.NewTone()
	fmt.Println(tone.HarmonicFreq(2))

	tone.Frequency = 440
	fmt.Println(tone.HarmonicFreq(1))
	fmt.Println(tone.HarmonicFreq(2))
	fmt.Println(tone.HarmonicFreq(3))

	// Output:
	// 0
	// 440
	// 880
	// 1320
}

func ExampleTone_Clone() {
	tone := tone.NewToneAt(523.2511)
	clone := tone.Clone()

	tone.Frequency = 440
	clone.Gain = 1

	fmt.Println(tone.Frequency, tone.Gain, len(tone.HarmonicGains))
	fmt.Println(clone.Frequency, clone.Gain, len(clone.HarmonicGains))

	// Output:
	// 440 0 20
	// 523.2511 1 20
}

func ExampleTone_Empty() {
	tone := tone.NewTone()
	fmt.Println(tone.Empty())

	tone.Frequency = 440
	fmt.Println(tone.Empty())

	tone.Frequency = 0
	tone.Gain = 1
	fmt.Println(tone.Empty())

	// Output:
	// true
	// false
	// false
}

func ExampleTrunc() {}
