package tone_test

import (
	"fmt"

	"github.com/green-aloe/enobox/context"
	"github.com/green-aloe/enobox/tone"
)

func ExampleNumHarmGains() {
	ctx := context.NewContext()
	numHarmGains := tone.NumHarmGains(ctx)

	fmt.Println(numHarmGains)

	// Output:
	// 20
}

func ExampleSetNumHarmGains() {
	for _, n := range []int{10, 100, tone.DefaultNumHarmGains} {
		tone.SetNumHarmGains(n)

		ctx := context.NewContext()
		numHarmGains := tone.NumHarmGains(ctx)

		fmt.Println(numHarmGains)
	}

	// Output:
	// 10
	// 100
	// 20
}
