package tone

import (
	"slices"

	"github.com/govalues/decimal"
	"github.com/green-aloe/enobox/context"
	"github.com/green-aloe/enobox/note"
)

const (
	// Maximum number of significant figures to use when truncating decimals
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
	// first element is the gain of the first harmonic (which has a frequency twice as high as the
	// fundamental frequency), the second element is the gain of the second harmonic (three times as
	// high), and so on.
	//
	// As an example, a value of 2 will double the amplitude of that harmonic relative to the
	// fundamental frequency, while a value of 0.5 will halve it.
	HarmonicGains []float32
}

// NewTone initializes a tone with default/zero values.
func NewTone(ctx context.Context) Tone {
	return NewToneAt(ctx, 0)
}

// NewToneAt initializes a tone with the specified fundamental frequency.
func NewToneAt(ctx context.Context, frequency float32) Tone {
	return Tone{
		Frequency:     frequency,
		HarmonicGains: make([]float32, NumHarmGains(ctx)),
	}
}

// NewToneFrom initializes a tone from the specified note and octave.
func NewToneFrom(ctx context.Context, note note.Note, octave int) Tone {
	return NewToneAt(ctx, note.Frequency(octave))
}

// NewToneWith initializes a tone with the specified fundamental frequency, gain, and harmonic
// gains. The harmonic gains are set directly in the tone, as opposed to allocating a new slice and
// copying over the values. If the number of harmonic gains provided does not match the number for
// the context, this grows or shrinks the slice to match the expected length.
func NewToneWith(ctx context.Context, frequency float32, gain float32, harmonicGains []float32) Tone {
	if want, have := NumHarmGains(ctx), len(harmonicGains); want != have {
		if want > cap(harmonicGains) {
			harmonicGains = slices.Grow(harmonicGains, want-have)
		}
		harmonicGains = harmonicGains[:want]
	}

	return Tone{
		Frequency:     frequency,
		Gain:          gain,
		HarmonicGains: harmonicGains,
	}
}

// HarmonicFreq calculates the frequency of one of the tone's harmonic. The fundamental frequency
// has an order of 1. The frequency is truncated to have no more than MaxSigFigs digits.
//
// As an example, if the tone has a fundamental frequency of 440Hz, then the first harmonic (order=2) is
// 880Hz and the second harmonic (order=3) is 1320Hz,
func (tone *Tone) HarmonicFreq(order int) float32 {
	if tone == nil || order <= 0 || tone.Frequency == 0 {
		return 0
	}

	freq := tone.Frequency
	if freq < 0 {
		freq = -freq
		order -= 2
	}

	fundFreq, _ := decimal.NewFromFloat64(float64(freq))
	multiplier, _ := decimal.New(int64(order), 0)
	harmFreq, _ := fundFreq.Mul(multiplier)
	f64, _ := harmFreq.Float64()

	return Trunc(float32(f64), MaxSigFigs)
}

// Clone returns a complete copy of the tone that has all of the same values as the original but
// does not share any memory with it.
func (tone *Tone) Clone() Tone {
	if tone == nil {
		return Tone{}
	}

	// Copy over all the basic fields.
	c := *tone

	// Copy over the extended fields.
	c.HarmonicGains = make([]float32, len(tone.HarmonicGains))
	copy(c.HarmonicGains, tone.HarmonicGains)

	return c
}

// Empty checks whether the tone does not have any values set.
func (tone *Tone) Empty() bool {
	if tone == nil {
		return true
	}

	if tone.Frequency != 0 {
		return false
	}

	if tone.Gain != 0 {
		return false
	}

	for _, harmGain := range tone.HarmonicGains {
		if harmGain != 0 {
			return false
		}
	}

	return true
}

// Reset resets the tone to its zero values. The harmonic gains are set to zero, but the slice header
// does not change.
func (tone *Tone) Reset() {
	if tone == nil {
		return
	}

	tone.Frequency = 0
	tone.Gain = 0
	for i := range tone.HarmonicGains {
		tone.HarmonicGains[i] = 0
	}
}
