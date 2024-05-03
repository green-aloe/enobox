package stack

import "sync"

// A stack is a first-in-last-out (FILO) data structure: the last element pushed is the first
// element popped. The zero value is an empty stack and ready to use, and stacks are safe for
// concurrent use.
type Stack struct {
	top   *node
	mutex sync.Mutex
}

// A node is a single element in a stack.
type node struct {
	v    any
	next *node
}

// Push adds a value to the top of the stack.
func (s *Stack) Push(v any) {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.top = &node{
		v:    v,
		next: s.top,
	}
}

// Pop removes and returns the value at the top of the stack.
func (s *Stack) Pop() any {
	if s == nil {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.top == nil {
		return nil
	}

	v := s.top.v
	s.top = s.top.next

	return v
}

// Peek returns the value at the top of the stack without removing it.
func (s *Stack) Peek() any {
	if s == nil {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.top == nil {
		return nil
	}

	return s.top.v
}

// Empty returns true if the stack is empty.
func (s *Stack) Empty() bool {
	if s == nil {
		return true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.top == nil
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	if s == nil {
		return 0
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	size := 0
	for n := s.top; n != nil; n = n.next {
		size++
	}

	return size
}

// Clear removes all elements from the stack.
func (s *Stack) Clear() {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.top = nil
}
