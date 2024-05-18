package note

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_ChordNames tests that the constants for chord names are defined correctly.
func Test_ChordNames(t *testing.T) {
	var name ChordName

	require.Equal(t, ChordName("maj"), Major)
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
		require.False(t, ChordName("").Valid())
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

// Test_NewChord tests that NewChord builds the correct chord depending on the root note and chord name.
func Test_NewChord(t *testing.T) {
	t.Run("invalid root", func(t *testing.T) {
		require.Zero(t, NewChord(Note(""), Major))
		require.Zero(t, NewChord(Note("Z"), Major))
	})

	t.Run("invalid chord", func(t *testing.T) {
		require.Zero(t, NewChord(C, ChordName("")))
		require.Zero(t, NewChord(C, ChordName("invalid")))
	})

	t.Run("invalid root and chord", func(t *testing.T) {
		require.Zero(t, NewChord(A+B, Major+Minor))
	})

	t.Run("valid", func(t *testing.T) {
		require.Equal(t, Chord{C, Major, []Note{C, E, G}}, NewChord(C, Major))
		require.Equal(t, Chord{CSharp, Major6, []Note{CSharp, F, GSharp, ASharp}}, NewChord(CSharp, Major6))
		require.Equal(t, Chord{DFlat, Dom7, []Note{DFlat, F, GSharp, B}}, NewChord(DFlat, Dom7))
		require.Equal(t, Chord{D, Major7, []Note{D, FSharp, A, CSharp}}, NewChord(D, Major7))
		require.Equal(t, Chord{DSharp, Augmented, []Note{DSharp, G, B}}, NewChord(DSharp, Augmented))
		require.Equal(t, Chord{EFlat, Augmented7, []Note{EFlat, G, B, CSharp}}, NewChord(EFlat, Augmented7))
		require.Equal(t, Chord{E, Minor, []Note{E, G, B}}, NewChord(E, Minor))
		require.Equal(t, Chord{F, Minor6, []Note{F, GSharp, C, D}}, NewChord(F, Minor6))
		require.Equal(t, Chord{FSharp, Minor7, []Note{FSharp, A, CSharp, E}}, NewChord(FSharp, Minor7))
		require.Equal(t, Chord{GFlat, MinorMajor7, []Note{GFlat, A, CSharp, F}}, NewChord(GFlat, MinorMajor7))
		require.Equal(t, Chord{G, Diminished, []Note{G, ASharp, CSharp}}, NewChord(G, Diminished))
		require.Equal(t, Chord{GSharp, Diminished7, []Note{GSharp, B, D, F}}, NewChord(GSharp, Diminished7))
		require.Equal(t, Chord{AFlat, HalfDiminished7, []Note{AFlat, B, D, FSharp}}, NewChord(AFlat, HalfDiminished7))
		require.Equal(t, Chord{A, Major, []Note{A, CSharp, E}}, NewChord(A, Major))
		require.Equal(t, Chord{ASharp, Major6, []Note{ASharp, D, F, G}}, NewChord(ASharp, Major6))
		require.Equal(t, Chord{BFlat, Dom7, []Note{BFlat, D, F, GSharp}}, NewChord(BFlat, Dom7))
		require.Equal(t, Chord{B, Major7, []Note{B, DSharp, FSharp, ASharp}}, NewChord(B, Major7))
	})
}

// Test_Chord_Valid tests that Chord's Valid method correctly reports if a chord is valid.
func Test_Chord_Valid(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		require.False(t, Chord{}.Valid())
	})

	t.Run("missing root", func(t *testing.T) {
		require.False(t, Chord{Note(""), Major, []Note{A, B, C}}.Valid())
	})

	t.Run("invalid root", func(t *testing.T) {
		require.False(t, Chord{Note("lkj2"), Major, []Note{A, B, C}}.Valid())
	})

	t.Run("mismatched root", func(t *testing.T) {
		require.False(t, Chord{B, Major, []Note{A, B, C}}.Valid())
		require.False(t, Chord{C, Major, []Note{A, B, C}}.Valid())
		require.False(t, Chord{D, Major, []Note{A, B, C}}.Valid())
	})

	t.Run("missing chord", func(t *testing.T) {
		require.False(t, Chord{A, ChordName(""), []Note{A, B, C}}.Valid())
	})

	t.Run("invalid chord", func(t *testing.T) {
		require.False(t, Chord{A, ChordName("1w9j2"), []Note{A, B, C}}.Valid())
	})

	t.Run("missing notes", func(t *testing.T) {
		require.False(t, Chord{A, Major, []Note{}}.Valid())
		require.False(t, Chord{A, Major, []Note{A}}.Valid())
		require.False(t, Chord{A, Major, []Note{A, B}}.Valid())
	})

	t.Run("invalid notes", func(t *testing.T) {
		require.False(t, Chord{A, Major, []Note{"irfj", B, C}}.Valid())
		require.False(t, Chord{A, Major, []Note{A, "skdj", C}}.Valid())
		require.False(t, Chord{A, Major, []Note{A, B, "234jf"}}.Valid())
		require.False(t, Chord{A, Major, []Note{"109fj", "asdvj", "1238"}}.Valid())
	})

	t.Run("valid", func(t *testing.T) {
		require.True(t, Chord{A, Major, []Note{A, B, C}}.Valid())
		require.True(t, Chord{G, MinorMajor7, []Note{G, A, B, C}}.Valid())
		require.True(t, Chord{FSharp, Minor7, []Note{FSharp, A, CSharp, E}}.Valid())
		require.True(t, NewChord(BFlat, Dom7).Valid())
	})
}

// Test_Chord_String tests that Chord's String method returns the correct string representation of
// the chord.
func Test_Chord_String(t *testing.T) {
	t.Run("invalid chord", func(t *testing.T) {
		require.Equal(t, "invalid chord", Chord{A, Major, nil}.String())
	})

	t.Run("valid chord", func(t *testing.T) {
		require.Equal(t, "Cmaj", NewChord(C, Major).String())
		require.Equal(t, "C♯maj6", NewChord(CSharp, Major6).String())
		require.Equal(t, "D♭dom7", NewChord(DFlat, Dom7).String())
		require.Equal(t, "Dmaj7", NewChord(D, Major7).String())
		require.Equal(t, "D♯aug", NewChord(DSharp, Augmented).String())
		require.Equal(t, "E♭aug7", NewChord(EFlat, Augmented7).String())
		require.Equal(t, "Emin", NewChord(E, Minor).String())
		require.Equal(t, "Fmin6", NewChord(F, Minor6).String())
		require.Equal(t, "F♯min7", NewChord(FSharp, Minor7).String())
		require.Equal(t, "G♭min/maj7", NewChord(GFlat, MinorMajor7).String())
		require.Equal(t, "Gdim", NewChord(G, Diminished).String())
		require.Equal(t, "G♯dim7", NewChord(GSharp, Diminished7).String())
		require.Equal(t, "A♭halfdim7", NewChord(AFlat, HalfDiminished7).String())
		require.Equal(t, "Amin/maj7", NewChord(A, MinorMajor7).String())
		require.Equal(t, "A♯dim", NewChord(ASharp, Diminished).String())
		require.Equal(t, "B♭dim7", NewChord(BFlat, Diminished7).String())
		require.Equal(t, "Bhalfdim7", NewChord(B, HalfDiminished7).String())
	})
}

// Test_Chord_Root tests that Chord's Root method returns the correct root note, or an empty note if
// the chord is invalid.
func Test_Chord_Root(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var chord Chord
		require.Zero(t, chord.Root())
	})

	t.Run("invalid", func(t *testing.T) {
		chord := Chord{Note("asdlkfj"), Major, []Note{A, B, C}}
		require.Zero(t, chord.Root())
	})

	t.Run("valid", func(t *testing.T) {
		chord := NewChord(G, MinorMajor7)
		require.Equal(t, G, chord.Root())
	})
}
