package main

import (
	"fmt"
)

func main() {
	vertices := make(map[string]int) // maps are like dictionaries (key,value)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	delete(vertices, "square")
	fmt.Println(vertices)
}
