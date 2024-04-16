package tone

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_Consts tests any package-level constant values.
func Test_Consts(t *testing.T) {
	require.Equal(t, 20, NumHarmGains)
}

// Test_NewTone tests that NewTone returns a Tone that has been initialized correctly.
func Test_NewTone(t *testing.T) {
	tone := NewTone()
	require.NotEmpty(t, tone)
	require.IsType(t, Tone{}, tone)
	require.Equal(t, float32(0), tone.Frequency)
	require.Equal(t, float32(0), tone.Gain)
	require.Len(t, tone.HarmonicGains, NumHarmGains)
	for _, gain := range tone.HarmonicGains {
		require.Equal(t, float32(0), gain)
	}
}

// Test_NewToneAt tests that NewToneAt returns a Tone that has been initialized with the correct
// fundamental frequency.
func Test_NewToneAt(t *testing.T) {
	type subtest struct {
		frequency float32
		name      string
	}

	subtests := []subtest{
		{0, "zero frequency"},
		{-10, "negative frequency"},
		{10, "positive frequency"},
		{23.1, "non-integer frequency"},
		{440, "A4 frequency"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			tone := NewToneAt(subtest.frequency)
			require.NotEmpty(t, tone)
			require.IsType(t, Tone{}, tone)
			require.Equal(t, subtest.frequency, tone.Frequency)
			require.Equal(t, float32(0), tone.Gain)
			require.Len(t, tone.HarmonicGains, NumHarmGains)
			for _, gain := range tone.HarmonicGains {
				require.Equal(t, float32(0), gain)
			}
		})
	}
}
