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
