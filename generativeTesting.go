package main

import (
	"fmt"
	"strings"
)

func Pad(input string, limit uint) string {
	fmt.Printf("the original string is %s and the length we need is %d \n", input, limit)
	strLen := uint(len(input))

	if strLen > limit {
		return input[:limit-1]
	}
	input +=  strings.Repeat(" ", int(limit - strLen))
	return input
}

func main() {
	result := Pad("Hello", 25)
	fmt.Printf("the padded string is --------%s --------- \n", result)
}

