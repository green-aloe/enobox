package tone

import (
	"github.com/govalues/decimal"
)

const (
	// Number of harmonic gains above the fundamental frequency that are tracked for each tone by default
	NumHarmGains = 20

	// MaxSigFigs is the maximum number of significant figures to use when rounding decimals
	MaxSigFigs = 6
)

// A Tone represents a single tone, which is a fundamental frequency and its harmonics at regular
// intervals above it. A new tone must be created with NewTone before it can be used.
type Tone struct {
	// Frequency is the tone's fundamental frequency.
	Frequency float32

	// Gain is a multiplier that increases or decreases the amplitude of the tone. It is a ratio of
	// the amplitude of the output signal to the amplitude of the base signal.
	//
	// As an example, a value of 2 will double the signal's strength, while a value of 0.5 will
	// halve it.
	Gain float32

	// HarmonicGains is a list of gains for each harmonic above the fundamental frequency. They are
	// ratios of the amplitude of the harmonic to the amplitude of the fundamental frequency. The
	// first element is the gain of the first harmonic, the second element is the gain of the second
	// harmonic, and so on.
	//
	// As an example, a value of 2 will double the amplitude of that harmonic relative to the
	// fundamental frequency, while a value of 0.5 will halve it.
	HarmonicGains []float32
}

// NewTone initializes a tone with default/zero values.
func NewTone() Tone {
	return Tone{
		HarmonicGains: make([]float32, NumHarmGains),
	}
}

// NewToneAt initializes a tone with the specified fundamental frequency.
func NewToneAt(frequency float32) Tone {
	tone := NewTone()
	tone.Frequency = frequency

	return tone
}

// HarmonicFreq calculates the frequency of the tone's nth harmonic. The frequency is truncated to
// have no more than MaxSigFigs significant figures.
//
// As an example, if the tone has a fundamental frequency of 440Hz, then the first harmonic (n=1) is
// 880Hz and the second harmonic (n=2) is 1320Hz,
func (tone Tone) HarmonicFreq(n int) float32 {
	if n <= 0 || tone.Frequency == 0 {
		return 0
	}

	freq := tone.Frequency
	if freq < 0 {
		freq = -freq
		n--
	} else {
		n++
	}

	fundFreq, _ := decimal.NewFromFloat64(float64(freq))
	multiplier, _ := decimal.New(int64(n), 0)
	harmFreq, _ := fundFreq.Mul(multiplier)
	f64, _ := harmFreq.Float64()

	return Trunc(float32(f64), MaxSigFigs)
}

// Clone returns a complete copy of the tone that has all of the same values as the original but
// does not share any memory with it.
func (tone Tone) Clone() Tone {
	// Copy over all the basic fields.
	c := tone

	// Copy over the extended fields.
	c.HarmonicGains = make([]float32, len(tone.HarmonicGains))
	copy(c.HarmonicGains, tone.HarmonicGains)

	return c
}
