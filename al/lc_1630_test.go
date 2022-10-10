package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

// [4,6,5,9,3,7]
// [0,0,2]
// [2,3,5]
func TestLc1630(t *testing.T) {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	rand.Intn((6 - 5 + 1) >> 1)
	fmt.Printf("%v\n", checkArithmeticSubarrays(nums, l, r))
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n := len(l)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		ans[i] = true
		start, end := l[i], r[i]
		seg := make([]int, end-start+1)
		copy(seg, nums[start:end+1])
		sort.Ints(seg)
		d := seg[0] - seg[1]
		for j := 1; j < len(seg)-1; j++ {
			if seg[j]-seg[j+1] != d {
				ans[i] = false
				break
			}
		}
	}
	return ans
}
