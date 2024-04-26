package tone

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_Notes tests that the constants for notes and accidentals are defined correctly.
func Test_Notes(t *testing.T) {
	t.Run("accidentals", func(t *testing.T) {
		var s string
		require.Equal(t, "♯", Sharp)
		require.IsType(t, s, Sharp)
		require.Equal(t, "♭", Flat)
		require.IsType(t, s, Flat)
	})

	t.Run("notes", func(t *testing.T) {
		var note Note
		require.Equal(t, Note("C"), C)
		require.IsType(t, note, C)
		require.Equal(t, Note("C♯"), CSharp)
		require.IsType(t, note, CSharp)
		require.Equal(t, Note("D♭"), DFlat)
		require.IsType(t, note, DFlat)
		require.Equal(t, Note("D"), D)
		require.IsType(t, note, D)
		require.Equal(t, Note("D♯"), DSharp)
		require.IsType(t, note, DSharp)
		require.Equal(t, Note("E♭"), EFlat)
		require.IsType(t, note, EFlat)
		require.Equal(t, Note("E"), E)
		require.IsType(t, note, E)
		require.Equal(t, Note("F"), F)
		require.IsType(t, note, F)
		require.Equal(t, Note("F♯"), FSharp)
		require.IsType(t, note, FSharp)
		require.Equal(t, Note("G♭"), GFlat)
		require.IsType(t, note, GFlat)
		require.Equal(t, Note("G"), G)
		require.IsType(t, note, G)
		require.Equal(t, Note("G♯"), GSharp)
		require.IsType(t, note, GSharp)
		require.Equal(t, Note("A♭"), AFlat)
		require.IsType(t, note, AFlat)
		require.Equal(t, Note("A"), A)
		require.IsType(t, note, A)
		require.Equal(t, Note("A♯"), ASharp)
		require.IsType(t, note, ASharp)
		require.Equal(t, Note("B♭"), BFlat)
		require.IsType(t, note, BFlat)
		require.Equal(t, Note("B"), B)
		require.IsType(t, note, B)
	})
}

// Test_NewToneFrom tests that NewToneFrom returns a tone that has been initialized with the correct
// fundamental frequency for various notes and octaves.
func Test_NewToneFrom(t *testing.T) {

	t.Run("invalid note", func(t *testing.T) {
		tone := NewToneFrom(Note("H"), 5)
		require.Empty(t, tone)

		tone = NewToneFrom(Note(C+"b"), 5)
		require.Empty(t, tone)
	})

	t.Run("invalid octave", func(t *testing.T) {
		tone := NewToneFrom(C, -2)
		require.Empty(t, tone)

		tone = NewToneFrom(C, 11)
		require.Empty(t, tone)
	})

	for _, note := range []Note{C, CSharp, DFlat, D, DSharp, EFlat, E, F, FSharp, GFlat, G, GSharp, AFlat, A, ASharp, BFlat, B} {
		for octave := -1; octave <= 10; octave++ {
			t.Run(fmt.Sprintf("%v%v", note, octave), func(t *testing.T) {
				tone := NewToneFrom(note, octave)
				require.NotEmpty(t, tone)
				require.IsType(t, Tone{}, tone)
				require.Greater(t, tone.Frequency, float32(8))
				require.Equal(t, noteFreqs[note][octave], tone.Frequency)
				require.Equal(t, float32(0), tone.Gain)
				require.Len(t, tone.HarmonicGains, NumHarmGains)
				for _, gain := range tone.HarmonicGains {
					require.Equal(t, float32(0), gain)
				}
			})
		}
	}
}
