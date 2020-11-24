package main

import (
	"fmt"
)

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	c := 0
	for c < 5 {
		fmt.Println(c)
		c++
	}

	// for with arrays
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	// for with maps
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}
}
