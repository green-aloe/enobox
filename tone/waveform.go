package tone

import "github.com/green-aloe/enobox/context"

var (
	// First 100 harmonic gains for a square waveform, truncated to 6 digits, calculated with this
	// formula:
	//   gain = 1 / (order + 2) if order is odd, 0 if order is even
	squareHarmGains = []float32{
		0, Trunc(1.0/3, MaxSigFigs), 0, Trunc(1.0/5, MaxSigFigs), 0, Trunc(1.0/7, MaxSigFigs), 0, Trunc(1.0/9, MaxSigFigs), 0, Trunc(1.0/11, MaxSigFigs),
		0, Trunc(1.0/13, MaxSigFigs), 0, Trunc(1.0/15, MaxSigFigs), 0, Trunc(1.0/17, MaxSigFigs), 0, Trunc(1.0/19, MaxSigFigs), 0, Trunc(1.0/21, MaxSigFigs),
		0, Trunc(1.0/23, MaxSigFigs), 0, Trunc(1.0/25, MaxSigFigs), 0, Trunc(1.0/27, MaxSigFigs), 0, Trunc(1.0/29, MaxSigFigs), 0, Trunc(1.0/31, MaxSigFigs),
		0, Trunc(1.0/33, MaxSigFigs), 0, Trunc(1.0/35, MaxSigFigs), 0, Trunc(1.0/37, MaxSigFigs), 0, Trunc(1.0/39, MaxSigFigs), 0, Trunc(1.0/41, MaxSigFigs),
		0, Trunc(1.0/43, MaxSigFigs), 0, Trunc(1.0/45, MaxSigFigs), 0, Trunc(1.0/47, MaxSigFigs), 0, Trunc(1.0/49, MaxSigFigs), 0, Trunc(1.0/51, MaxSigFigs),
		0, Trunc(1.0/53, MaxSigFigs), 0, Trunc(1.0/55, MaxSigFigs), 0, Trunc(1.0/57, MaxSigFigs), 0, Trunc(1.0/59, MaxSigFigs), 0, Trunc(1.0/61, MaxSigFigs),
		0, Trunc(1.0/63, MaxSigFigs), 0, Trunc(1.0/65, MaxSigFigs), 0, Trunc(1.0/67, MaxSigFigs), 0, Trunc(1.0/69, MaxSigFigs), 0, Trunc(1.0/71, MaxSigFigs),
		0, Trunc(1.0/73, MaxSigFigs), 0, Trunc(1.0/75, MaxSigFigs), 0, Trunc(1.0/77, MaxSigFigs), 0, Trunc(1.0/79, MaxSigFigs), 0, Trunc(1.0/81, MaxSigFigs),
		0, Trunc(1.0/83, MaxSigFigs), 0, Trunc(1.0/85, MaxSigFigs), 0, Trunc(1.0/87, MaxSigFigs), 0, Trunc(1.0/89, MaxSigFigs), 0, Trunc(1.0/91, MaxSigFigs),
		0, Trunc(1.0/93, MaxSigFigs), 0, Trunc(1.0/95, MaxSigFigs), 0, Trunc(1.0/97, MaxSigFigs), 0, Trunc(1.0/99, MaxSigFigs), 0, Trunc(1.0/101, MaxSigFigs),
	}

	// First 100 harmonic gains for a triangle waveform, truncated to 6 digits, calculated with this
	// formula:
	//   gain = 1 / (order^2 + 1) if order is odd, 0 if order is even
	triangleHarmGains = []float32{
		0, Trunc(1.0/9, MaxSigFigs), 0, Trunc(1.0/25, MaxSigFigs), 0, Trunc(1.0/49, MaxSigFigs), 0, Trunc(1.0/81, MaxSigFigs), 0, Trunc(1.0/121, MaxSigFigs),
		0, Trunc(1.0/169, MaxSigFigs), 0, Trunc(1.0/225, MaxSigFigs), 0, Trunc(1.0/289, MaxSigFigs), 0, Trunc(1.0/361, MaxSigFigs), 0, Trunc(1.0/441, MaxSigFigs),
		0, Trunc(1.0/529, MaxSigFigs), 0, Trunc(1.0/625, MaxSigFigs), 0, Trunc(1.0/729, MaxSigFigs), 0, Trunc(1.0/841, MaxSigFigs), 0, Trunc(1.0/961, MaxSigFigs),
		0, Trunc(1.0/1089, MaxSigFigs), 0, Trunc(1.0/1225, MaxSigFigs), 0, Trunc(1.0/1369, MaxSigFigs), 0, Trunc(1.0/1521, MaxSigFigs), 0, Trunc(1.0/1681, MaxSigFigs),
		0, Trunc(1.0/1849, MaxSigFigs), 0, Trunc(1.0/2025, MaxSigFigs), 0, Trunc(1.0/2209, MaxSigFigs), 0, Trunc(1.0/2401, MaxSigFigs), 0, Trunc(1.0/2601, MaxSigFigs),
		0, Trunc(1.0/2809, MaxSigFigs), 0, Trunc(1.0/3025, MaxSigFigs), 0, Trunc(1.0/3249, MaxSigFigs), 0, Trunc(1.0/3481, MaxSigFigs), 0, Trunc(1.0/3721, MaxSigFigs),
		0, Trunc(1.0/3969, MaxSigFigs), 0, Trunc(1.0/4225, MaxSigFigs), 0, Trunc(1.0/4489, MaxSigFigs), 0, Trunc(1.0/4761, MaxSigFigs), 0, Trunc(1.0/5041, MaxSigFigs),
		0, Trunc(1.0/5329, MaxSigFigs), 0, Trunc(1.0/5625, MaxSigFigs), 0, Trunc(1.0/5929, MaxSigFigs), 0, Trunc(1.0/6241, MaxSigFigs), 0, Trunc(1.0/6561, MaxSigFigs),
		0, Trunc(1.0/6889, MaxSigFigs), 0, Trunc(1.0/7225, MaxSigFigs), 0, Trunc(1.0/7569, MaxSigFigs), 0, Trunc(1.0/7921, MaxSigFigs), 0, Trunc(1.0/8281, MaxSigFigs),
		0, Trunc(1.0/8649, MaxSigFigs), 0, Trunc(1.0/9025, MaxSigFigs), 0, Trunc(1.0/9409, MaxSigFigs), 0, Trunc(1.0/9801, MaxSigFigs), 0, Trunc(1.0/10201, MaxSigFigs),
	}

	// First 100 harmonic gains for a sawtooth waveform, truncated to 6 digits, calculated with this
	// formula:
	//   gain = 1 / order
	sawtoothHarmGains = []float32{
		Trunc(1.0/2, MaxSigFigs), Trunc(1.0/3, MaxSigFigs), Trunc(1.0/4, MaxSigFigs), Trunc(1.0/5, MaxSigFigs), Trunc(1.0/6, MaxSigFigs),
		Trunc(1.0/7, MaxSigFigs), Trunc(1.0/8, MaxSigFigs), Trunc(1.0/9, MaxSigFigs), Trunc(1.0/10, MaxSigFigs), Trunc(1.0/11, MaxSigFigs),
		Trunc(1.0/12, MaxSigFigs), Trunc(1.0/13, MaxSigFigs), Trunc(1.0/14, MaxSigFigs), Trunc(1.0/15, MaxSigFigs), Trunc(1.0/16, MaxSigFigs),
		Trunc(1.0/17, MaxSigFigs), Trunc(1.0/18, MaxSigFigs), Trunc(1.0/19, MaxSigFigs), Trunc(1.0/20, MaxSigFigs), Trunc(1.0/21, MaxSigFigs),
		Trunc(1.0/22, MaxSigFigs), Trunc(1.0/23, MaxSigFigs), Trunc(1.0/24, MaxSigFigs), Trunc(1.0/25, MaxSigFigs), Trunc(1.0/26, MaxSigFigs),
		Trunc(1.0/27, MaxSigFigs), Trunc(1.0/28, MaxSigFigs), Trunc(1.0/29, MaxSigFigs), Trunc(1.0/30, MaxSigFigs), Trunc(1.0/31, MaxSigFigs),
		Trunc(1.0/32, MaxSigFigs), Trunc(1.0/33, MaxSigFigs), Trunc(1.0/34, MaxSigFigs), Trunc(1.0/35, MaxSigFigs), Trunc(1.0/36, MaxSigFigs),
		Trunc(1.0/37, MaxSigFigs), Trunc(1.0/38, MaxSigFigs), Trunc(1.0/39, MaxSigFigs), Trunc(1.0/40, MaxSigFigs), Trunc(1.0/41, MaxSigFigs),
		Trunc(1.0/42, MaxSigFigs), Trunc(1.0/43, MaxSigFigs), Trunc(1.0/44, MaxSigFigs), Trunc(1.0/45, MaxSigFigs), Trunc(1.0/46, MaxSigFigs),
		Trunc(1.0/47, MaxSigFigs), Trunc(1.0/48, MaxSigFigs), Trunc(1.0/49, MaxSigFigs), Trunc(1.0/50, MaxSigFigs), Trunc(1.0/51, MaxSigFigs),
		Trunc(1.0/52, MaxSigFigs), Trunc(1.0/53, MaxSigFigs), Trunc(1.0/54, MaxSigFigs), Trunc(1.0/55, MaxSigFigs), Trunc(1.0/56, MaxSigFigs),
		Trunc(1.0/57, MaxSigFigs), Trunc(1.0/58, MaxSigFigs), Trunc(1.0/59, MaxSigFigs), Trunc(1.0/60, MaxSigFigs), Trunc(1.0/61, MaxSigFigs),
		Trunc(1.0/62, MaxSigFigs), Trunc(1.0/63, MaxSigFigs), Trunc(1.0/64, MaxSigFigs), Trunc(1.0/65, MaxSigFigs), Trunc(1.0/66, MaxSigFigs),
		Trunc(1.0/67, MaxSigFigs), Trunc(1.0/68, MaxSigFigs), Trunc(1.0/69, MaxSigFigs), Trunc(1.0/70, MaxSigFigs), Trunc(1.0/71, MaxSigFigs),
		Trunc(1.0/72, MaxSigFigs), Trunc(1.0/73, MaxSigFigs), Trunc(1.0/74, MaxSigFigs), Trunc(1.0/75, MaxSigFigs), Trunc(1.0/76, MaxSigFigs),
		Trunc(1.0/77, MaxSigFigs), Trunc(1.0/78, MaxSigFigs), Trunc(1.0/79, MaxSigFigs), Trunc(1.0/80, MaxSigFigs), Trunc(1.0/81, MaxSigFigs),
		Trunc(1.0/82, MaxSigFigs), Trunc(1.0/83, MaxSigFigs), Trunc(1.0/84, MaxSigFigs), Trunc(1.0/85, MaxSigFigs), Trunc(1.0/86, MaxSigFigs),
		Trunc(1.0/87, MaxSigFigs), Trunc(1.0/88, MaxSigFigs), Trunc(1.0/89, MaxSigFigs), Trunc(1.0/90, MaxSigFigs), Trunc(1.0/91, MaxSigFigs),
		Trunc(1.0/92, MaxSigFigs), Trunc(1.0/93, MaxSigFigs), Trunc(1.0/94, MaxSigFigs), Trunc(1.0/95, MaxSigFigs), Trunc(1.0/96, MaxSigFigs),
		Trunc(1.0/97, MaxSigFigs), Trunc(1.0/98, MaxSigFigs), Trunc(1.0/99, MaxSigFigs), Trunc(1.0/100, MaxSigFigs), Trunc(1.0/101, MaxSigFigs),
	}
)

// NewSquareTone creates a new tone that has a square waveform.
//
// In a square waveform, the even-ordered harmonics have no gain, while the odd-ordered harmonics
// have gains that are the inverse of their orders. This waveform has a duty cycle of 50%, meaning
// that is has equal high and low periods.
//
// For example, a square tone that has a fundamental frequency (order=1) of 100Hz and a gain of 1
// has these first four harmonics:
//   - Harmonic 1: order = 2, frequency = 200Hz, gain = 0
//   - Harmonic 2: order = 3, frequency = 300Hz, gain = 1/3
//   - Harmonic 3: order = 4, frequency = 400Hz, gain = 0
//   - Harmonic 4: order = 5, frequency = 500Hz, gain = 1/5
func NewSquareTone(ctx context.Context, frequency float32) Tone {
	tone := NewToneAt(ctx, frequency)
	copy(tone.HarmonicGains, squareHarmGains)

	return tone
}

// NewTriangleTone creates a new tone that has a triangle waveform.
//
// In a triangle waveform, the even-ordered harmonics have no gain, while the odd-ordered harmonics
// have gains that are the inverse square of their orders (which also makes this waveform an
// integral of a square waveform). This leads to a much steeper roll-off than other waveforms.
//
// For example, a triangle tone that has a fundamental frequency (order=1) of 100Hz and a gain of 1
// has these first four harmonics:
//   - Harmonic 1: order = 2, frequency = 200Hz, gain = 0
//   - Harmonic 2: order = 3, frequency = 300Hz, gain = 1/9
//   - Harmonic 3: order = 4, frequency = 400Hz, gain = 0
//   - Harmonic 4: order = 5, frequency = 500Hz, gain = 1/25
func NewTriangleTone(ctx context.Context, frequency float32) Tone {
	tone := NewToneAt(ctx, frequency)
	copy(tone.HarmonicGains, triangleHarmGains)

	return tone
}

// NewSawtoothTone creates a new tone that has a sawtooth waveform.
//
// In a sawtooth waveform, each harmonic has a gain that is the inverse of its order. Because it is
// composed of every harmonic of the fundamental frequency, a sawtooth waveform is brighter and
// richer than other waveforms.
//
// For example, a sawtooth tone that has a fundamental frequency (order=1) of 100Hz and a gain of 1
// has these first four harmonics:
//   - Harmonic 1: order = 2, frequency = 200Hz, gain = 1/2
//   - Harmonic 2: order = 3, frequency = 300Hz, gain = 1/3
//   - Harmonic 3: order = 4, frequency = 400Hz, gain = 1/4
//   - Harmonic 4: order = 5, frequency = 500Hz, gain = 1/5
func NewSawtoothTone(ctx context.Context, frequency float32) Tone {
	tone := NewToneAt(ctx, frequency)
	copy(tone.HarmonicGains, sawtoothHarmGains)

	return tone
}
