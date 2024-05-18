package note

const (
	Major           ChordName = "maj"      // Major chord
	Major6          ChordName = "maj6"     // Major sixth chord
	Dom7            ChordName = "dom7"     // Dominant seventh chord
	Major7          ChordName = "maj7"     // Major seventh chord
	Augmented       ChordName = "aug"      // Augmented chord
	Augmented7      ChordName = "aug7"     // Augmented seventh chord
	Minor           ChordName = "min"      // Minor chord
	Minor6          ChordName = "min6"     // Minor sixth chord
	Minor7          ChordName = "min7"     // Minor seventh chord
	MinorMajor7     ChordName = "min/maj7" // Minor-major seventh chord
	Diminished      ChordName = "dim"      // Diminished chord
	Diminished7     ChordName = "dim7"     // Diminished seventh chord
	HalfDiminished7 ChordName = "halfdim7" // Half-diminished seventh chord
)

// A chord name is a pre-defined name of a chord.
type ChordName string

// Valid reports if the chord name is valid.
func (name ChordName) Valid() bool {
	_, ok := chordToSemitonesList[name]
	return ok
}

// A Chord is a list of three or more notes that starts with a base note and goes up in ascending order.
type Chord struct {
	root  Note
	name  ChordName
	notes []Note
}

// NewChord creates a new chord (list of notes) from a root note and a pre-defined chord name. If
// the root note or chord name is invalid, this returns an empty chord.
func NewChord(root Note, name ChordName) Chord {
	if !root.Valid() || !name.Valid() {
		return Chord{}
	}

	var notes []Note
	for _, semitone := range chordToSemitonesList[name] {
		notes = append(notes, root.IncrementBy(semitone))
	}

	return Chord{
		root:  root,
		name:  name,
		notes: notes,
	}
}

// Valid reports if the chord is valid.
func (c Chord) Valid() bool {
	if !c.root.Valid() || !c.name.Valid() || len(c.notes) < 3 || c.root != c.notes[0] {
		return false
	}

	for _, note := range c.notes {
		if !note.Valid() {
			return false
		}
	}

	return true
}

// String returns the string representation of the chord. If the chord is invalid, this returns
// "invalid chord".
func (c Chord) String() string {
	if !c.Valid() {
		return "invalid chord"
	}

	return string(c.root) + string(c.name)
}
