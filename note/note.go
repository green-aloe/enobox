package note

const (
	Sharp = "♯"
	Flat  = "♭"

	C      Note = "C"
	CSharp Note = C + Sharp
	DFlat  Note = D + Flat
	D      Note = "D"
	DSharp Note = D + Sharp
	EFlat  Note = E + Flat
	E      Note = "E"
	F      Note = "F"
	FSharp Note = F + Sharp
	GFlat  Note = G + Flat
	G      Note = "G"
	GSharp Note = G + Sharp
	AFlat  Note = A + Flat
	A      Note = "A"
	ASharp Note = A + Sharp
	BFlat  Note = B + Flat
	B      Note = "B"
)

type Note string

// Valid reports if the note is valid.
func (note Note) Valid() bool {
	_, ok := noteToSemitonesAboveC[note]
	return ok
}
