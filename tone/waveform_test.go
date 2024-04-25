package tone

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_NewSquareTone tests that NewSquareTone returns a tone that has a square waveform.
func Test_NewSquareTone(t *testing.T) {
	wantHarmGains := []float32{
		0, 0.333333, 0, 0.200000, 0, 0.142857, 0, 0.111111, 0, 0.0909090,
		0, 0.0769230, 0, 0.0666666, 0, 0.0588235, 0, 0.0526315, 0, 0.0476190,
		0, 0.0434782, 0, 0.0400000, 0, 0.0370370, 0, 0.0344827, 0, 0.0322580,
		0, 0.0303030, 0, 0.0285714, 0, 0.0270270, 0, 0.0256410, 0, 0.0243902,
		0, 0.0232558, 0, 0.0222222, 0, 0.0212765, 0, 0.0204081, 0, 0.0196078,
		0, 0.0188679, 0, 0.0181818, 0, 0.0175438, 0, 0.0169491, 0, 0.0163934,
		0, 0.0158730, 0, 0.0153846, 0, 0.0149253, 0, 0.0144927, 0, 0.0140845,
		0, 0.0136986, 0, 0.0133333, 0, 0.0129870, 0, 0.0126582, 0, 0.0123456,
		0, 0.0120481, 0, 0.0117647, 0, 0.0114942, 0, 0.0112359, 0, 0.0109890,
		0, 0.0107526, 0, 0.0105263, 0, 0.0103092, 0, 0.0101010, 0, 0.00990099,
	}

	t.Run("precalculations", func(t *testing.T) {
		require.Equal(t, len(wantHarmGains), len(squareHarmGains))
		for i, have := range squareHarmGains {
			want := wantHarmGains[i]
			require.Equal(t, want, have)
		}
	})

	type subtest struct {
		frequency float32
		name      string
	}

	subtests := []subtest{
		{0, "no frequency"},
		{1.1, "small float frequency"},
		{440, "medium integer frequency"},
		{112233.44, "high float frequency"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			// Test positive frequencies.
			tone := NewSquareTone(subtest.frequency)
			require.Equal(t, subtest.frequency, tone.Frequency)
			require.Len(t, tone.HarmonicGains, NumHarmGains)
			for i, harmGain := range tone.HarmonicGains {
				require.Equal(t, wantHarmGains[i], harmGain)
			}

			// Test negative frequencies.
			tone = NewSquareTone(-subtest.frequency)
			require.Equal(t, -subtest.frequency, tone.Frequency)
			require.Len(t, tone.HarmonicGains, NumHarmGains)
			for i, harmGain := range tone.HarmonicGains {
				require.Equal(t, wantHarmGains[i], harmGain)
			}
		})
	}
}

// Test_NewTriangleTone tests that NewTriangleTone returns a tone that has a triangle waveform.
func Test_NewTriangleTone(t *testing.T) {
	wantHarmGains := []float32{
		0, 0.111111, 0, 0.0400000, 0, 0.0204081, 0, 0.0123456, 0, 0.00826446,
		0, 0.00591716, 0, 0.00444444, 0, 0.00346020, 0, 0.00277008, 0, 0.00226757,
		0, 0.00189035, 0, 0.00160000, 0, 0.00137174, 0, 0.00118906, 0, 0.00104058,
		0, 0.000918273, 0, 0.000816326, 0, 0.000730460, 0, 0.000657462, 0, 0.000594884,
		0, 0.000540832, 0, 0.000493827, 0, 0.000452693, 0, 0.000416493, 0, 0.000384467,
		0, 0.000355998, 0, 0.000330578, 0, 0.000307787, 0, 0.000287273, 0, 0.000268744,
		0, 0.000251952, 0, 0.000236686, 0, 0.000222766, 0, 0.000210039, 0, 0.000198373,
		0, 0.000187652, 0, 0.000177777, 0, 0.000168662, 0, 0.000160230, 0, 0.000152415,
		0, 0.000145158, 0, 0.000138408, 0, 0.000132117, 0, 0.000126246, 0, 0.000120758,
		0, 0.000115620, 0, 0.000110803, 0, 0.000106281, 0, 0.000102030, 0, 0.0000980296,
	}

	t.Run("precalculations", func(t *testing.T) {
		require.Equal(t, len(wantHarmGains), len(triangleHarmGains))
		for i, have := range triangleHarmGains {
			want := wantHarmGains[i]
			require.Equal(t, want, have)
		}
	})

	type subtest struct {
		frequency float32
		name      string
	}

	subtests := []subtest{
		{0, "no frequency"},
		{1.1, "small float frequency"},
		{440, "medium integer frequency"},
		{112233.44, "high float frequency"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			// Test positive frequencies.
			tone := NewTriangleTone(subtest.frequency)
			require.Equal(t, subtest.frequency, tone.Frequency)
			require.Len(t, tone.HarmonicGains, NumHarmGains)
			for i, harmGain := range tone.HarmonicGains {
				require.Equal(t, wantHarmGains[i], harmGain)
			}

			// Test negative frequencies.
			tone = NewTriangleTone(-subtest.frequency)
			require.Equal(t, -subtest.frequency, tone.Frequency)
			require.Len(t, tone.HarmonicGains, NumHarmGains)
			for i, harmGain := range tone.HarmonicGains {
				require.Equal(t, wantHarmGains[i], harmGain)
			}
		})
	}
}

// Test_NewSawtoothTone tests that NewSawtoothTone returns a tone that has a sawtooth waveform.
func Test_NewSawtoothTone(t *testing.T) {
	wantHarmGains := []float32{
		0.500000, 0.333333, 0.250000, 0.200000, 0.166666, 0.142857, 0.125000, 0.111111, 0.100000, 0.0909090,
		0.0833333, 0.0769230, 0.0714285, 0.0666666, 0.0625000, 0.0588235, 0.0555555, 0.0526315, 0.0500000, 0.0476190,
		0.0454545, 0.0434782, 0.0416666, 0.0400000, 0.0384615, 0.0370370, 0.0357142, 0.0344827, 0.0333333, 0.0322580,
		0.0312500, 0.0303030, 0.0294117, 0.0285714, 0.0277777, 0.0270270, 0.0263157, 0.0256410, 0.0250000, 0.0243902,
		0.0238095, 0.0232558, 0.0227272, 0.0222222, 0.0217391, 0.0212765, 0.0208333, 0.0204081, 0.0200000, 0.0196078,
		0.0192307, 0.0188679, 0.0185185, 0.0181818, 0.0178571, 0.0175438, 0.0172413, 0.0169491, 0.0166666, 0.0163934,
		0.0161290, 0.0158730, 0.0156250, 0.0153846, 0.0151515, 0.0149253, 0.0147058, 0.0144927, 0.0142857, 0.0140845,
		0.0138888, 0.0136986, 0.0135135, 0.0133333, 0.0131578, 0.0129870, 0.0128205, 0.0126582, 0.0125000, 0.0123456,
		0.0121951, 0.0120481, 0.0119047, 0.0117647, 0.0116279, 0.0114942, 0.0113636, 0.0112359, 0.0111111, 0.0109890,
		0.0108695, 0.0107526, 0.0106382, 0.0105263, 0.0104166, 0.0103092, 0.0102040, 0.0101010, 0.0100000, 0.00990099,
	}

	t.Run("precalculations", func(t *testing.T) {
		require.Equal(t, len(wantHarmGains), len(sawtoothHarmGains))
		for i, have := range sawtoothHarmGains {
			want := wantHarmGains[i]
			require.Equal(t, want, have)
		}
	})

	type subtest struct {
		frequency float32
		name      string
	}

	subtests := []subtest{
		{0, "no frequency"},
		{1.1, "small float frequency"},
		{440, "medium integer frequency"},
		{112233.44, "high float frequency"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			// Test positive frequencies.
			tone := NewSawtoothTone(subtest.frequency)
			require.Equal(t, subtest.frequency, tone.Frequency)
			require.Len(t, tone.HarmonicGains, NumHarmGains)
			for i, harmGain := range tone.HarmonicGains {
				require.Equal(t, wantHarmGains[i], harmGain)
			}

			// Test negative frequencies.
			tone = NewSawtoothTone(-subtest.frequency)
			require.Equal(t, -subtest.frequency, tone.Frequency)
			require.Len(t, tone.HarmonicGains, NumHarmGains)
			for i, harmGain := range tone.HarmonicGains {
				require.Equal(t, wantHarmGains[i], harmGain)
			}
		})
	}
}
