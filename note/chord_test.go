package note

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_ChordNames tests that the constants for chord names are defined correctly.
func Test_ChordNames(t *testing.T) {
	var name ChordName

	require.Equal(t, ChordName(""), Major)
	require.IsType(t, name, Major)
	require.Equal(t, ChordName("maj6"), Major6)
	require.IsType(t, name, Major6)
	require.Equal(t, ChordName("dom7"), Dom7)
	require.IsType(t, name, Dom7)
	require.Equal(t, ChordName("maj7"), Major7)
	require.IsType(t, name, Major7)
	require.Equal(t, ChordName("aug"), Augmented)
	require.IsType(t, name, Augmented)
	require.Equal(t, ChordName("aug7"), Augmented7)
	require.IsType(t, name, Augmented7)
	require.Equal(t, ChordName("min"), Minor)
	require.IsType(t, name, Minor)
	require.Equal(t, ChordName("min6"), Minor6)
	require.IsType(t, name, Minor6)
	require.Equal(t, ChordName("min7"), Minor7)
	require.IsType(t, name, Minor7)
	require.Equal(t, ChordName("min/maj7"), MinorMajor7)
	require.IsType(t, name, MinorMajor7)
	require.Equal(t, ChordName("dim"), Diminished)
	require.IsType(t, name, Diminished)
	require.Equal(t, ChordName("dim7"), Diminished7)
	require.IsType(t, name, Diminished7)
	require.Equal(t, ChordName("halfdim7"), HalfDiminished7)
	require.IsType(t, name, HalfDiminished7)
}
