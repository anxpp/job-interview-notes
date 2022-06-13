package main

import (
	"fmt"
	"testing"
)

func TestLc0007(t *testing.T) {
	reverse(-123)
}

func reverse(x int) int {
	var ans int
	for x != 0 {
		if ans < (-1<<31)/10 || ans > (1<<31-1)/10 {
			return 0
		}
		fmt.Printf("x%%10=%d,x/10=%d", x%10, x/10)
		ans = ans*10 + x%10
		x /= 10
	}
	return ans
}
