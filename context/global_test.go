package context

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_Consts tests any package-level constant values.
func Test_Consts(t *testing.T) {
	require.Equal(t, 44_100, DefaultSampleRate)
}

// Test_SampleRate tests that SampleRate returns the correct global sample rate.
func Test_SampleRate(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	t.Run("default", func(t *testing.T) {
		require.Equal(t, 44_100, SampleRate())
	})

	t.Run("custom rate", func(t *testing.T) {
		sampleRate = 20_000
		require.Equal(t, 20_000, SampleRate())
	})

	t.Run("zero rate", func(t *testing.T) {
		sampleRate = 0
		require.Equal(t, 44_100, SampleRate())
	})

	t.Run("negative rate", func(t *testing.T) {
		sampleRate = -20_000
		require.Equal(t, 44_100, SampleRate())
	})
}

// Test_SetSampleRate tests that SetSampleRate sets the global sample rate correctly.
func Test_SetSampleRate(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	t.Run("custom rate", func(t *testing.T) {
		SetSampleRate(1_000)
		require.Equal(t, 1_000, SampleRate())
	})

	t.Run("zero rate", func(t *testing.T) {
		SetSampleRate(0)
		require.Equal(t, 1_000, SampleRate())
	})

	t.Run("negative rate", func(t *testing.T) {
		SetSampleRate(-1_000)
		require.Equal(t, 1_000, SampleRate())
	})
}

// Test_SampleRate_Concurrency tests that SampleRate and SetSampleRate can be called concurrently.
func Test_SampleRate_Concurrency(t *testing.T) {
	require.Equal(t, 44_100, SampleRate())
	defer SetSampleRate(DefaultSampleRate)

	numRoutines := 100
	numCycles := 10_000

	var wgStart sync.WaitGroup
	wgStart.Add(1)
	var wgEnd sync.WaitGroup

	for i := 0; i < numRoutines; i++ {
		wgEnd.Add(1)
		go func() {
			defer wgEnd.Done()

			wgStart.Wait()
			for i := 0; i < numCycles; i++ {
				SetSampleRate(SampleRate() + 1)
			}
		}()
	}

	wgStart.Done()
	wgEnd.Wait()
}
