package note_test

import (
	"fmt"

	"github.com/green-aloe/enobox/note"
)

func ExampleNote_Valid() {
	note1 := note.C
	valid1 := note1.Valid()

	note2 := note.Note("X")
	valid2 := note2.Valid()

	fmt.Println(valid1, valid2)

	// Output:
	// true false
}

func ExampleNote_Frequency() {
	note := note.C
	octave := 4
	freq := note.Frequency(octave)

	fmt.Println(freq)

	// Output:
	// 261.6256
}

func ExampleNote_IncrementBy() {
	note1 := note.EFlat
	note2 := note1.IncrementBy(3)
	note3 := note1.IncrementBy(-3)

	fmt.Println(note1, note2, note3)

	// Output:
	// E♭ F♯ C
}
