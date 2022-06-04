package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	filename := "profile"
	f, _ := os.Create(filename)
	//_ = pprof.StartCPUProfile(f)
	defer func() {
		_ = pprof.WriteHeapProfile(f)
	}()
	//defer pprof.StopCPUProfile()
	rand.Seed(time.Now().Unix())
	const l = 1 << 20
	nums := make([]int, l)
	for i := range nums {
		nums[i] = rand.Intn(l)
	}
	defer func() func() {
		now := time.Now()
		return func() {
			fmt.Printf("time duration: %d\n", time.Now().Sub(now)/time.Millisecond)
		}
	}()()
	sort20220525(nums, 0, l-1)
	//fmt.Printf("%v\n", nums)
}

func sort20220525(nums []int, left, right int) {
	p := rand.Intn(right-left+1) + left
	pv := nums[p]
	l, r := left, right
	for l < r {
		for l <= p && nums[l] <= pv {
			l++
		}
		if l <= p {
			nums[p] = nums[l]
			p = l
		}
		for r >= p && nums[r] >= pv {
			r--
		}
		if r >= p {
			nums[p] = nums[r]
			p = r
		}
	}
	nums[p] = pv
	if left < p {
		sort20220525(nums, left, p-1)
	}
	if right > p {
		sort20220525(nums, p+1, right)
	}
}

func sort20220524(nums []int, left, right int) {
	p := rand.Intn(right-left+1) + left
	pv := nums[p]
	l, r := left, right
	for l < r {
		for l <= p && nums[l] <= pv {
			l++
		}
		if l <= p {
			nums[p] = nums[l]
			p = l
		}
		for r >= p && nums[r] >= pv {
			r--
		}
		if r >= p {
			nums[p] = nums[r]
			p = r
		}
	}
	nums[p] = pv
	var wg sync.WaitGroup
	if left < p {
		wg.Add(1)
		go func() {
			sort20220524(nums, left, p-1)
			wg.Done()
		}()
	}
	if right > p {
		wg.Add(1)
		go func() {
			sort20220524(nums, p+1, right)
			wg.Done()
		}()
	}
	wg.Wait()
}

func sort0(nums []int, left, right int) {
	p, pv := left, nums[left]
	l, r := left, right
	for l < r {
		for l <= p && nums[l] <= pv {
			l++
		}
		if l <= p {
			nums[p] = nums[l]
			p = l
		}
		for r >= p && nums[r] >= pv {
			r--
		}
		if r >= p {
			nums[p] = nums[r]
			p = r
		}
	}
	nums[p] = pv
	if left < p {
		sort0(nums, left, p-1)
	}
	if right > p {
		sort0(nums, p+1, right)
	}
}
