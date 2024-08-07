package tone

import (
	"fmt"
	"testing"

	"github.com/green-aloe/enobox/context"
	"github.com/green-aloe/enobox/note"
	"github.com/stretchr/testify/require"
)

// Test_NewTone tests that NewTone returns a tone that has been initialized correctly.
func Test_NewTone(t *testing.T) {
	defer SetNumHarmGains(DefaultNumHarmGains)

	for _, numHarmGains := range []int{0, 1, 100, DefaultNumHarmGains} {
		SetNumHarmGains(numHarmGains)

		ctx := context.NewContext()

		tone := NewTone(ctx)
		require.NotEmpty(t, tone)
		require.IsType(t, Tone{}, tone)
		require.Zero(t, tone.Frequency)
		require.Zero(t, tone.Gain)
		require.Len(t, tone.HarmonicGains, numHarmGains)
		for _, gain := range tone.HarmonicGains {
			require.Zero(t, gain)
		}
	}
}

// Test_NewToneAt tests that NewToneAt returns a tone that has been initialized with the correct
// fundamental frequency.
func Test_NewToneAt(t *testing.T) {
	defer SetNumHarmGains(DefaultNumHarmGains)

	type subtest struct {
		frequency float32
		name      string
	}

	subtests := []subtest{
		{0, "zero frequency"},
		{-10, "negative frequency"},
		{10, "positive frequency"},
		{23.1, "non-integer frequency"},
		{440, "A4 frequency"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			for _, numHarmGains := range []int{0, 1, 100, DefaultNumHarmGains} {
				SetNumHarmGains(numHarmGains)

				ctx := context.NewContext()

				tone := NewToneAt(ctx, subtest.frequency)
				require.NotEmpty(t, tone)
				require.IsType(t, Tone{}, tone)
				require.Equal(t, subtest.frequency, tone.Frequency)
				require.Zero(t, tone.Gain)
				require.Len(t, tone.HarmonicGains, numHarmGains)
				for _, gain := range tone.HarmonicGains {
					require.Zero(t, gain)
				}
			}
		})
	}
}

// Test_NewToneFrom tests that NewToneFrom returns a tone that has been initialized with the correct
// fundamental frequency for various notes and octaves.
func Test_NewToneFrom(t *testing.T) {
	defer SetNumHarmGains(DefaultNumHarmGains)

	t.Run("invalid note", func(t *testing.T) {
		ctx := context.NewContext()

		tone := NewToneFrom(ctx, note.Note("H"), 5)
		require.Equal(t, NewTone(ctx), tone)

		tone = NewToneFrom(ctx, note.Note(note.C+"b"), 5)
		require.Equal(t, NewTone(ctx), tone)
	})

	t.Run("invalid octave", func(t *testing.T) {
		ctx := context.NewContext()

		tone := NewToneFrom(ctx, note.C, -2)
		require.Equal(t, NewTone(ctx), tone)

		tone = NewToneFrom(ctx, note.C, 11)
		require.Equal(t, NewTone(ctx), tone)
	})

	for _, note := range []note.Note{
		note.C, note.CSharp, note.DFlat, note.D, note.DSharp, note.EFlat, note.E,
		note.F, note.FSharp, note.GFlat, note.G, note.GSharp, note.AFlat, note.A,
		note.ASharp, note.BFlat, note.B,
	} {
		for octave := -1; octave <= 10; octave++ {
			t.Run(fmt.Sprintf("%v%v", note, octave), func(t *testing.T) {
				for _, numHarmGains := range []int{0, 1, 100, DefaultNumHarmGains} {
					SetNumHarmGains(numHarmGains)

					ctx := context.NewContext()

					tone := NewToneFrom(ctx, note, octave)
					require.NotEmpty(t, tone)
					require.IsType(t, Tone{}, tone)
					require.Greater(t, tone.Frequency, float32(8))
					require.Equal(t, note.Frequency(octave), tone.Frequency)
					require.Zero(t, tone.Gain)
					require.Len(t, tone.HarmonicGains, numHarmGains)
					for _, gain := range tone.HarmonicGains {
						require.Zero(t, gain)
					}
				}
			})
		}
	}
}

// Test_NewToneWith tests that NewToneWith returns a tone that has been initialized with the correct
// fundamental frequency, gain, and harmonic gains.
func Test_NewToneWith(t *testing.T) {
	defer SetNumHarmGains(DefaultNumHarmGains)

	t.Run("zero length", func(t *testing.T) {
		ctx := context.NewContext()
		require.Equal(t, 20, NumHarmGains(ctx))

		tone := NewToneWith(ctx, 123.456, 0.789, nil)

		require.IsType(t, Tone{}, tone)
		require.Equal(t, float32(123.456), tone.Frequency)
		require.Equal(t, float32(0.789), tone.Gain)
		require.Equal(t, []float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, tone.HarmonicGains)
		require.Equal(t, 20, cap(tone.HarmonicGains))
	})

	t.Run("short length, short capacity", func(t *testing.T) {
		ctx := context.NewContext()
		require.Equal(t, 20, NumHarmGains(ctx))

		gains := make([]float32, 5, 8)
		for i := range gains {
			gains[i] = float32(i + 1)
		}
		tone := NewToneWith(ctx, 123.456, 0.789, gains)

		require.IsType(t, Tone{}, tone)
		require.Equal(t, float32(123.456), tone.Frequency)
		require.Equal(t, float32(0.789), tone.Gain)
		require.Equal(t, []float32{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, tone.HarmonicGains)
		require.Equal(t, 20, cap(tone.HarmonicGains))

		// We needed more harmonic gains than were provided. The provided slice was too short, and
		// it did not have enough capacity to grow in place. We should have allocated a new slice
		// and used a new backing array.
		for i := range gains {
			gains[i] += 100
		}
		require.Equal(t, []float32{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, tone.HarmonicGains)
	})

	t.Run("short length, long capacity", func(t *testing.T) {
		ctx := context.NewContext()
		require.Equal(t, 20, NumHarmGains(ctx))

		gains := make([]float32, 5, 30)
		for i := range gains {
			gains[i] = float32(i + 1)
		}
		tone := NewToneWith(ctx, 123.456, 0.789, gains)

		require.IsType(t, Tone{}, tone)
		require.Equal(t, float32(123.456), tone.Frequency)
		require.Equal(t, float32(0.789), tone.Gain)
		require.Equal(t, []float32{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, tone.HarmonicGains)
		require.Equal(t, 30, cap(tone.HarmonicGains))

		// We needed more harmonic gains than were provided. The provided slice was too short, but
		// it had enough capacity to grow in place. We should not have allocated a new slice and
		// should be using the same backing array.
		for i := range gains {
			gains[i] += 100
		}
		require.Equal(t, []float32{101, 102, 103, 104, 105, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, tone.HarmonicGains)
	})

	t.Run("expected length", func(t *testing.T) {
		ctx := context.NewContext()
		require.Equal(t, 20, NumHarmGains(ctx))

		gains := make([]float32, 20)
		for i := range gains {
			gains[i] = float32(i + 1)
		}
		tone := NewToneWith(ctx, 123.456, 0.789, gains)

		require.IsType(t, Tone{}, tone)
		require.Equal(t, float32(123.456), tone.Frequency)
		require.Equal(t, float32(0.789), tone.Gain)
		require.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, tone.HarmonicGains)
		require.Equal(t, 20, cap(tone.HarmonicGains))

		// We did not need to change anything about the slice of harmonic gains provided. We should
		// still be using the same backing array.
		for i := range gains {
			gains[i] += 100
		}
		require.Equal(t, []float32{101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120}, tone.HarmonicGains)
	})

	t.Run("long length", func(t *testing.T) {
		ctx := context.NewContext()
		require.Equal(t, 20, NumHarmGains(ctx))

		gains := make([]float32, 30)
		for i := range gains {
			gains[i] = float32(i + 1)
		}
		tone := NewToneWith(ctx, 123.456, 0.789, gains)

		require.IsType(t, Tone{}, tone)
		require.Equal(t, float32(123.456), tone.Frequency)
		require.Equal(t, float32(0.789), tone.Gain)
		require.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, tone.HarmonicGains)
		require.Equal(t, 30, cap(tone.HarmonicGains))

		// We needed to shorten the slice of harmonic gains provided, but the backing array should
		// not have changed.
		for i := range gains {
			gains[i] += 100
		}
		require.Equal(t, []float32{101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120}, tone.HarmonicGains)
	})

	for _, frequency := range []float32{0, -10, 10, 23.1, 440} {
		for _, gain := range []float32{0, 0.789, 1} {
			for _, numHarmGains := range []int{0, 1, DefaultNumHarmGains, 100} {
				SetNumHarmGains(numHarmGains)

				ctx := context.NewContext()

				gains := make([]float32, numHarmGains)
				for i := range gains {
					gains[i] = float32(i + 1)
				}
				tone := NewToneWith(ctx, frequency, gain, gains)

				require.IsType(t, Tone{}, tone)
				require.Equal(t, frequency, tone.Frequency)
				require.Equal(t, gain, tone.Gain)
				require.Len(t, tone.HarmonicGains, numHarmGains)
			}
		}
	}
}

// Test_Tone_HarmonicFreq tests that Tone's HarmonicFreq method returns the correct frequency for a
// variety of tones and harmonics.
func Test_Tone_HarmonicFreq(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		for i := range 10 {
			var tone *Tone
			require.Zero(t, tone.HarmonicFreq(i))
		}
	})

	ctx := context.NewContext()

	type subtest struct {
		tone Tone
		n    int
		want float32
		name string
	}

	subtests := []subtest{
		{Tone{}, 0, 0, "empty tone, no harmonic"},
		{Tone{}, -1, 0, "empty tone, negative harmonic"},
		{Tone{}, 1, 0, "empty tone, positive harmonic"},
		{NewTone(ctx), 0, 0, "new tone, no harmonic"},
		{NewTone(ctx), -1, 0, "new tone, negative harmonic"},
		{NewTone(ctx), 1, 0, "new tone, positive harmonic"},
		{NewToneAt(ctx, 0), 0, 0, "zero frequency, no harmonic"},
		{NewToneAt(ctx, 0), -10, 0, "zero frequency, negative harmonic"},
		{NewToneAt(ctx, 0), 10, 0, "zero frequency, positive harmonic"},
		{NewToneAt(ctx, -10), 0, 0, "negative frequency, no harmonic"},
		{NewToneAt(ctx, -10), -1, 0, "negative frequency, negative harmonic"},
		{NewToneAt(ctx, -10), 2, 0, "negative frequency, positive harmonic to zero"},
		{NewToneAt(ctx, -10), 9, 70, "negative frequency, positive harmonic"},
		{NewToneAt(ctx, 440), 0, 0, "positive frequency, no harmonic"},
		{NewToneAt(ctx, 440), -2, 0, "positive frequency, negative harmonic"},
		{NewToneAt(ctx, 440), 5, 2200, "positive frequency, positive harmonic"},
		{NewToneAt(ctx, 23.1), 0, 0, "non-integer frequency, no harmonic"},
		{NewToneAt(ctx, 587.3295), -1, 0, "non-integer frequency, negative harmonic"},
		{NewToneAt(ctx, 87.30706), 9, 785.763, "non-integer frequency, positive harmonic"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			have := subtest.tone.HarmonicFreq(subtest.n)
			require.Equal(t, subtest.want, have)
		})
	}

	t.Run("table", func(t *testing.T) {
		fundFreq := float32(34.6478)
		wantsNeg := []float32{
			-34.6478, 0, 34.6478, 69.2956, 103.943, 138.591, 173.239, 207.886, 242.534, 277.182, 311.830, 346.478, 381.125, 415.773, 450.421, 485.069, 519.717, 554.364, 589.012, 623.660, 658.308, 692.956, 727.603, 762.251, 796.899, 831.547, 866.195, 900.842, 935.490, 970.138, 1004.78, 1039.43, 1074.08, 1108.72, 1143.37, 1178.02, 1212.67, 1247.32, 1281.96, 1316.61, 1351.26, 1385.91,
			1420.55, 1455.20, 1489.85, 1524.50, 1559.15, 1593.79, 1628.44, 1663.09, 1697.74, 1732.39, 1767.03, 1801.68, 1836.33, 1870.98, 1905.62, 1940.27, 1974.92, 2009.57, 2044.22, 2078.86, 2113.51, 2148.16, 2182.81, 2217.45, 2252.10, 2286.75, 2321.40, 2356.05, 2390.69, 2425.34, 2459.99, 2494.64, 2529.28, 2563.93, 2598.58, 2633.23, 2667.88, 2702.52, 2737.17, 2771.82,
			2806.47, 2841.11, 2875.76, 2910.41, 2945.06, 2979.71, 3014.35, 3049.00, 3083.65, 3118.30, 3152.95, 3187.59, 3222.24, 3256.89, 3291.54, 3326.18, 3360.83, 3395.48, 3430.13, 3464.78, 3499.42, 3534.07, 3568.72, 3603.37, 3638.01, 3672.66, 3707.31, 3741.96, 3776.61, 3811.25, 3845.90, 3880.55, 3915.20, 3949.84, 3984.49, 4019.14, 4053.79, 4088.44, 4123.08, 4157.73,
			4192.38, 4227.03, 4261.67, 4296.32, 4330.97, 4365.62, 4400.27, 4434.91, 4469.56, 4504.21, 4538.86, 4573.51, 4608.15, 4642.80, 4677.45, 4712.10, 4746.74, 4781.39, 4816.04, 4850.69, 4885.34, 4919.98, 4954.63, 4989.28, 5023.93, 5058.57, 5093.22, 5127.87, 5162.52, 5197.17, 5231.81, 5266.46, 5301.11, 5335.76, 5370.40, 5405.05, 5439.70, 5474.35, 5509.00, 5543.64,
			5578.29, 5612.94, 5647.59, 5682.23, 5716.88, 5751.53, 5786.18, 5820.83, 5855.47, 5890.12, 5924.77, 5959.42, 5994.06, 6028.71, 6063.36, 6098.01, 6132.66, 6167.30, 6201.95, 6236.60, 6271.25, 6305.90, 6340.54, 6375.19, 6409.84, 6444.49, 6479.13, 6513.78, 6548.43, 6583.08, 6617.73, 6652.37, 6687.02, 6721.67, 6756.32, 6790.96, 6825.61, 6860.26, 6894.91, 6929.56,
			6964.20, 6998.85, 7033.50, 7068.15, 7102.79, 7137.44, 7172.09, 7206.74, 7241.39, 7276.03, 7310.68, 7345.33, 7379.98, 7414.62, 7449.27, 7483.92, 7518.57, 7553.22, 7587.86, 7622.51, 7657.16, 7691.81, 7726.45, 7761.10, 7795.75, 7830.40, 7865.05, 7899.69, 7934.34, 7968.99, 8003.64, 8038.28, 8072.93, 8107.58, 8142.23, 8176.88, 8211.52, 8246.17, 8280.82, 8315.47,
			8350.12, 8384.76, 8419.41, 8454.06, 8488.71, 8523.35, 8558.00, 8592.65, 8627.30, 8661.95, 8696.59, 8731.24, 8765.89, 8800.54, 8835.18, 8869.83, 8904.48, 8939.13, 8973.78, 9008.42, 9043.07, 9077.72, 9112.37, 9147.02, 9181.66, 9216.31, 9250.96, 9285.61, 9320.25, 9354.90, 9389.55, 9424.20, 9458.85, 9493.49, 9528.14, 9562.79, 9597.44, 9632.08, 9666.73, 9701.38,
			9736.03, 9770.68, 9805.32, 9839.97, 9874.62, 9909.27, 9943.91, 9978.56, 10013.2, 10047.8, 10082.5, 10117.1, 10151.8, 10186.4, 10221.1, 10255.7, 10290.3, 10325.0, 10359.6, 10394.3, 10428.9, 10463.6, 10498.2, 10532.9, 10567.5, 10602.2, 10636.8, 10671.5, 10706.1, 10740.8, 10775.4, 10810.1, 10844.7, 10879.4, 10914.0, 10948.7, 10983.3, 11018.0, 11052.6, 11087.2,
			11121.9, 11156.5, 11191.2, 11225.8, 11260.5, 11295.1, 11329.8, 11364.4, 11399.1, 11433.7, 11468.4, 11503.0, 11537.7, 11572.3, 11607.0, 11641.6, 11676.3, 11710.9, 11745.6, 11780.2, 11814.9, 11849.5, 11884.1, 11918.8, 11953.4, 11988.1, 12022.7, 12057.4, 12092.0, 12126.7, 12161.3, 12196.0, 12230.6, 12265.3, 12299.9, 12334.6, 12369.2, 12403.9, 12438.5, 12473.2,
			12507.8, 12542.5, 12577.1, 12611.8, 12646.4, 12681.0, 12715.7, 12750.3, 12785.0, 12819.6, 12854.3, 12888.9, 12923.6, 12958.2, 12992.9, 13027.5, 13062.2, 13096.8, 13131.5, 13166.1, 13200.8, 13235.4, 13270.1, 13304.7, 13339.4, 13374.0, 13408.6, 13443.3, 13477.9, 13512.6, 13547.2, 13581.9, 13616.5, 13651.2, 13685.8, 13720.5, 13755.1, 13789.8, 13824.4, 13859.1,
			13893.7, 13928.4, 13963.0, 13997.7, 14032.3, 14067.0, 14101.6, 14136.3, 14170.9, 14205.5, 14240.2, 14274.8, 14309.5, 14344.1, 14378.8, 14413.4, 14448.1, 14482.7, 14517.4, 14552.0, 14586.7, 14621.3, 14656.0, 14690.6, 14725.3, 14759.9, 14794.6, 14829.2, 14863.9, 14898.5, 14933.2, 14967.8, 15002.4, 15037.1, 15071.7, 15106.4, 15141.0, 15175.7, 15210.3, 15245.0,
			15279.6, 15314.3, 15348.9, 15383.6, 15418.2, 15452.9, 15487.5, 15522.2, 15556.8, 15591.5, 15626.1, 15660.8, 15695.4, 15730.1, 15764.7, 15799.3, 15834.0, 15868.6, 15903.3, 15937.9, 15972.6, 16007.2, 16041.9, 16076.5, 16111.2, 16145.8, 16180.5, 16215.1, 16249.8, 16284.4, 16319.1, 16353.7, 16388.4, 16423.0, 16457.7, 16492.3, 16527.0, 16561.6, 16596.2, 16630.9,
			16665.5, 16700.2, 16734.8, 16769.5, 16804.1, 16838.8, 16873.4, 16908.1, 16942.7, 16977.4, 17012.0, 17046.7, 17081.3, 17116.0, 17150.6, 17185.3, 17219.9, 17254.6, 17289.2, 17323.9, 17358.5, 17393.1, 17427.8, 17462.4, 17497.1, 17531.7, 17566.4, 17601.0, 17635.7, 17670.3, 17705.0, 17739.6, 17774.3, 17808.9, 17843.6, 17878.2, 17912.9, 17947.5, 17982.2, 18016.8,
			18051.5, 18086.1, 18120.7, 18155.4, 18190.0, 18224.7, 18259.3, 18294.0, 18328.6, 18363.3, 18397.9, 18432.6, 18467.2, 18501.9, 18536.5, 18571.2, 18605.8, 18640.5, 18675.1, 18709.8, 18744.4, 18779.1, 18813.7, 18848.4, 18883.0, 18917.7, 18952.3, 18986.9, 19021.6, 19056.2, 19090.9, 19125.5, 19160.2, 19194.8, 19229.5, 19264.1, 19298.8, 19333.4, 19368.1, 19402.7,
			19437.4, 19472.0, 19506.7, 19541.3, 19576.0, 19610.6, 19645.3, 19679.9, 19714.5, 19749.2, 19783.8, 19818.5, 19853.1, 19887.8, 19922.4, 19957.1, 19991.7, 20026.4, 20061.0, 20095.7, 20130.3, 20165.0, 20199.6, 20234.3, 20268.9, 20303.6, 20338.2, 20372.9, 20407.5, 20442.2, 20476.8, 20511.4, 20546.1, 20580.7, 20615.4, 20650.0, 20684.7, 20719.3, 20754.0, 20788.6,
			20823.3, 20857.9, 20892.6, 20927.2, 20961.9, 20996.5, 21031.2, 21065.8, 21100.5, 21135.1, 21169.8, 21204.4, 21239.1, 21273.7, 21308.3, 21343.0, 21377.6, 21412.3, 21446.9, 21481.6, 21516.2, 21550.9, 21585.5, 21620.2, 21654.8, 21689.5, 21724.1, 21758.8, 21793.4, 21828.1, 21862.7, 21897.4, 21932.0, 21966.7, 22001.3, 22036.0, 22070.6, 22105.2, 22139.9, 22174.5,
			22209.2, 22243.8, 22278.5, 22313.1, 22347.8, 22382.4, 22417.1, 22451.7, 22486.4, 22521.0, 22555.7, 22590.3, 22625.0, 22659.6, 22694.3, 22728.9, 22763.6, 22798.2, 22832.9, 22867.5, 22902.1, 22936.8, 22971.4, 23006.1, 23040.7, 23075.4, 23110.0, 23144.7, 23179.3, 23214.0, 23248.6, 23283.3, 23317.9, 23352.6, 23387.2, 23421.9, 23456.5, 23491.2, 23525.8, 23560.5,
			23595.1, 23629.8, 23664.4, 23699.0, 23733.7, 23768.3, 23803.0, 23837.6, 23872.3, 23906.9, 23941.6, 23976.2, 24010.9, 24045.5, 24080.2, 24114.8, 24149.5, 24184.1, 24218.8, 24253.4, 24288.1, 24322.7, 24357.4, 24392.0, 24426.7, 24461.3, 24495.9, 24530.6, 24565.2, 24599.9, 24634.5, 24669.2, 24703.8, 24738.5, 24773.1, 24807.8, 24842.4, 24877.1, 24911.7, 24946.4,
			24981.0, 25015.7, 25050.3, 25085.0, 25119.6, 25154.3, 25188.9, 25223.6, 25258.2, 25292.8, 25327.5, 25362.1, 25396.8, 25431.4, 25466.1, 25500.7, 25535.4, 25570.0, 25604.7, 25639.3, 25674.0, 25708.6, 25743.3, 25777.9, 25812.6, 25847.2, 25881.9, 25916.5, 25951.2, 25985.8, 26020.4, 26055.1, 26089.7, 26124.4, 26159.0, 26193.7, 26228.3, 26263.0, 26297.6, 26332.3,
			26366.9, 26401.6, 26436.2, 26470.9, 26505.5, 26540.2, 26574.8, 26609.5, 26644.1, 26678.8, 26713.4, 26748.1, 26782.7, 26817.3, 26852.0, 26886.6, 26921.3, 26955.9, 26990.6, 27025.2, 27059.9, 27094.5, 27129.2, 27163.8, 27198.5, 27233.1, 27267.8, 27302.4, 27337.1, 27371.7, 27406.4, 27441.0, 27475.7, 27510.3, 27545.0, 27579.6, 27614.2, 27648.9, 27683.5, 27718.2,
			27752.8, 27787.5, 27822.1, 27856.8, 27891.4, 27926.1, 27960.7, 27995.4, 28030.0, 28064.7, 28099.3, 28134.0, 28168.6, 28203.3, 28237.9, 28272.6, 28307.2, 28341.9, 28376.5, 28411.1, 28445.8, 28480.4, 28515.1, 28549.7, 28584.4, 28619.0, 28653.7, 28688.3, 28723.0, 28757.6, 28792.3, 28826.9, 28861.6, 28896.2, 28930.9, 28965.5, 29000.2, 29034.8, 29069.5, 29104.1,
			29138.8, 29173.4, 29208.0, 29242.7, 29277.3, 29312.0, 29346.6, 29381.3, 29415.9, 29450.6, 29485.2, 29519.9, 29554.5, 29589.2, 29623.8, 29658.5, 29693.1, 29727.8, 29762.4, 29797.1, 29831.7, 29866.4, 29901.0, 29935.7, 29970.3, 30004.9, 30039.6, 30074.2, 30108.9, 30143.5, 30178.2, 30212.8, 30247.5, 30282.1, 30316.8, 30351.4, 30386.1, 30420.7, 30455.4, 30490.0,
			30524.7, 30559.3, 30594.0, 30628.6, 30663.3, 30697.9, 30732.6, 30767.2, 30801.8, 30836.5, 30871.1, 30905.8, 30940.4, 30975.1, 31009.7, 31044.4, 31079.0, 31113.7, 31148.3, 31183.0, 31217.6, 31252.3, 31286.9, 31321.6, 31356.2, 31390.9, 31425.5, 31460.2, 31494.8, 31529.4, 31564.1, 31598.7, 31633.4, 31668.0, 31702.7, 31737.3, 31772.0, 31806.6, 31841.3, 31875.9,
			31910.6, 31945.2, 31979.9, 32014.5, 32049.2, 32083.8, 32118.5, 32153.1, 32187.8, 32222.4, 32257.1, 32291.7, 32326.3, 32361.0, 32395.6, 32430.3, 32464.9, 32499.6, 32534.2, 32568.9, 32603.5, 32638.2, 32672.8, 32707.5, 32742.1, 32776.8, 32811.4, 32846.1, 32880.7, 32915.4, 32950.0, 32984.7, 33019.3, 33054.0, 33088.6, 33123.2, 33157.9, 33192.5, 33227.2, 33261.8,
			33296.5, 33331.1, 33365.8, 33400.4, 33435.1, 33469.7, 33504.4, 33539.0, 33573.7, 33608.3, 33643.0, 33677.6, 33712.3, 33746.9, 33781.6, 33816.2, 33850.9, 33885.5, 33920.1, 33954.8, 33989.4, 34024.1, 34058.7, 34093.4, 34128.0, 34162.7, 34197.3, 34232.0, 34266.6, 34301.3, 34335.9, 34370.6, 34405.2, 34439.9, 34474.5, 34509.2, 34543.8, 34578.5, 34613.1, 34647.8,
		}
		wantsPos := wantsNeg[2:]

		// Test a tone with a negative fundamental frequency.
		negTone := NewToneAt(ctx, -fundFreq)
		for i, want := range wantsNeg {
			have := negTone.HarmonicFreq(i + 1)
			require.Equal(t, want, have)
		}

		// Test a tone with a positive fundamental frequency.
		posTone := NewToneAt(ctx, fundFreq)
		for i, want := range wantsPos {
			have := posTone.HarmonicFreq(i + 1)
			require.Equal(t, want, have)
		}
	})
}

// Test_Tone_Clone tests that Tone's Clone method creates a deep copy of the tone that has all of
// the same values as the original but does not share any memory with it.
func Test_Tone_Clone(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var tone *Tone
		require.Empty(t, tone.Clone())
	})

	ctx := context.NewContext()

	type subtest struct {
		tone Tone
		name string
	}

	subtests := []subtest{
		{Tone{}, "empty"},
		{NewTone(ctx), "new"},
		{NewToneAt(ctx, 42), "frequency only"},
		{Tone{0, 0, []float32{}}, "empty, no harmonics"},
		{Tone{1, 1, []float32{}}, "frequency and gain only"},
		{Tone{0, 0, []float32{.41, 103.3}}, "harmonics only"},
		{Tone{1.1, 2.2, []float32{0, 1.1, 0.03}}, "all fields"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			clone := subtest.tone.Clone()

			// Make sure all fields have the same values.
			require.Equal(t, subtest.tone.Frequency, clone.Frequency)
			require.Equal(t, subtest.tone.Gain, clone.Gain)
			require.Equal(t, len(subtest.tone.HarmonicGains), len(clone.HarmonicGains))
			for i := range subtest.tone.HarmonicGains {
				require.Equal(t, subtest.tone.HarmonicGains[i], clone.HarmonicGains[i])
			}

			// Make sure no fields share memory.
			subtest.tone.Frequency++
			require.NotEqual(t, subtest.tone.Frequency, clone.Frequency)
			subtest.tone.Gain++
			require.NotEqual(t, subtest.tone.Gain, clone.Gain)
			for i := range subtest.tone.HarmonicGains {
				subtest.tone.HarmonicGains[i]++
				require.NotEqual(t, subtest.tone.HarmonicGains[i], clone.HarmonicGains[i])
			}
		})
	}
}

// Test_Tone_Empty tests that Tone's Empty method correctly determines whether a tone is empty.
func Test_Tone_Empty(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var tone *Tone
		require.True(t, tone.Empty())
	})

	ctx := context.NewContext()

	type subtest struct {
		want bool
		tone Tone
		name string
	}

	subtests := []subtest{
		{true, Tone{}, "empty"},
		{true, NewTone(ctx), "new"},
		{false, Tone{10, 0, nil}, "frequency only"},
		{false, Tone{0, 10, nil}, "gain only"},
		{false, Tone{0, 0, []float32{0.1, 0.2}}, "harmonics only"},
		{false, Tone{10, 10, nil}, "frequency and gain"},
		{false, Tone{10, 10, []float32{0.1, 0.2}}, "all fields"},
		{false, Tone{10, 10, []float32{-20}}, "all fields, negative harmonic gain"},
		{false, Tone{-20, 0, nil}, "negative frequency"},
		{false, Tone{0, -20, nil}, "negative gain"},
		{false, NewSquareTone(ctx, 10), "square tone"},
		{false, NewTriangleTone(ctx, 20), "triangle tone"},
		{false, NewSawtoothTone(ctx, 30), "sawtooth tone"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.tone.Empty())
		})
	}
}

// Test_Tone_Reset tests that Tone's Reset method correctly resets a tone to its default values.
func Test_Tone_Reset(t *testing.T) {
	defer SetNumHarmGains(DefaultNumHarmGains)

	t.Run("nil", func(t *testing.T) {
		var tone *Tone

		require.NotPanics(t, func() { tone.Reset() })
	})

	t.Run("uninitialized", func(t *testing.T) {
		var tone Tone

		tone.Reset()
		require.True(t, tone.Empty())
		require.Len(t, tone.HarmonicGains, 0)
	})

	for _, numHarmGains := range []int{0, 1, DefaultNumHarmGains, 100} {
		SetNumHarmGains(numHarmGains)

		ctx := context.NewContext()

		t.Run("initialized", func(t *testing.T) {
			tone := NewTone(ctx)

			tone.Reset()
			require.True(t, tone.Empty())
			require.Len(t, tone.HarmonicGains, numHarmGains)
		})

		t.Run("initialized with frequency", func(t *testing.T) {
			tone := NewToneAt(ctx, 123.11)

			tone.Reset()
			require.True(t, tone.Empty())
			require.Len(t, tone.HarmonicGains, numHarmGains)
		})

		t.Run("initialized with harmonics", func(t *testing.T) {
			tone := NewToneAt(ctx, 52.24)
			for i := range tone.HarmonicGains {
				tone.HarmonicGains[i] = float32(i) + 1.1
			}

			tone.Reset()
			require.True(t, tone.Empty())
			require.Len(t, tone.HarmonicGains, numHarmGains)
		})
	}
}
