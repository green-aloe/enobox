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

// Test_ChordName_Valid tests that ChordName's Valid method correctly reports if a chord name is valid.
func Test_ChordName_Valid(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		require.True(t, Major.Valid())
		require.True(t, Major6.Valid())
		require.True(t, Dom7.Valid())
		require.True(t, Major7.Valid())
		require.True(t, Augmented.Valid())
		require.True(t, Augmented7.Valid())
		require.True(t, Minor.Valid())
		require.True(t, Minor6.Valid())
		require.True(t, Minor7.Valid())
		require.True(t, MinorMajor7.Valid())
		require.True(t, Diminished.Valid())
		require.True(t, Diminished7.Valid())
		require.True(t, HalfDiminished7.Valid())
	})

	t.Run("empty", func(t *testing.T) {
		require.True(t, ChordName("").Valid())
	})

	t.Run("gibberish", func(t *testing.T) {
		require.False(t, ChordName("l4tjgq3").Valid())
	})

	t.Run("misspellings", func(t *testing.T) {
		require.False(t, ChordName("Maj6").Valid())
		require.False(t, (Minor + Diminished).Valid())
		require.False(t, (Diminished + "6").Valid())
	})
}
