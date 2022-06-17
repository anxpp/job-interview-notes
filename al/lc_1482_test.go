package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func Test1482(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(1<<4) + 8
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(1 << 5)
	}
	sort.Ints(list)
	target := 24
	i := sort.Search(n, func(i int) bool {
		return list[i] >= target
	})
	fmt.Printf("n=%d,%v\nsearched: i=%d\n", n, list, i)

	fmt.Printf("%d\n", minDays([]int{1, 10, 3, 10, 2}, 3, 1))
}

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	var l, r = 1<<31 - 1, 0
	for _, n := range bloomDay {
		if n < l {
			l = n
		}
		if n > r {
			r = n
		}
	}
	var check = func(mid int) bool {
		var start, cnt int
		for i := range bloomDay {
			if bloomDay[i] < mid {
				start = i + 1
				continue
			}
			if i-start+1 >= k {
				cnt++
				if cnt >= m {
					return true
				}
				start = i + 1
			}
		}
		return false
	}
	for l < r {
		mid := l + (r-l)>>1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
