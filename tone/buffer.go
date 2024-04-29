package tone

import (
	"sync"

	"github.com/green-aloe/enobox/context"
)

var (
	bufferPool = sync.Pool{
		New: func() any {
			return Buffer{
				Tones: make([]Tone, context.SampleRate()),
			}
		},
	}

	zeroTones = make([]Tone, context.SampleRate())
)

type Buffer struct {
	Tones []Tone
}

// NewBuffer returns a new buffer with all default (zero) values.
// This should be released back to the system with Release after use.
func NewBuffer() Buffer {
	// Grab a buffer from the pool and ensure that it has only default values.
	buffer := bufferPool.Get().(Buffer)
	buffer.Reset()

	return buffer
}

// Release releases the buffer back to the system.
// After this, the buffer should not be used again.
func (buffer *Buffer) Release() {
	if buffer == nil {
		return
	}

	bufferPool.Put(buffer)
}

// Reset resets a buffer to its zero values.
func (buffer *Buffer) Reset() {
	if buffer == nil {
		return
	}

	copy(buffer.Tones, zeroTones)
}
