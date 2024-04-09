package context

import (
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
