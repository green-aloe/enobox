//go:build !darwin

package buffer

import (
	"github.com/green-aloe/enobox/context"
	"github.com/green-aloe/enobox/tone"
)

type Buffer struct {
	Tones []tone.Tone
}

func NewBuffer(ctx context.Context) Buffer {
	tones := make([]tone.Tone, ctx.SampleRate())
	for i := range tones {
		tones[i] = tone.NewTone(ctx)
	}

	return Buffer{
		Tones: tones,
	}
}

// Release releases the buffer back to the system.
// After this, the buffer should not be used again.
func (buffer *Buffer) Release() {
}

// Reset resets a buffer to its zero values.
func (buffer *Buffer) Reset() {
}
