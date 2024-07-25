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
