package context

import "sync"

const (
	// DefaultSampleRate is the default number of samples output per second (Hz).
	DefaultSampleRate = 44_100
)

var (
	// Number of samples output per second (Hz)
	sampleRate      = DefaultSampleRate
	sampleRateMutex sync.RWMutex
)

// SampleRate returns the global sample rate.
func SampleRate() int {
	sampleRateMutex.RLock()
	defer sampleRateMutex.RUnlock()

	if sampleRate > 0 {
		return sampleRate
	}

	return DefaultSampleRate
}

// SetSampleRate sets the global sample rate (number of samples per second (Hz) in the output). The
// rate must be greater than zero.
func SetSampleRate(rate int) {
	sampleRateMutex.Lock()
	defer sampleRateMutex.Unlock()

	if rate > 0 {
		sampleRate = rate
	}
}
