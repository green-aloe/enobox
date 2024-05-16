package note

const (
	Major           ChordName = ""         // Major chord
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

type ChordName string
