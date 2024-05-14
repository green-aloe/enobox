package note

import (
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

// Test_Note_Valid tests that Note's Valid method correctly reports if a note is valid.
func Test_Note_Valid(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		require.True(t, C.Valid())
		require.True(t, CSharp.Valid())
		require.True(t, DFlat.Valid())
		require.True(t, D.Valid())
		require.True(t, DSharp.Valid())
		require.True(t, EFlat.Valid())
		require.True(t, E.Valid())
		require.True(t, F.Valid())
		require.True(t, FSharp.Valid())
		require.True(t, GFlat.Valid())
		require.True(t, G.Valid())
		require.True(t, GSharp.Valid())
		require.True(t, AFlat.Valid())
		require.True(t, A.Valid())
		require.True(t, ASharp.Valid())
		require.True(t, BFlat.Valid())
		require.True(t, B.Valid())

	})

	t.Run("accidentals", func(t *testing.T) {
		require.True(t, (C + Sharp).Valid())
		require.True(t, (D + Flat).Valid())
		require.True(t, (D + Sharp).Valid())
		require.True(t, (E + Flat).Valid())
		require.True(t, (F + Sharp).Valid())
		require.True(t, (G + Flat).Valid())
		require.True(t, (G + Sharp).Valid())
		require.True(t, (A + Flat).Valid())
		require.True(t, (A + Sharp).Valid())
		require.True(t, (B + Flat).Valid())

		require.False(t, (C + Flat).Valid())
		require.False(t, (E + Sharp).Valid())
		require.False(t, (F + Flat).Valid())
		require.False(t, (B + Sharp).Valid())

		require.False(t, Note(Sharp).Valid())
		require.False(t, Note(Flat).Valid())
	})

	t.Run("empty", func(t *testing.T) {
		require.False(t, Note("").Valid())
	})

	t.Run("gibberish", func(t *testing.T) {
		require.False(t, Note("aklsdjf").Valid())
	})

	t.Run("misspellings", func(t *testing.T) {
		require.False(t, (A + B).Valid())
		require.False(t, (A + "b").Valid())
		require.False(t, (C + "sharp").Valid())
	})
}
