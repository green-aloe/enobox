package note_test

import (
	"fmt"

	"github.com/green-aloe/enobox/note"
)

func ExampleNote_Valid() {
	fmt.Println(note.C.Valid())
	fmt.Println(note.Note("H").Valid())

	// Output:
	// true
	// false
}

func ExampleNote_Frequency() {
	fmt.Println(note.C.Frequency(4))

	// Output:
	// 261.6256
}

func ExampleNote_IncrementBy() {
	note := note.EFlat
	fmt.Println(note.IncrementBy(3), note)
	fmt.Println(note.IncrementBy(-3), note)

	// Output:
	// F♯ E♭
	// C E♭
}
