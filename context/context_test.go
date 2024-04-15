package context

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_NewContext tests that NewContext returns a Context that has been initialized correctly.
func Test_NewContext(t *testing.T) {
	ctx := NewContext()
	require.NotEmpty(t, ctx)
	require.IsType(t, Context{}, ctx)
	require.Equal(t, NewTime(), ctx.Time())
	require.Equal(t, DefaultSampleRate, ctx.SampleRate())
}

// Test_Context_Time tests that Context's Time method returns the correct timestamp.
func Test_Context_Time(t *testing.T) {

	t.Run("empty", func(t *testing.T) {
		var ctx Context
		require.Equal(t, Time{}, ctx.Time())
	})

	t.Run("initial", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, 0, ctx.Time().Second())
		require.Equal(t, 1, ctx.Time().Sample())
	})

	t.Run("updated", func(t *testing.T) {
		ctx := NewContext()
		ctx.time = ctx.time.ShiftBy(10)
		require.Equal(t, 0, ctx.Time().Second())
		require.Equal(t, 11, ctx.Time().Sample())
	})

	t.Run("immutable", func(t *testing.T) {
		ctx := NewContext()
		time := ctx.Time()
		time.second, time.sample = 10, 11
		require.Equal(t, 0, ctx.Time().Second())
		require.Equal(t, 1, ctx.Time().Sample())
		require.Equal(t, 10, time.Second())
		require.Equal(t, 11, time.Sample())
	})
}

// Test_Context_SampleRate tests that Context's SampleRate method returns the correct sample rate.
func Test_Context_SampleRate(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	t.Run("empty", func(t *testing.T) {
		var ctx Context
		require.Equal(t, 0, ctx.SampleRate())
	})

	t.Run("initial", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, DefaultSampleRate, ctx.SampleRate())
	})

	t.Run("updated", func(t *testing.T) {
		SetSampleRate(100)
		ctx := NewContext()
		require.Equal(t, 100, ctx.SampleRate())
	})
}

// Test_Context_NyqistFrequency tests that Context's NyqistFrequency method returns the correct
// frequency.
func Test_Context_NyqistFrequency(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	t.Run("empty", func(t *testing.T) {
		var ctx Context
		require.Equal(t, float32(0), ctx.NyqistFrequency())
	})

	t.Run("initial", func(t *testing.T) {
		ctx := NewContext()
		require.Equal(t, float32(22_050), ctx.NyqistFrequency())
	})

	t.Run("updated", func(t *testing.T) {
		SetSampleRate(100)
		ctx := NewContext()
		require.Equal(t, float32(50), ctx.NyqistFrequency())
	})
}
