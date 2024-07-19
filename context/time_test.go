package context

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_NewTime tests that NewTime returns a Time object with the correct values.
func Test_NewTime(t *testing.T) {
	t.Run("types", func(t *testing.T) {
		time := NewTime()
		require.IsType(t, Time{}, time)
		require.IsType(t, int(0), time.second)
		require.IsType(t, int(0), time.sample)
		require.IsType(t, int(0), time.sampleRate)
	})

	t.Run("values", func(t *testing.T) {
		time := NewTime()
		require.Equal(t, 0, time.second)
		require.Equal(t, 1, time.sample)
		require.Equal(t, 44_100, time.sampleRate)
	})
}

// Test_NewTimeWith tests that NewTimeWith returns a Time object with the correct values.
func Test_NewTimeWith(t *testing.T) {
	t.Run("types", func(t *testing.T) {
		time := NewTimeWith(100)
		require.IsType(t, Time{}, time)
	})

	t.Run("invalid", func(t *testing.T) {
		for _, sampleRate := range []int{-100, -1, 0} {
			require.PanicsWithValue(t, "invalid time", func() {
				NewTimeWith(sampleRate)
			})
		}
	})

	t.Run("values", func(t *testing.T) {
		for _, sampleRate := range []int{1, 100, 44_100, 48_000} {
			time := NewTimeWith(sampleRate)
			require.Equal(t, 0, time.second)
			require.Equal(t, 1, time.sample)
			require.Equal(t, sampleRate, time.sampleRate)
		}
	})
}

// Test_Time_Second tests that Time's Second method returns the correct number of complete seconds
// that have elapsed so far.
func Test_Time_Second(t *testing.T) {
	type subtest struct {
		time Time
		want int
		name string
	}

	subtests := []subtest{
		{NewTime(), 0, "empty"},
		{Time{100, 0, 0}, 100, "seconds without samples, no rate"},
		{Time{0, 100, 0}, 0, "samples without seconds, no rate"},
		{Time{50, 100, 0}, 50, "seconds and samples, no rate"},
		{Time{100, 0, 44_100}, 100, "seconds without samples, has rate"},
		{Time{0, 100, 44_100}, 0, "samples without seconds, has rate"},
		{Time{50, 100, 44_100}, 50, "seconds and samples, has rate"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.time.second)
			require.Equal(t, subtest.want, subtest.time.Second())
		})
	}
}

// Test_Time_Sample tests that Time's Sample method returns the correct sample number in the current
// second.
func Test_Time_Sample(t *testing.T) {
	type subtest struct {
		time Time
		want int
		name string
	}

	subtests := []subtest{
		{NewTime(), 1, "empty"},
		{Time{1, 100, 0}, 100, "samples without seconds, no rate"},
		{Time{100, 1, 0}, 1, "seconds without samples, no rate"},
		{Time{50, 70, 0}, 70, "samples and seconds, no rate"},
		{Time{1, 100, 44_100}, 100, "samples without seconds, has rate"},
		{Time{100, 1, 44_100}, 1, "seconds without samples, has rate"},
		{Time{50, 70, 44_100}, 70, "samples and seconds, has rate"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.time.sample)
			require.Equal(t, subtest.want, subtest.time.Sample())
		})
	}
}

// Test_Time_SampleRate tests that Time's SampleRate method returns the correct sample rate no
// matter what the underlying time is.
func Test_Time_SampleRate(t *testing.T) {
	type subtest struct {
		time Time
		want int
		name string
	}

	subtests := []subtest{
		{NewTime(), 44_100, "default"},
		{NewTimeWith(0), 0, "empty"},
		{Time{1, 100, 0}, 0, "samples without seconds, no rate"},
		{Time{100, 1, 0}, 0, "seconds without samples, no rate"},
		{Time{50, 70, 0}, 0, "samples and seconds, no rate"},
		{Time{1, 100, 44_100}, 44_100, "samples without seconds, has rate"},
		{Time{100, 1, 44_100}, 44_100, "seconds without samples, has rate"},
		{Time{50, 70, 44_100}, 44_100, "samples and seconds, has rate"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.time.sampleRate)
			require.Equal(t, subtest.want, subtest.time.SampleRate())
		})
	}
}

// Test_Time_ShiftBy tests that Time's ShiftBy method shifts the timestamp by the correct number of
// samples.
func Test_Time_ShiftBy(t *testing.T) {
	type subtest struct {
		initial    Time
		numSamples int
		want       Time
		name       string
	}

	subtests := []subtest{
		// shift down
		{Time{0, 10, 44_100}, -5, Time{0, 5, 44_100}, "shift down"},
		{Time{1, 10, 44_100}, -9, Time{1, 1, 44_100}, "shift down to boundary"},
		{Time{10, 10, 44_100}, -10, Time{9, 44_100, 44_100}, "underflow without remainder"},
		{Time{10, 10, 44_100}, -11, Time{9, 44_099, 44_100}, "underflow with remainder"},
		{Time{0, 10, 100}, -5, Time{0, 5, 100}, "shift down with low sample rate"},
		{Time{1, 10, 100}, -9, Time{1, 1, 100}, "shift down to boundary with low sample rate"},
		{Time{10, 10, 100}, -10, Time{9, 100, 100}, "underflow without remainder with low sample rate"},
		{Time{10, 10, 100}, -11, Time{9, 99, 100}, "underflow with remainder with low sample rate"},
		{Time{10, 10, 44_100}, -44_100, Time{9, 10, 44_100}, "shift down full second off boundary"},
		{Time{10, 1, 44_100}, -44_100, Time{9, 1, 44_100}, "shift down full second on boundary"},
		{Time{10, 10, 44_100}, -44_100*3 - 10_000, Time{6, 34_110, 44_100}, "shift down multiple seconds"},

		// no shift
		{Time{10, 10, 44_100}, 0, Time{10, 10, 44_100}, "no shift"},

		// shift up
		{Time{10, 10, 44_100}, 10, Time{10, 20, 44_100}, "shift up"},
		{Time{10, 44_095, 44_100}, 5, Time{10, 44_100, 44_100}, "shift up to boundary"},
		{Time{10, 44_095, 44_100}, 6, Time{11, 1, 44_100}, "overflow without remainder"},
		{Time{10, 44_095, 44_100}, 10, Time{11, 5, 44_100}, "overflow with remainder"},
		{Time{10, 10, 100}, 10, Time{10, 20, 100}, "shift up with low sample rate"},
		{Time{10, 95, 100}, 5, Time{10, 100, 100}, "shift up to boundary with low sample rate"},
		{Time{10, 95, 100}, 6, Time{11, 1, 100}, "overflow without remainder with low sample rate"},
		{Time{10, 95, 100}, 10, Time{11, 5, 100}, "overflow with remainder with low sample rate"},
		{Time{10, 10, 44_100}, 44_100, Time{11, 10, 44_100}, "shift up full second off boundary"},
		{Time{10, 1, 44_100}, 44_100, Time{11, 1, 44_100}, "shift up full second on boundary"},
		{Time{10, 10, 44_100}, 44_100*3 + 10_000, Time{13, 10_010, 44_100}, "shift up multiple seconds"},

		// edge cases
		{Time{0, 10, 44_100}, -100, NewTime(), "shift below zero seconds"},
		{Time{math.MaxInt, 44_100, 44_100}, 1, Time{0, 1, 44_100}, "increment max time"},
		{NewTime(), math.MaxInt, Time{209146758205323, 31508, 44_100}, "increment by max amount"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			have := subtest.initial.ShiftBy(subtest.numSamples)
			require.Equal(t, subtest.want, have)
		})
	}

	t.Run("immutable", func(t *testing.T) {
		time1 := NewTime()
		time2 := time1.ShiftBy(5)
		require.NotEqual(t, time1, time2)
		require.Equal(t, 1, time1.Sample())
		require.Equal(t, 6, time2.Sample())
	})
}

// Test_Time_Increment tests that Time's Increment method always increments the timestamp by one
// sample and never modifies the receiver.
func Test_Time_Increment(t *testing.T) {

	t.Run("increment", func(t *testing.T) {
		time := NewTime()
		require.Equal(t, 1, time.Sample())
		time = time.Increment()
		require.Equal(t, 2, time.Sample())
	})

	t.Run("seconds boundary", func(t *testing.T) {
		time := NewTime().ShiftBy(44_100 - 1)
		require.Equal(t, 0, time.Second())
		require.Equal(t, 44_100, time.Sample())

		time = time.Increment()
		require.Equal(t, 1, time.Second())
		require.Equal(t, 1, time.Sample())
	})

	t.Run("immutable", func(t *testing.T) {
		time1 := NewTime()
		time2 := time1.Increment()
		require.NotEqual(t, time1, time2)
		require.Equal(t, 1, time1.Sample())
		require.Equal(t, 2, time2.Sample())
	})
}

// Test_Time_Decrement tests that Time's Decrement method always decrements the timestamp by one
// sample and never modifies the receiver.
func Test_Time_Decrement(t *testing.T) {

	t.Run("increment", func(t *testing.T) {
		time := NewTime().ShiftBy(5)
		require.Equal(t, 6, time.Sample())
		time = time.Decrement()
		require.Equal(t, 5, time.Sample())
	})

	t.Run("seconds boundary", func(t *testing.T) {
		time := NewTime().ShiftBy(44_100)
		require.Equal(t, 1, time.Second())
		require.Equal(t, 1, time.Sample())

		time = time.Decrement()
		require.Equal(t, 0, time.Second())
		require.Equal(t, 44_100, time.Sample())
	})

	t.Run("minimum", func(t *testing.T) {
		time := NewTime()
		require.Equal(t, 0, time.Second())
		require.Equal(t, 1, time.Sample())

		time = time.Decrement()
		require.Equal(t, 0, time.Second())
		require.Equal(t, 1, time.Sample())
	})

	t.Run("immutable", func(t *testing.T) {
		time1 := NewTime().ShiftBy(5)
		time2 := time1.Decrement()
		require.NotEqual(t, time1, time2)
		require.Equal(t, 6, time1.Sample())
		require.Equal(t, 5, time2.Sample())
	})
}

// Test_Time_Duration tests that Time's Duration method calculates the correct duration between two
// timestamps.
func Test_Time_Duration(t *testing.T) {
	type subtest struct {
		t1   Time
		t2   Time
		want string
		name string
	}

	subtests := []subtest{
		{NewTime(), NewTime(), "0s", "empty to empty"},
		{Time{0, 1, 44_100}, Time{0, 2, 44_100}, "23µs", "one sample"},
		{Time{0, 1, 44_100}, Time{0, 3, 44_100}, "45µs", "two samples"},
		{Time{0, 1, 44_100}, Time{0, 11, 44_100}, "227µs", "ten samples"},
		{Time{0, 44_100, 44_100}, Time{1, 1, 44_100}, "23µs", "seconds boundary"},
		{Time{0, 44_100/2 + 1, 44_100}, NewTime(), "500ms", "half second"},
		{Time{1, 1, 44_100}, NewTime(), "1s", "one second"},
		{Time{2, 1, 44_100}, NewTime(), "2s", "two seconds"},
		{Time{100, 1, 44_100}, Time{50, 1, 44_100}, "50s", "fifty seconds"},
		{Time{1, 1, 44_100}, Time{0, 10_000, 44_100}, "773.265ms", "10k samples"},
		{Time{22, 44_100, 44_100}, Time{23, 44_100, 44_100}, "1s", "one second boundary"},
		{NewTime(), Time{10_000, 1, 44_100}, "2h46m40s", "hours"},
		{NewTime(), Time{10_000, 10_000, 44_100}, "2h46m40.226735s", "hours and seconds"},
		{Time{0, 1, 100}, Time{0, 2, 100}, "10ms", "one sample, low rate"},
		{Time{1, 1, 100}, NewTimeWith(100), "1s", "one second, low rate"},
		{Time{22, 100, 100}, Time{23, 100, 100}, "1s", "one second boundary, low rate"},
		{Time{10, 10, 100}, Time{10, 10, 44_100}, "0s", "different rates"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.t1.Duration(subtest.t2).String())
			require.Equal(t, subtest.want, subtest.t2.Duration(subtest.t1).String())
		})
	}
}

// Test_Time_Equal tests that Time's Equal method correctly determines when two timestamps are equal.
func Test_Time_Equal(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		time1 := NewTime()
		time2 := NewTime()
		require.True(t, time1.Equal(time2))
		require.True(t, time2.Equal(time1))
	})

	t.Run("not empty", func(t *testing.T) {
		time1 := Time{111, 2222, 44_100}
		time2 := Time{111, 2222, 44_100}
		require.True(t, time1.Equal(time2))
		require.True(t, time2.Equal(time1))
	})

	t.Run("different seconds", func(t *testing.T) {
		time1 := Time{1, 1, 44_100}
		time2 := Time{2, 1, 44_100}
		require.False(t, time1.Equal(time2))
		require.False(t, time2.Equal(time1))
	})

	t.Run("different samples", func(t *testing.T) {
		time1 := Time{1, 10, 44_100}
		time2 := Time{1, 11, 44_100}
		require.False(t, time1.Equal(time2))
		require.False(t, time2.Equal(time1))
	})

	t.Run("different sample rates", func(t *testing.T) {
		time1 := Time{1, 10, 44_100}
		time2 := Time{1, 10, 48_000}
		require.False(t, time1.Equal(time2))
		require.False(t, time2.Equal(time1))
	})

	t.Run("all different", func(t *testing.T) {
		time1 := Time{10, 88, 100}
		time2 := Time{200, 999, 200}
		require.False(t, time1.Equal(time2))
		require.False(t, time2.Equal(time1))
	})

	t.Run("shifted", func(t *testing.T) {
		time1 := NewTime()
		time2 := NewTime()
		require.True(t, time1.Equal(time2))
		require.True(t, time2.Equal(time1))

		time1 = time1.ShiftBy(20)
		require.False(t, time1.Equal(time2))
		require.False(t, time2.Equal(time1))

		time2 = time2.ShiftBy(20)
		require.True(t, time1.Equal(time2))
		require.True(t, time2.Equal(time1))

		time1 = time1.Decrement()
		require.False(t, time1.Equal(time2))
		require.False(t, time2.Equal(time1))

		time2 = time2.Decrement()
		require.True(t, time1.Equal(time2))
		require.True(t, time2.Equal(time1))

		time1 = time1.Increment().Increment()
		require.False(t, time1.Equal(time2))
		require.False(t, time2.Equal(time1))

		time2 = time2.Increment().Increment()
		require.True(t, time1.Equal(time2))
		require.True(t, time2.Equal(time1))
	})
}

// Test_Time_Before tests that Time's Before method correctly determines when one timestamp is
// earlier than another one.
func Test_Time_Before(t *testing.T) {
	type subtest struct {
		time1 Time
		time2 Time
		want  bool
		name  string
	}

	subtests := []subtest{
		{NewTime(), NewTime(), false, "equal empty"},
		{Time{1, 1, 44_100}, Time{1, 1, 44_100}, false, "equal set"},
		{NewTime(), Time{1, 1, 44_100}, true, "empty to set"},
		{Time{1, 1, 44_100}, NewTime(), false, "set to empty"},
		{Time{1, 1, 44_100}, Time{1, 2, 44_100}, true, "lower sample"},
		{Time{1, 2, 44_100}, Time{1, 1, 44_100}, false, "higher sample"},
		{Time{1, 2, 44_100}, Time{2, 1, 44_100}, true, "lower second"},
		{Time{2, 1, 44_100}, Time{1, 2, 44_100}, false, "higher second"},
		{Time{1, 1, 44_100}, Time{2, 2, 44_100}, true, "lower second, lower sample"},
		{Time{2, 2, 44_100}, Time{1, 1, 44_100}, false, "higher second, higher sample"},
		{Time{1, 44_100, 44_100}, Time{2, 1, 44_100}, true, "seconds boundary"},
		{Time{1, 1, 100}, Time{1, 1, 100}, false, "equal set, low rate"},
		{Time{1, 1, 100}, Time{1, 2, 100}, true, "lower, low rate"},
		{Time{1, 2, 100}, Time{1, 1, 100}, false, "higher, low rate"},
		{Time{1, 1, 100}, Time{2, 2, 44_100}, false, "different rates"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.time1.Before(subtest.time2))
		})
	}
}

// Test_Time_After tests that Time's After method correctly determines when one timestamp is later
// than another one.
func Test_Time_After(t *testing.T) {
	type subtest struct {
		time1 Time
		time2 Time
		want  bool
		name  string
	}

	subtests := []subtest{
		{NewTime(), NewTime(), false, "equal empty"},
		{Time{1, 1, 44_100}, Time{1, 1, 44_100}, false, "equal set"},
		{NewTime(), Time{1, 1, 44_100}, false, "empty to set"},
		{Time{1, 1, 44_100}, NewTime(), true, "set to empty"},
		{Time{1, 1, 44_100}, Time{1, 2, 44_100}, false, "lower sample"},
		{Time{1, 2, 44_100}, Time{1, 1, 44_100}, true, "higher sample"},
		{Time{1, 2, 44_100}, Time{2, 1, 44_100}, false, "lower second"},
		{Time{2, 1, 44_100}, Time{1, 2, 44_100}, true, "higher second"},
		{Time{1, 1, 44_100}, Time{2, 2, 44_100}, false, "lower second, lower sample"},
		{Time{2, 2, 44_100}, Time{1, 1, 44_100}, true, "higher second, higher sample"},
		{Time{1, 44_100, 44_100}, Time{2, 1, 44_100}, false, "seconds boundary"},
		{Time{1, 1, 100}, Time{1, 1, 100}, false, "equal set, low rate"},
		{Time{1, 1, 100}, Time{1, 2, 100}, false, "lower, low rate"},
		{Time{1, 2, 100}, Time{1, 1, 100}, true, "higher, low rate"},
		{Time{2, 2, 44_100}, Time{1, 1, 100}, false, "different rates"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.time1.After(subtest.time2))
		})
	}
}

// Test_Time_String tests that Time's String method returns the correct string representation of the
// timestamp.
func Test_Time_String(t *testing.T) {
	type subtest struct {
		time Time
		want string
	}

	subtests := []subtest{
		{NewTime(), "0 seconds, sample 1/44100"},
		{Time{0, 2, 44_100}, "0 seconds, sample 2/44100"},
		{Time{1, 1, 44_100}, "1 second, sample 1/44100"},
		{Time{100, 100, 44_100}, "100 seconds, sample 100/44100"},
		{Time{10_000, 44_100, 44_100}, "10000 seconds, sample 44100/44100"},
		{Time{20, 55, 100}, "20 seconds, sample 55/100"},
		{Time{-1, 10, 100}, "invalid time: -1 second, sample 10/100"},
		{Time{10, -1, 100}, "invalid time: 10 seconds, sample -1/100"},
		{Time{10, 0, 100}, "invalid time: 10 seconds, sample 0/100"},
		{Time{20, 101, 100}, "invalid time: 20 seconds, sample 101/100"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.want, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.time.String())
		})
	}
}

// Test_Time_Empty tests that Time's Empty method correctly determines when a timestamp is empty.
func Test_Time_Empty(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var time Time
		require.True(t, time.Empty())
	})

	t.Run("not empty", func(t *testing.T) {
		time := NewTime()
		require.False(t, time.Empty())

		time = time.Decrement()
		require.False(t, time.Empty())

		time = NewTimeWith(44_100)
		require.False(t, time.Empty())

		time = time.ShiftBy(-44_100)
		require.False(t, time.Empty())
	})
}
