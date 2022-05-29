package test

import (
	"fmt"
	"testing"
)

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func TestFunc(t *testing.T) {
	var slice []int
	println("cap0", cap(slice))
	slice = append(slice, 1, 2)
	slice = append(slice, 3)
	println("cap1", cap(slice))
	newSlice := AddElement(slice, 4)
	println("cap2", cap(slice))
	println("cap3", cap(newSlice))
	fmt.Println(&slice[0] == &newSlice[0])
}
