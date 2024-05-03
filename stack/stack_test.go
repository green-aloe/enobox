package stack

import (
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_Stack_Push tests that Stack's Push method adds a value to the top of the stack for various
// stack configurations.
func Test_Stack_Push(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack
		require.NotPanics(t, func() { s.Push(1) })
		require.NotPanics(t, func() { s.Push(1) })
		require.NotPanics(t, func() { s.Push(1) })
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack
		require.NotPanics(t, func() { s.Push(1) })
		require.NotPanics(t, func() { s.Push(1) })
		require.NotPanics(t, func() { s.Push(1) })
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack
		require.NotPanics(t, func() { s.Push("1") })
		require.NotPanics(t, func() { s.Push("2") })
		require.NotPanics(t, func() { s.Push("3") })
		require.Equal(t, 3, s.Size())
		require.Equal(t, "3", s.Pop())
		require.Equal(t, "2", s.Pop())
		require.Equal(t, "1", s.Pop())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack

		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				require.NotPanics(t, func() { s.Push(i) })
			}(i)
		}

		wg.Wait()
		require.Equal(t, 100, s.Size())
	})
}

// Test_Stack_Pop tests that Stack's Pop method removes and returns the value at the top of the
// stack for various stack configurations.
func Test_Stack_Pop(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack
		require.Nil(t, s.Pop())
		require.Nil(t, s.Pop())
		require.Nil(t, s.Pop())
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack
		require.Nil(t, s.Pop())
		require.Nil(t, s.Pop())
		require.Nil(t, s.Pop())
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack
		for i := 1; i <= 10; i++ {
			s.Push(i)
		}
		for i := 10; i >= 1; i-- {
			require.Equal(t, i, s.Pop())
		}
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack

		var want []int
		for i := 0; i < 100; i++ {
			s.Push(i)
			want = append(want, i)
		}

		ch := make(chan any, 100)
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Pop()
			}()
		}

		var have []int
		for i := 0; i < 100; i++ {
			v := <-ch
			i, ok := v.(int)
			require.True(t, ok)
			have = append(have, i)
		}
		sort.Slice(have, func(i, j int) bool { return have[i] < have[j] })
		require.Equal(t, want, have)
		require.Len(t, ch, 0)
	})
}

// Test_Stack_Peek tests that Stack's Peek method returns the value at the top of the stack without
// removing it for various stack configurations.
func Test_Stack_Peek(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack
		require.Nil(t, s.Peek())
		require.Nil(t, s.Peek())
		require.Nil(t, s.Peek())
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack
		require.Nil(t, s.Peek())
		require.Nil(t, s.Peek())
		require.Nil(t, s.Peek())
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack
		s.Push("a")
		s.Push("b")
		s.Push("c")
		require.Equal(t, "c", s.Peek())
		require.Equal(t, "c", s.Peek())
		require.Equal(t, "c", s.Peek())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack
		s.Push(1.1)
		s.Push(2.2)
		s.Push(3.3)

		ch := make(chan any, 100)
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Peek()
			}()
		}

		for i := 0; i < 100; i++ {
			v := <-ch
			f, ok := v.(float64)
			require.True(t, ok)
			require.Equal(t, 3.3, f)
		}
		require.Len(t, ch, 0)
	})
}

// Test_Stack_Empty tests that Stack's Empty method returns true if the stack is empty for various
// stack configurations.
func Test_Stack_Empty(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack
		require.True(t, s.Empty())
		require.True(t, s.Empty())
		require.True(t, s.Empty())
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack
		require.True(t, s.Empty())
		require.True(t, s.Empty())
		require.True(t, s.Empty())
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack
		for _, r := range "the quick brown fox jumps over the lazy dog" {
			s.Push(r)
			require.False(t, s.Empty())
		}
		require.False(t, s.Empty())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack

		ch := make(chan bool, 100)
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Empty()
			}()
		}

		for i := 0; i < 100; i++ {
			b := <-ch
			require.True(t, b)
		}
		require.Len(t, ch, 0)

		s.Push('😀')
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Empty()
			}()
		}

		for i := 0; i < 100; i++ {
			b := <-ch
			require.False(t, b)
		}
		require.Len(t, ch, 0)
	})
}

// Test_Stack_Size tests that Stack's Size method returns the number of elements in the stack for
// various stack configurations.
func Test_Stack_Size(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack
		require.Zero(t, s.Size())
		require.Zero(t, s.Size())
		require.Zero(t, s.Size())
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack
		require.Zero(t, s.Size())
		require.Zero(t, s.Size())
		require.Zero(t, s.Size())
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack
		s.Push("a")
		require.Equal(t, 1, s.Size())
		s.Push("b")
		require.Equal(t, 2, s.Size())
		s.Push("c")
		require.Equal(t, 3, s.Size())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack

		ch := make(chan int, 100)
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Size()
			}()
		}

		for i := 0; i < 100; i++ {
			i := <-ch
			require.Zero(t, i)
		}
		require.Len(t, ch, 0)

		for i := 0; i < 100; i++ {
			s.Push(i)
		}

		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Size()
			}()
		}

		for i := 0; i < 100; i++ {
			i := <-ch
			require.Equal(t, 100, i)
		}
		require.Len(t, ch, 0)
	})
}

// Test_Stack_Clear tests that Stack's Clear method removes all elements from the stack for various
// stack configurations.
func Test_Stack_Clear(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack
		require.NotPanics(t, func() { s.Clear() })
		require.True(t, s.Empty())

		require.NotPanics(t, func() { s.Clear() })
		require.NotPanics(t, func() { s.Clear() })
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack
		require.NotPanics(t, func() { s.Clear() })
		require.True(t, s.Empty())

		require.NotPanics(t, func() { s.Clear() })
		require.NotPanics(t, func() { s.Clear() })
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack
		s.Push("a")
		s.Push("b")
		s.Push("c")
		require.False(t, s.Empty())
		require.NotPanics(t, func() { s.Clear() })
		require.True(t, s.Empty())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack

		for i := 0; i < 100; i++ {
			go func() {
				s.Clear()
			}()
		}

		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.Push(1)
				s.Clear()
				s.Push(2)
				s.Clear()
				s.Push(3)
				s.Clear()
			}()
		}
		wg.Wait()
		require.True(t, s.Empty())
	})
}

// Test_differentTypes tests that Stack can handle values of different types.
func Test_differentTypes(t *testing.T) {
	var s Stack
	s.Push(1)
	s.Push("2")
	s.Push(3.0)
	require.Equal(t, 3.0, s.Pop())
	require.Equal(t, "2", s.Pop())
	require.Equal(t, 1, s.Pop())
}
