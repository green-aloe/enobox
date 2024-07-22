package tone

import (
	"sync"

	"github.com/green-aloe/enobox/context"
)

const (
	// DefaultNumHarmGains is the default number of harmonic gains above the fundamental frequency
	// that are tracked for each tone.
	DefaultNumHarmGains = 20
)

var (
	// Number of harmonic gains above the fundamental frequency
	numHarmGains      = DefaultNumHarmGains
	numHarmGainsMutex sync.RWMutex
	numHarmGainsKey   ctxKey
)

type ctxKey struct{}

func init() {
	context.AddDecorator(func(ctx context.Context) context.Context {
		return ctx.WithValue(numHarmGainsKey, numHarmGains)
	})
}
