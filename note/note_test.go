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

// Test_Note_Frequency tests that Note's Frequency method returns the correct frequency for various
// notes and octaves.
func Test_Note_Frequency(t *testing.T) {
	t.Run("invalid note", func(t *testing.T) {
		require.Zero(t, Note("").Frequency(0))
		require.Zero(t, Note("asglkaj3rqjw").Frequency(0))
		require.Zero(t, (A + B).Frequency(0))
	})

	t.Run("invalid octave", func(t *testing.T) {
		require.Zero(t, C.Frequency(-2))
		require.Zero(t, C.Frequency(11))
	})

	t.Run("valid", func(t *testing.T) {
		for i, wantFrequency := range []float32{
			8.175799, 16.35160, 32.70320, 65.40639, 130.8128, 261.6256,
			523.2511, 1046.502, 2093.005, 4186.009, 8372.018, 16744.04,
		} {
			require.Equal(t, wantFrequency, (C).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			8.661957, 17.32391, 34.64783, 69.29566, 138.5913, 277.1826,
			554.3653, 1108.731, 2217.461, 4434.922, 8869.844, 17739.69,
		} {
			require.Equal(t, wantFrequency, (CSharp).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			8.661957, 17.32391, 34.64783, 69.29566, 138.5913, 277.1826,
			554.3653, 1108.731, 2217.461, 4434.922, 8869.844, 17739.69,
		} {
			require.Equal(t, wantFrequency, (DFlat).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			9.177024, 18.35405, 36.70810, 73.41619, 146.8324, 293.6648,
			587.3295, 1174.659, 2349.318, 4698.636, 9397.273, 18794.55,
		} {
			require.Equal(t, wantFrequency, (D).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			9.722718, 19.44544, 38.89087, 77.78175, 155.5635, 311.1270,
			622.2540, 1244.508, 2489.016, 4978.032, 9956.063, 19912.13,
		} {
			require.Equal(t, wantFrequency, (DSharp).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			9.722718, 19.44544, 38.89087, 77.78175, 155.5635, 311.1270,
			622.2540, 1244.508, 2489.016, 4978.032, 9956.063, 19912.13,
		} {
			require.Equal(t, wantFrequency, (EFlat).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			10.30086, 20.60172, 41.20344, 82.40689, 164.8138, 329.6276,
			659.2551, 1318.510, 2637.020, 5274.041, 10548.08, 21096.16,
		} {
			require.Equal(t, wantFrequency, (E).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			10.91338, 21.82676, 43.65353, 87.30706, 174.6141, 349.2282,
			698.4565, 1396.913, 2793.826, 5587.652, 11175.30, 22350.61,
		} {
			require.Equal(t, wantFrequency, (F).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			11.56233, 23.12465, 46.24930, 92.49861, 184.9972, 369.9944,
			739.9888, 1479.978, 2959.955, 5919.911, 11839.82, 23679.64,
		} {
			require.Equal(t, wantFrequency, (FSharp).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			11.56233, 23.12465, 46.24930, 92.49861, 184.9972, 369.9944,
			739.9888, 1479.978, 2959.955, 5919.911, 11839.82, 23679.64,
		} {
			require.Equal(t, wantFrequency, (GFlat).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			12.24986, 24.49971, 48.99943, 97.99886, 195.9977, 391.9954,
			783.9909, 1567.982, 3135.963, 6271.927, 12543.85, 25087.71,
		} {
			require.Equal(t, wantFrequency, (G).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			12.97827, 25.95654, 51.91309, 103.8262, 207.6523, 415.3047,
			830.6094, 1661.219, 3322.438, 6644.875, 13289.75, 26579.50,
		} {
			require.Equal(t, wantFrequency, (GSharp).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			12.97827, 25.95654, 51.91309, 103.8262, 207.6523, 415.3047,
			830.6094, 1661.219, 3322.438, 6644.875, 13289.75, 26579.50,
		} {
			require.Equal(t, wantFrequency, (AFlat).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			13.75000, 27.50000, 55.00000, 110.0000, 220.0000, 440.0000,
			880.0000, 1760.000, 3520.000, 7040.000, 14080.00, 28160.00,
		} {
			require.Equal(t, wantFrequency, (A).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			14.56762, 29.13524, 58.27047, 116.5409, 233.0819, 466.1638,
			932.3275, 1864.655, 3729.310, 7458.620, 14917.24, 29834.48,
		} {
			require.Equal(t, wantFrequency, (ASharp).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			14.56762, 29.13524, 58.27047, 116.5409, 233.0819, 466.1638,
			932.3275, 1864.655, 3729.310, 7458.620, 14917.24, 29834.48,
		} {
			require.Equal(t, wantFrequency, (BFlat).Frequency(-1+i))
		}

		for i, wantFrequency := range []float32{
			15.43385, 30.86771, 61.73541, 123.4708, 246.9417, 493.8833,
			987.7666, 1975.533, 3951.066, 7902.133, 15804.27, 31608.53,
		} {
			require.Equal(t, wantFrequency, (B).Frequency(-1+i))
		}
	})
}

// Test_Note_IncrementBy tests that Note's IncrementBy method correctly increments or decrements a
// note by a given number of half steps.
func Test_Note_IncrementBy(t *testing.T) {
	t.Run("invalid note", func(t *testing.T) {
		require.Zero(t, Note("").IncrementBy(0))
		require.Zero(t, Note("aslkjql34kfj").IncrementBy(1))
		require.Zero(t, Note(B+"sharp").IncrementBy(-1))
	})

	t.Run("no shift", func(t *testing.T) {
		require.Equal(t, C, C.IncrementBy(0))
		require.Equal(t, CSharp, CSharp.IncrementBy(0))
		require.Equal(t, DFlat, DFlat.IncrementBy(0))
		require.Equal(t, D, D.IncrementBy(0))
		require.Equal(t, DSharp, DSharp.IncrementBy(0))
		require.Equal(t, EFlat, EFlat.IncrementBy(0))
		require.Equal(t, E, E.IncrementBy(0))
		require.Equal(t, F, F.IncrementBy(0))
		require.Equal(t, FSharp, FSharp.IncrementBy(0))
		require.Equal(t, GFlat, GFlat.IncrementBy(0))
		require.Equal(t, G, G.IncrementBy(0))
		require.Equal(t, GSharp, GSharp.IncrementBy(0))
		require.Equal(t, AFlat, AFlat.IncrementBy(0))
		require.Equal(t, A, A.IncrementBy(0))
		require.Equal(t, ASharp, ASharp.IncrementBy(0))
		require.Equal(t, BFlat, BFlat.IncrementBy(0))
		require.Equal(t, B, B.IncrementBy(0))
	})

	t.Run("positive shift", func(t *testing.T) {
		for _, wantNotes := range [][]Note{
			{C, CSharp, D, DSharp, E, F, FSharp, G, GSharp, A, ASharp, B},
			{CSharp, D, DSharp, E, F, FSharp, G, GSharp, A, ASharp, B, C},
			{DFlat, D, DSharp, E, F, FSharp, G, GSharp, A, ASharp, B, C},
			{D, DSharp, E, F, FSharp, G, GSharp, A, ASharp, B, C, CSharp},
			{DSharp, E, F, FSharp, G, GSharp, A, ASharp, B, C, CSharp, D},
			{EFlat, E, F, FSharp, G, GSharp, A, ASharp, B, C, CSharp, D},
			{E, F, FSharp, G, GSharp, A, ASharp, B, C, CSharp, D, DSharp},
			{F, FSharp, G, GSharp, A, ASharp, B, C, CSharp, D, DSharp, E},
			{FSharp, G, GSharp, A, ASharp, B, C, CSharp, D, DSharp, E, F},
			{GFlat, G, GSharp, A, ASharp, B, C, CSharp, D, DSharp, E, F},
			{G, GSharp, A, ASharp, B, C, CSharp, D, DSharp, E, F, FSharp},
			{GSharp, A, ASharp, B, C, CSharp, D, DSharp, E, F, FSharp, G},
			{AFlat, A, ASharp, B, C, CSharp, D, DSharp, E, F, FSharp, G},
			{A, ASharp, B, C, CSharp, D, DSharp, E, F, FSharp, G, GSharp},
			{ASharp, B, C, CSharp, D, DSharp, E, F, FSharp, G, GSharp, A},
			{BFlat, B, C, CSharp, D, DSharp, E, F, FSharp, G, GSharp, A},
			{B, C, CSharp, D, DSharp, E, F, FSharp, G, GSharp, A, ASharp},
		} {
			base := wantNotes[0]
			for multiplier := range 10 {
				for i, wantNote := range wantNotes {
					require.Equal(t, wantNote, base.IncrementBy((12*multiplier)+i))
				}
			}
		}
	})

	t.Run("negative shift", func(t *testing.T) {
		for _, wantNotes := range [][]Note{
			{C, B, ASharp, A, GSharp, G, FSharp, F, E, DSharp, D, CSharp},
			{CSharp, C, B, ASharp, A, GSharp, G, FSharp, F, E, DSharp, D},
			{DFlat, C, B, ASharp, A, GSharp, G, FSharp, F, E, DSharp, D},
			{D, CSharp, C, B, ASharp, A, GSharp, G, FSharp, F, E, DSharp},
			{DSharp, D, CSharp, C, B, ASharp, A, GSharp, G, FSharp, F, E},
			{EFlat, D, CSharp, C, B, ASharp, A, GSharp, G, FSharp, F, E},
			{E, DSharp, D, CSharp, C, B, ASharp, A, GSharp, G, FSharp, F},
			{F, E, DSharp, D, CSharp, C, B, ASharp, A, GSharp, G, FSharp},
			{FSharp, F, E, DSharp, D, CSharp, C, B, ASharp, A, GSharp, G},
			{GFlat, F, E, DSharp, D, CSharp, C, B, ASharp, A, GSharp, G},
			{G, FSharp, F, E, DSharp, D, CSharp, C, B, ASharp, A, GSharp},
			{GSharp, G, FSharp, F, E, DSharp, D, CSharp, C, B, ASharp, A},
			{AFlat, G, FSharp, F, E, DSharp, D, CSharp, C, B, ASharp, A},
			{A, GSharp, G, FSharp, F, E, DSharp, D, CSharp, C, B, ASharp},
			{ASharp, A, GSharp, G, FSharp, F, E, DSharp, D, CSharp, C, B},
			{BFlat, A, GSharp, G, FSharp, F, E, DSharp, D, CSharp, C, B},
			{B, ASharp, A, GSharp, G, FSharp, F, E, DSharp, D, CSharp, C},
		} {
			base := wantNotes[0]
			for multiplier := range 10 {
				for i, wantNote := range wantNotes {
					require.Equal(t, wantNote, base.IncrementBy(-((12 * multiplier) + i)))
				}
			}
		}
	})
}
