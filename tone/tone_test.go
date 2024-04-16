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
}
