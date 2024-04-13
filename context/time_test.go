package context

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_NewTime tests that NewTime returns a Time object with the correct values.
func Test_NewTime(t *testing.T) {
	time := NewTime()
	require.IsType(t, Time{}, time)
	require.IsType(t, int(0), time.Second)
	require.IsType(t, int(0), time.Sample)
	require.Equal(t, 0, time.Second)
	require.Equal(t, 1, time.Sample)
}

// Test_ShiftBy tests that Time's ShiftBy method shifts the timestamp by the correct number of
// samples.
func Test_ShiftBy(t *testing.T) {
	type subtest struct {
		initial    Time
		numSamples int
		want       Time
		name       string
	}

	subtests := []subtest{
		// shift down
		{Time{0, 10}, -5, Time{0, 5}, "shift down"},
		{Time{1, 10}, -9, Time{1, 1}, "shift down to boundary"},
		{Time{10, 10}, -10, Time{9, 44_100}, "underflow without remainder"},
		{Time{10, 10}, -11, Time{9, 44_099}, "underflow with remainder"},
		{Time{10, 10}, -44_100, Time{9, 10}, "shift down full second off boundary"},
		{Time{10, 1}, -44_100, Time{9, 1}, "shift down full second on boundary"},
		{Time{10, 10}, -44_100*3 - 10_000, Time{6, 34_110}, "shift down multiple seconds"},

		// no shift
		{Time{10, 10}, 0, Time{10, 10}, "no shift"},

		// shift up
		{Time{10, 10}, 10, Time{10, 20}, "shift up"},
		{Time{10, 44_095}, 5, Time{10, 44_100}, "shift up to boundary"},
		{Time{10, 44_095}, 6, Time{11, 1}, "overflow without remainder"},
		{Time{10, 44_095}, 10, Time{11, 5}, "overflow with remainder"},
		{Time{10, 10}, 44_100, Time{11, 10}, "shift up full second off boundary"},
		{Time{10, 1}, 44_100, Time{11, 1}, "shift up full second on boundary"},
		{Time{10, 10}, 44_100*3 + 10_000, Time{13, 10_010}, "shift up multiple seconds"},

		// edge cases
		{Time{0, 10}, -100, NewTime(), "shift below zero seconds"},
		{Time{math.MaxInt, 44_100}, 1, Time{0, 1}, "increment max time"},
		{NewTime(), math.MaxInt, Time{209146758205323, 31508}, "increment by max amount"},
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
		require.Equal(t, 1, time1.Sample)
		require.Equal(t, 6, time2.Sample)
	})
}

// Test_Increment tests that Time's Increment method always increments the timestamp by one sample
// and never modifies the receiver.
func Test_Increment(t *testing.T) {

	t.Run("increment", func(t *testing.T) {
		time := NewTime()
		require.Equal(t, 1, time.Sample)
		time = time.Increment()
		require.Equal(t, 2, time.Sample)
	})

	t.Run("seconds boundary", func(t *testing.T) {
		time := NewTime().ShiftBy(SampleRate() - 1)
		require.Equal(t, 0, time.Second)
		require.Equal(t, SampleRate(), time.Sample)

		time = time.Increment()
		require.Equal(t, 1, time.Second)
		require.Equal(t, 1, time.Sample)
	})

	t.Run("immutable", func(t *testing.T) {
		time1 := NewTime()
		time2 := time1.Increment()
		require.NotEqual(t, time1, time2)
		require.Equal(t, 1, time1.Sample)
		require.Equal(t, 2, time2.Sample)
	})
}

// Test_Decrement tests that Time's Decrement method always decrements the timestamp by one sample
// and never modifies the receiver.
func Test_Decrement(t *testing.T) {

	t.Run("increment", func(t *testing.T) {
		time := NewTime().ShiftBy(5)
		require.Equal(t, 6, time.Sample)
		time = time.Decrement()
		require.Equal(t, 5, time.Sample)
	})

	t.Run("seconds boundary", func(t *testing.T) {
		time := NewTime().ShiftBy(SampleRate())
		require.Equal(t, 1, time.Second)
		require.Equal(t, 1, time.Sample)

		time = time.Decrement()
		require.Equal(t, 0, time.Second)
		require.Equal(t, SampleRate(), time.Sample)
	})

	t.Run("minimum", func(t *testing.T) {
		time := NewTime()
		require.Equal(t, 0, time.Second)
		require.Equal(t, 1, time.Sample)

		time = time.Decrement()
		require.Equal(t, 0, time.Second)
		require.Equal(t, 1, time.Sample)
	})

	t.Run("immutable", func(t *testing.T) {
		time1 := NewTime().ShiftBy(5)
		time2 := time1.Decrement()
		require.NotEqual(t, time1, time2)
		require.Equal(t, 6, time1.Sample)
		require.Equal(t, 5, time2.Sample)
	})
}

// Test_Duration tests that Time's Duration method calculates the correct duration between two
// timestamps.
func Test_Duration(t *testing.T) {
	type subtest struct {
		t1   Time
		t2   Time
		want string
		name string
	}

	subtests := []subtest{
		{NewTime(), NewTime(), "0s", "empty to empty"},
		{Time{0, 1}, Time{0, 2}, "23µs", "one sample"},
		{Time{0, 1}, Time{0, 3}, "45µs", "two samples"},
		{Time{0, 1}, Time{0, 11}, "227µs", "ten samples"},
		{Time{0, SampleRate()}, Time{1, 1}, "23µs", "seconds boundary"},
		{Time{0, SampleRate()/2 + 1}, NewTime(), "500ms", "half second"},
		{Time{1, 1}, NewTime(), "1s", "one second"},
		{Time{2, 1}, NewTime(), "2s", "two seconds"},
		{Time{100, 1}, Time{50, 1}, "50s", "fifty seconds"},
		{Time{1, 1}, Time{0, 10_000}, "773.265ms", "10k samples"},
		{Time{22, SampleRate()}, Time{23, SampleRate()}, "1s", "one second boundary"},
		{NewTime(), Time{10_000, 1}, "2h46m40s", "hours"},
		{NewTime(), Time{10_000, 10_000}, "2h46m40.226735s", "hours and seconds"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.t1.Duration(subtest.t2).String())
			require.Equal(t, subtest.want, subtest.t2.Duration(subtest.t1).String())
		})
	}
}

// Test_Equal tests that Time's Equal method correctly determines when two timestamps are equal.
func Test_Equal(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		time1 := NewTime()
		time2 := NewTime()
		require.True(t, time1.Equal(time2))
		require.True(t, time2.Equal(time1))
	})

	t.Run("non-empty", func(t *testing.T) {
		time1 := Time{111, 2222}
		time2 := Time{111, 2222}
		require.True(t, time1.Equal(time2))
		require.True(t, time2.Equal(time1))
	})

	t.Run("different seconds", func(t *testing.T) {
		time1 := Time{1, 1}
		time2 := Time{2, 1}
		require.False(t, time1.Equal(time2))
		require.False(t, time2.Equal(time1))
	})

	t.Run("different samples", func(t *testing.T) {
		time1 := Time{1, 10}
		time2 := Time{1, 11}
		require.False(t, time1.Equal(time2))
		require.False(t, time2.Equal(time1))
	})

	t.Run("different seconds and samples", func(t *testing.T) {
		time1 := Time{10, 88}
		time2 := Time{200, 999}
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

// Test_Before tests that Time's Before method correctly determines when one timestamp is earlier
// than another one.
func Test_Before(t *testing.T) {
	type subtest struct {
		time1 Time
		time2 Time
		want  bool
		name  string
	}

	subtests := []subtest{
		{NewTime(), NewTime(), false, "equal empty"},
		{Time{1, 1}, Time{1, 1}, false, "equal set"},
		{NewTime(), Time{1, 1}, true, "empty to set"},
		{Time{1, 1}, NewTime(), false, "set to empty"},
		{Time{1, 1}, Time{1, 2}, true, "lower sample"},
		{Time{1, 2}, Time{1, 1}, false, "higher sample"},
		{Time{1, 2}, Time{2, 1}, true, "lower second"},
		{Time{2, 1}, Time{1, 2}, false, "higher second"},
		{Time{1, 1}, Time{2, 2}, true, "lower second, lower sample"},
		{Time{2, 2}, Time{1, 1}, false, "higher second, higher sample"},
		{Time{1, SampleRate()}, Time{2, 1}, true, "seconds boundary"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.time1.Before(subtest.time2))
		})
	}
}

// Test_After tests that Time's After method correctly determines when one timestamp is later than
// another one.
func Test_After(t *testing.T) {
	type subtest struct {
		time1 Time
		time2 Time
		want  bool
		name  string
	}

	subtests := []subtest{
		{NewTime(), NewTime(), false, "equal empty"},
		{Time{1, 1}, Time{1, 1}, false, "equal set"},
		{NewTime(), Time{1, 1}, false, "empty to set"},
		{Time{1, 1}, NewTime(), true, "set to empty"},
		{Time{1, 1}, Time{1, 2}, false, "lower sample"},
		{Time{1, 2}, Time{1, 1}, true, "higher sample"},
		{Time{1, 2}, Time{2, 1}, false, "lower second"},
		{Time{2, 1}, Time{1, 2}, true, "higher second"},
		{Time{1, 1}, Time{2, 2}, false, "lower second, lower sample"},
		{Time{2, 2}, Time{1, 1}, true, "higher second, higher sample"},
		{Time{1, SampleRate()}, Time{2, 1}, false, "seconds boundary"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.time1.After(subtest.time2))
		})
	}
}

// Test_String tests that Time's String method returns the correct string representation of the
// timestamp.
func Test_String(t *testing.T) {
	type subtest struct {
		time Time
		want string
	}

	subtests := []subtest{
		{NewTime(), "0 seconds, sample 1/44100"},
		{Time{0, 2}, "0 seconds, sample 2/44100"},
		{Time{1, 1}, "1 second, sample 1/44100"},
		{Time{100, 100}, "100 seconds, sample 100/44100"},
		{Time{10_000, SampleRate()}, "10000 seconds, sample 44100/44100"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.want, func(t *testing.T) {
			require.Equal(t, subtest.want, subtest.time.String())
		})
	}
}
