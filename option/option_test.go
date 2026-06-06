package option

import (
	"fmt"
	"strconv"
)

func ExampleOption() {
	v1, _ := Some(12345).
		Map(func(i int) int { return i * 2 }).
		Map(strconv.Itoa).
		OrElse("oops").
		Unwrap()
	fmt.Println(v1)

	v2, _ := None[int]().
		Map(func(i int) int { return i * 2 }).
		Map(strconv.Itoa).
		OrElse("oops").
		Unwrap()
	fmt.Println(v2)
	// Output:
	// 24690
	// oops
}
