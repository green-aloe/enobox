package stack_test

import (
	"fmt"
	"math"

	"github.com/green-aloe/enobox/stack"
)

func ExampleStack_Push() {
	var s stack.Stack[int]
	s.Push(123)
	fmt.Println(s.Count(), s.Peek())

	// Output:
	// 1 123
}

func ExampleStack_Pop() {
	var s stack.Stack[string]
	fmt.Println(s.Pop())

	s.Push("hello")
	s.Push("world")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// Output:
	//
	// world
	// hello
	//
}

func ExampleStack_Peek() {
	var s stack.Stack[[]string]
	fmt.Println(s.Peek())

	s.Push([]string{"a", "b", "c"})
	s.Push([]string{"d", "e", "f"})
	fmt.Println(s.Peek())
	fmt.Println(s.Peek())

	// Output:
	// []
	// [d e f]
	// [d e f]
}

func ExampleStack_Empty() {
	var s stack.Stack[float64]
	fmt.Println(s.Empty())

	s.Push(math.Pi)
	fmt.Println(s.Empty())

	s.Pop()
	fmt.Println(s.Empty())

	// Output:
	// true
	// false
	// true
}

func ExampleStack_Count() {
	var s stack.Stack[byte]
	fmt.Println(s.Count())

	for _, b := range []byte{'a', 'b', 'c'} {
		s.Push(b)
		fmt.Println(s.Count())
	}

	s.Pop()
	fmt.Println(s.Count())

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 2
}

func ExampleStack_Clear() {
	var s stack.Stack[bool]
	s.Push(true)
	s.Push(false)
	fmt.Println(s.Count())

	s.Clear()
	fmt.Println(s.Count())

	// Output:
	// 2
	// 0
}
