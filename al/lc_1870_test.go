package main

import (
	"testing"
)

func Test1870(t *testing.T) {
	println(minSpeedOnTime([]int{1, 1, 100000}, 2.01))
}

func minSpeedOnTime(dist []int, hour float64) int {
	var n = len(dist)
	var hr = int(int64(hour * 100))
	var check = func(m int) bool {
		// 前面花掉的时间
		var times int
		for i := 0; i < n-1; i++ {
			times += (dist[i] + m - 1) / m
		}
		// 最后的路程＜=剩余时间该速度能行驶的距离
		return dist[n-1]*100 <= (hr-times*100)*m
	}
	var l, r int = 1, 1e7
	for l < r {
		m := l + (r-l)>>1
		if check(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	if check(r) {
		return r
	}
	return -1
}
