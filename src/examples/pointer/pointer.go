package main

import (
	"fmt"
)

func main() {
	i := 7
	fmt.Println(i)  // print value
	fmt.Println(&i) // print memory address of variable

	// passing pointer to function
	inc(&i)
	fmt.Println(i)
	fmt.Println(&i)
}

// takes a pointer and increments its value. Doesn't return value
func inc(x *int) {
	*x++
}
