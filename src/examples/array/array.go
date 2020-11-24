package main

import (
	"fmt"
)

func main() {
	var a [5]int
	a[2] = 7

	b := [2]int{2, 4}

	c := []int{2, 3, 45, 6} // slice of ints with no fixed length
	c = append(c, 12)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
