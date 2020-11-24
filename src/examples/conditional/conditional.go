package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 7
	sum := x + y

	if sum > 6 {
		fmt.Println("Greater than 6")
	} else if sum < 2 {
		fmt.Println("Less than 2")
	} else {
		fmt.Println("Neither")
	}

}
