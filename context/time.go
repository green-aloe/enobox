package context

import (
	"fmt"
	"time"
)

// Time is a timestamp of a single sample of audio data.
type Time struct {
	// number of complete seconds that have elapsed so far.
	second int
	// sample number in the current second
	sample int
	// number of samples in one second
	sampleRate int
}

// NewTime returns a timestamp with the lowest value. It uses the global sample rate to know how
// many samples are in one second.
func NewTime() Time {
	return NewTimeWith(SampleRate())
}

// NewTimeWith returns a timestamp with the lowest value. It uses the provided sample rate to know
// how many samples are in one second.
func NewTimeWith(sampleRate int) Time {
	return NewTimeAt(0, 1, sampleRate)
}

// NewTimeAt returns a timestamp with the provided data. It panics if second is less than 0, sample
// is less than 1, or sample is greater than sampleRate.
func NewTimeAt(second, sample, sampleRate int) Time {
	if second < 0 || sample < 1 || sample > sampleRate {
		panic("invalid time")
	}

	return Time{
		second:     second,
		sample:     sample,
		sampleRate: sampleRate,
	}
}

// Second returns the number of complete seconds that have elapsed so far. For example, if three
// seconds have elapsed and the generator is on the fourth second of audio data, then Second will
// return 3.
func (t Time) Second() int {
	return t.second
}

// Sample returns the sample number in the current second. For example, if the generator is playing
// the 14th sample of the 4th second, then Sample will return 14 and Second will return 3.
//
// This field is 1-based, i.e. the first sample is 1, the second sample is 2, etc. This value can
// never exceed the global sample rate.
func (t Time) Sample() int {
	return t.sample
}

// SampleRate returns the number of samples in one second.
func (t Time) SampleRate() int {
	return t.sampleRate
}

// ShiftBy shifts the timestamp by the number of samples and returns the new timestamp. This does
// not modify the receiver. The number of samples can be positive or negative.
func (t Time) ShiftBy(samples int) Time {
	t.second += samples / t.sampleRate
	t.sample += samples % t.sampleRate

	switch {
	case t.sample > t.sampleRate:
		// We're overflowing the second. Bump up to the next second.
		t.second++
		t.sample -= t.sampleRate

	case t.sample <= 0:
		// We're underflowing the second. Bump down to the previous second.
		t.second--
		t.sample += t.sampleRate
	}

	// If we went below 0 seconds, then reset the timestamp.
	if t.second < 0 {
		t = NewTimeWith(t.sampleRate)
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

// Duration calculates the duration between two timestamps to the nearest microsecond. The returned
// duration is always positive. This returns 0 if the timestamps do not have the same sample rate.
func (t Time) Duration(t2 Time) time.Duration {
	if t.sampleRate != t2.sampleRate {
		return 0
	}

	secondsDiff := t.second - t2.second
	secondsDur := time.Duration(secondsDiff) * time.Second

	samplesDiff := t.sample - t2.sample
	samplesConv := float64(samplesDiff) / float64(t.sampleRate)
	samplesDur := time.Duration(samplesConv * float64(time.Second))

	diff := (secondsDur + samplesDur).Abs()
	rounded := diff.Round(time.Microsecond)

	return rounded
}

// Equal returns true if the two timestamps are equal.
func (t Time) Equal(t2 Time) bool {
	return t == t2
}

// Before returns true if t represents a time that is earlier/lower than t2 does. This returns false
// if the timestamps do not have the same sample rate.
func (t Time) Before(t2 Time) bool {
	if t.sampleRate != t2.sampleRate {
		return false
	}

	if t.second != t2.second {
		return t.second < t2.second
	}

	return t.sample < t2.sample
}

// After returns true if t represents a time that is later/higher than t2 does. This returns false
// if the timestamps do not have the same sample rate.
func (t Time) After(t2 Time) bool {
	if t.sampleRate != t2.sampleRate {
		return false
	}

	if t.second != t2.second {
		return t.second > t2.second
	}

	return t.sample > t2.sample
}

// String returns the human-readable representation of t.
func (t Time) String() string {

	suffix := "s"
	if t.second == 1 || t.second == -1 {
		suffix = ""
	}

	s := fmt.Sprintf("%v second%s, sample %v/%v", t.second, suffix, t.sample, t.sampleRate)
	if t.second < 0 || t.sample < 1 || t.sample > t.sampleRate {
		s = "invalid time: " + s
	}

	return s
}

// Empty checks if the timestamp is the zero value.
func (t Time) Empty() bool {
	return t == Time{}
}
