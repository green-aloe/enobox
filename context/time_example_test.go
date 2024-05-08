package context_test

import (
	"fmt"

	"github.com/green-aloe/enobox/context"
)

func ExampleNewTime() {
	time := context.NewTime()
	fmt.Println(time)

	// Output:
	// 0 seconds, sample 1/44100
}

func ExampleNewTimeWith() {
	time := context.NewTimeWith(48_000)
	fmt.Println(time)

	// Output:
	// 0 seconds, sample 1/48000
}

func ExampleTime_Second() {
	time := context.NewTime()
	fmt.Println(time.Second())

	// Output:
	// 0
}

func ExampleTime_Sample() {
	time := context.NewTime()
	fmt.Println(time.Sample())

	// Output:
	// 1
}

func ExampleTime_ShiftBy() {
	time1 := context.NewTime()
	fmt.Println(time1)

	time2 := time1.ShiftBy(context.SampleRate() + 100)
	fmt.Println(time2)

	// The original time does not change.
	fmt.Println(time1)

	// Output:
	// 0 seconds, sample 1/44100
	// 1 second, sample 101/44100
	// 0 seconds, sample 1/44100
}

func ExampleTime_Increment() {
	time1 := context.NewTime()
	fmt.Println(time1)

	time2 := time1.Increment()
	fmt.Println(time2)

	// The original time does not change.
	fmt.Println(time1)

	// Output:
	// 0 seconds, sample 1/44100
	// 0 seconds, sample 2/44100
	// 0 seconds, sample 1/44100
}

func ExampleTime_Decrement() {
	time1 := context.NewTime()
	fmt.Println(time1)

	time2 := time1.ShiftBy(10).Decrement()
	fmt.Println(time2)

	// The original time does not change.
	fmt.Println(time1)

	// Output:
	// 0 seconds, sample 1/44100
	// 0 seconds, sample 10/44100
	// 0 seconds, sample 1/44100
}

func ExampleTime_Duration() {
	time1 := context.NewTime()
	time2 := context.NewTime().ShiftBy(context.SampleRate())

	fmt.Println(time1.Duration(time2))
	fmt.Println(time2.Duration(time1))

	// Output:
	// 1s
	// 1s
}

func ExampleTime_Equal() {
	time1 := context.NewTime()
	time2 := context.NewTime()

	fmt.Println(time1.Equal(time2))

	time2 = time2.Increment()
	fmt.Println(time1.Equal(time2))

	// Output:
	// true
	// false
}

func ExampleTime_Before() {
	time1 := context.NewTime()
	time2 := context.NewTime()

	fmt.Println(time1.Before(time2))

	time2 = time2.Increment()
	fmt.Println(time1.Before(time2))
	fmt.Println(time2.Before(time1))

	// Output:
	// false
	// true
	// false
}

func ExampleTime_After() {
	time1 := context.NewTime()
	time2 := context.NewTime()

	fmt.Println(time1.After(time2))

	time2 = time2.Increment()
	fmt.Println(time1.After(time2))
	fmt.Println(time2.After(time1))

	// Output:
	// false
	// false
	// true
}

func ExampleTime_String() {
	time := context.NewTime()
	fmt.Println(time)

	time = time.ShiftBy(context.SampleRate()).ShiftBy(100)
	fmt.Println(time)

	time = time.ShiftBy(context.SampleRate()).ShiftBy(100)
	fmt.Println(time)

	// Output:
	// 0 seconds, sample 1/44100
	// 1 second, sample 101/44100
	// 2 seconds, sample 201/44100
}
