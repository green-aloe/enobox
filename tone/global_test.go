package tone

import (
	"testing"

	"github.com/green-aloe/enobox/context"
	"github.com/stretchr/testify/require"
)

// Test_Consts tests any package-level constant values.
func Test_Consts(t *testing.T) {
	require.Equal(t, 20, DefaultNumHarmGains)

	require.Equal(t, 6, MaxSigFigs)
}

// Test_init tests that a context decorator is added on package initialization with the default
// number of harmonic gains.
func Test_init(t *testing.T) {
	ctx := context.NewContext()
	require.Equal(t, 20, NumHarmGains(ctx))
}
