package pipeline

import (
	"fmt"
	"strconv"
	"strings"
)

func ExamplePipeline() {
	x := Apply(
		12345,
		New(func(x int) int {
			return x * 2
		}).
			Then(strconv.Itoa).
			Then(func(s string) []string {
				return strings.Split(s, "")
			}).
			Then(func(s []string) string {
				return strings.Join(s, "-")
			}))
	fmt.Println(x)
	// Output:
	// 2-4-6-9-0
}
