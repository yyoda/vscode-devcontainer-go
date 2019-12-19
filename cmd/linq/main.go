package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	fmt.Printf("input: %v\n", numbers)

	result := koazee.StreamOf(numbers).
		Reduce(func(acc, val int) int {
			return acc + val
		}).
		Int()

	fmt.Println(result)
}
