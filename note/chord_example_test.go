package note_test

import (
	"fmt"

	"github.com/green-aloe/enobox/note"
)

func ExampleChordName_Valid() {
	validChordName := note.Major7
	invalidChordName := note.ChordName("invalid chord")

	fmt.Println(validChordName.Valid(), invalidChordName.Valid())

	// Output:
	// true false
}

func ExampleNewChord() {
	validChord1 := note.NewChord(note.C, note.Major)
	validChord2 := note.NewChord(note.ASharp, note.Augmented7)

	invalidChord := note.NewChord(note.Note("bad note"), note.Diminished)

	fmt.Println(validChord1)
	fmt.Println(validChord2)
	fmt.Println(invalidChord)

	// Output:
	// Cmaj
	// A♯aug7
	// invalid chord
}

func ExampleChord_Valid() {
	validChord := note.NewChord(note.DFlat, note.Dom7)
	invalidChord := note.NewChord(note.C, note.ChordName("invalid chord"))

	fmt.Println(validChord.Valid(), invalidChord.Valid())

	// Output:
	// true false
}

func ExampleChord_String() {
	validChord1 := note.NewChord(note.G, note.Minor)
	validChord2 := note.NewChord(note.D, note.MinorMajor7)

	invalidChord := note.NewChord(note.Note("bad note"), note.Major)

	fmt.Println(validChord1)
	fmt.Println(validChord2)
	fmt.Println(invalidChord)

	// Output:
	// Gmin
	// Dmin/maj7
	// invalid chord
}

func ExampleChord_Root() {
	chord := note.NewChord(note.A, note.Minor7)
	fmt.Println(chord.Root())

	// Output:
	// A
}

func ExampleChord_Name() {
	chord := note.NewChord(note.EFlat, note.HalfDiminished7)
	fmt.Println(chord.Name())

	// Output:
	// halfdim7
}

func ExampleChord_Notes() {
	chord := note.NewChord(note.FSharp, note.Minor6)
	fmt.Println(chord.Notes())

	// Output:
	// [F♯ A C♯ D♯]
}
