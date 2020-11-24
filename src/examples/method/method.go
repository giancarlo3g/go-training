package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Phone []int

func Print(s Stringer) {
	fmt.Println(s.String())
}

func (p Phone) String() string {
	text := ""
	for _, digit := range p {
		text += string('0' + digit)
	}
	return text
}

func main() {
	number := Phone{2, 3, 5, 7, 2}
	Print(number)
}
