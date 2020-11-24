package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result := sum(2, 3)
	fmt.Println(result)

	value, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

// accepts two integers, return an integer
func sum(x int, y int) int {
	return x + y
}

// accepts one float, returns float and error
// Go does not have exceptions
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
