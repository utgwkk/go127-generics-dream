package pipeline

import (
	"fmt"
	"strconv"
)

func ExamplePipelineErr() {
	x, err := ApplyErr(
		"abcde",
		NewErr(strconv.Atoi).
			Then(New(func(i int) int {
				return i * 10
			}).Err()).
			Then(New(strconv.Itoa).Err()),
	)
	fmt.Println(x)
	fmt.Println(err)
	// Output:
	// strconv.Atoi: parsing "abcde": invalid syntax
}
