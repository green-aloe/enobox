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

func ExampleNewTimeAt() {
	time := context.NewTimeAt(3, 25, context.DefaultSampleRate)

	fmt.Println(time)

	// Output:
	// 3 seconds, sample 25/44100
}

func ExampleTime_Second() {
	time := context.NewTime()
	second := time.Second()

	fmt.Println(second)

	// Output:
	// 0
}

func ExampleTime_Sample() {
	time := context.NewTime()
	sample := time.Sample()

	fmt.Println(sample)

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

	duration1 := time1.Duration(time2)
	duration2 := time2.Duration(time1)

	fmt.Println(duration1, duration2)

	// Output:
	// 1s 1s
}

func ExampleTime_Equal() {
	time1 := context.NewTime()
	time2 := context.NewTime()

	isEqual1 := time1.Equal(time2)

	time2 = time2.Increment()
	isEqual2 := time1.Equal(time2)

	fmt.Println(isEqual1, isEqual2)

	// Output:
	// true false
}

func ExampleTime_Before() {
	time1 := context.NewTime()
	time2 := context.NewTime()

	before1a := time1.Before(time2)
	before1b := time2.Before(time1)

	time2 = time2.Increment()
	before2a := time1.Before(time2)
	before2b := time2.Before(time1)

	fmt.Println(before1a, before1b)
	fmt.Println(before2a, before2b)

	// Output:
	// false false
	// true false
}

func ExampleTime_After() {
	time1 := context.NewTime()
	time2 := context.NewTime()

	after1a := time1.After(time2)
	after1b := time2.After(time1)

	time2 = time2.Increment()
	after2a := time1.After(time2)
	after2b := time2.After(time1)

	fmt.Println(after1a, after1b)
	fmt.Println(after2a, after2b)

	// Output:
	// false false
	// false true
}

func ExampleTime_String() {
	time1 := context.NewTime()
	time2 := time1.ShiftBy(context.SampleRate()).ShiftBy(100)
	time3 := time2.ShiftBy(context.SampleRate()).ShiftBy(100)

	fmt.Println(time1)
	fmt.Println(time2)
	fmt.Println(time3)

	// Output:
	// 0 seconds, sample 1/44100
	// 1 second, sample 101/44100
	// 2 seconds, sample 201/44100
}
