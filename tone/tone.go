package tone

import (
	"fmt"
	"math"
	"strconv"

	"github.com/govalues/decimal"
)

const (
	// Number of harmonic gains above the fundamental frequency that are tracked for each tone by default
	NumHarmGains = 20
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

// Trunc truncates a  decimal to have no more than n digits.
func Trunc(f float32, n int) float32 {
	if f == 0 || n <= 0 {
		return 0
	}

	s := strconv.FormatFloat(float64(f), 'f', -1, 32)
	d := decimal.MustParse(s)

	numDigits := d.Prec()
	numRight := d.Scale()
	numLeft := numDigits - numRight

	// If the number has more digits in it than we want, we need to truncate it.
	if numDigits > n {

		// If the left side has all the digits that we want, then we can treat the number as an
		// integer and do some easy math to truncate it. Otherwise, we'll continue to parse it as a
		// float and truncate the decimal places.
		if numLeft >= n {
			m := int32(f)
			pow := int32(math.Pow10(numLeft - n))
			m /= pow
			m *= pow
			f = float32(m)
		} else {
			// Calculate how many decimal places we need to keep, and truncate the number.
			decPlaces := n - numLeft
			d = d.Trunc(decPlaces)

			f64, ok := d.Float64()
			if !ok {
				panic(fmt.Sprintf("Unable to convert '%v' to float64", d))
			}
			f = float32(f64)
		}
	}

	return f
}
