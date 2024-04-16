package tone

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

func NewTone() Tone {
	return Tone{
		HarmonicGains: make([]float32, NumHarmGains),
	}
}
