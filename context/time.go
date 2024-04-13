package context

import (
	"fmt"
	"time"
)

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

// ShiftBy shifts the timestamp by the number of samples and returns the new timestamp. This does
// not modify the receiver. The number of samples can be positive or negative.
func (t Time) ShiftBy(samples int) Time {
	t.Second += samples / SampleRate()
	t.Sample += samples % SampleRate()

	switch {
	case t.Sample > SampleRate():
		// We're overflowing the second. Bump up to the next second.
		t.Second++
		t.Sample -= SampleRate()

	case t.Sample <= 0:
		// We're underflowing the second. Bump down to the previous second.
		t.Second--
		t.Sample += SampleRate()
	}

	// If we went below 0 seconds, then reset the timestamp.
	if t.Second < 0 {
		t = NewTime()
	}

	return t
}

// Increment increments the timestamp by one sample and returns the new timestamp.
// This does not modify the receiver. This is an alias for t.ShiftBy(1).
func (t Time) Increment() Time {
	return t.ShiftBy(1)
}

// Decrement decrements the timestamp by one sample and returns the new timestamp. This does not
// modify the receiver. The timestamp can never go below 0. This is an alias for t.ShiftBy(-1).
func (t Time) Decrement() Time {
	return t.ShiftBy(-1)
}

// Duration calculates the duration between two timestamps to the nearest microsecond.
func (t Time) Duration(t2 Time) time.Duration {
	secondsDiff := t.Second - t2.Second
	secondsDur := time.Duration(secondsDiff) * time.Second

	samplesDiff := t.Sample - t2.Sample
	samplesConv := float64(samplesDiff) / float64(SampleRate())
	samplesDur := time.Duration(samplesConv * float64(time.Second))

	diff := (secondsDur + samplesDur).Abs()
	rounded := diff.Round(time.Microsecond)

	return rounded
}

// Equal returns true if the two timestamps are equal.
func (t Time) Equal(t2 Time) bool {
	return t == t2
}

// Before returns true if t represents a time that is earlier/lower than t2 does.
func (t Time) Before(t2 Time) bool {
	if t.Second != t2.Second {
		return t.Second < t2.Second
	}

	return t.Sample < t2.Sample
}

// After returns true if t represents a time that is later/higher than t2 does.
func (t Time) After(t2 Time) bool {
	if t.Second != t2.Second {
		return t.Second > t2.Second
	}

	return t.Sample > t2.Sample
}

// String returns the human-readable representation of t.
func (t Time) String() string {
	suffix := "s"
	if t.Second == 1 {
		suffix = ""
	}
	return fmt.Sprintf("%v second%s, sample %v/%v", t.Second, suffix, t.Sample, SampleRate())
}
