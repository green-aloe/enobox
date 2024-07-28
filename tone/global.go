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
		numHarmGainsMutex.RLock()
		defer numHarmGainsMutex.RUnlock()

		return ctx.WithValue(numHarmGainsKey, numHarmGains)
	})
}

// NumHarmGains returns the number of harmonic gains in a tone for this context, or 0 if no value is set.
func NumHarmGains(ctx context.Context) int {
	if ctx == nil {
		return 0
	}

	if v := ctx.Value(numHarmGainsKey); v != nil {
		if n, ok := v.(int); ok && n > 0 {
			return n
		}
	}

	return 0
}

// SetNumHarmGains sets the global number of harmonic gains in a tone. All contexts created after
// this is called will use the value set here. The rate cannot be a negative numbers.
func SetNumHarmGains(rate int) {
	numHarmGainsMutex.Lock()
	defer numHarmGainsMutex.Unlock()

	if rate > 0 {
		numHarmGains = rate
	}
}
