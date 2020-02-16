package main

import (
	"fmt"
	"strings"
	"testing"
	"testing/quick"
)

func Pad(input string, limit uint) string {
	fmt.Printf("the original string is %s and the length we need is %d \n", input, limit)
	strLen := uint(len(input))

	if strLen > limit {
		return input[:limit]
	}
	input +=  strings.Repeat(" ", int(limit - strLen))
	return input
}

func main() {
	result := Pad("Hello", 25)
	fmt.Printf("the padded string is --------%s --------- \n", result)
}

func TestPad(t *testing.T){
	closureFn := func(inp string, l uint8) bool {
		res := Pad(inp, uint(l))
		return len(res) ==  int(l)
	}
	if err := quick.Check(closureFn,  &quick.Config{MaxCount: 100}); err != nil{
		t.Errorf("%s",err)
	}
}
