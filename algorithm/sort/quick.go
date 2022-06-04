package sort

import "math/rand"

func QuickSort(nums []Sortable, left, right int) {
	p := rand.Intn(right-left+1) + left
	pv := nums[p]
	l, r := left, right
	for l < r {
		for l <= p && nums[l].Less(pv) {
			l++
		}
		if l <= p {
			nums[p] = nums[l]
			p = l
		}
		for r >= p && pv.Less(nums[r]) {
			r--
		}
		if r >= p {
			nums[p] = nums[r]
			p = r
		}
	}
	nums[p] = pv
	if left < p {
		QuickSort(nums, left, p-1)
	}
	if right > p {
		QuickSort(nums, p+1, right)
	}
}
