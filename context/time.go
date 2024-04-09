package context

// Time is a timestamp of a single sample of audio data.
type Time struct {
	// Second is the number of complete seconds that have elapsed so far.
	//
	// Example: If 3 seconds have elapsed and the generator is playing the 4th second of audio data,
	// then this will be 3.
	Second int
	// Sample is the sample number in the current second. This field is 1-based, i.e. the first
	// sample is 1, the second sample is 2, etc. This value can never exceed the global sample rate.
	//
	// Example: If the generator is playing the 14th sample of the 4th second, then Second will be 3
	// and Sample will be 14.
	Sample int
}

// NewTime returns a timestamp with the lowest value.
func NewTime() Time {
	return Time{
		Second: 0,
		Sample: 1,
	}
}
