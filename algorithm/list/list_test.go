package list

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestList(t *testing.T) {
	n := 1 << 6
	arr := make([]int, n)
	var arrayList, linkedList List = NewArrayList(), NewLinkedList()
	for i := 0; i < n; i++ {
		item := rand.Intn(n << 8)
		arr[i] = item
		arrayList.Add(item)
		linkedList.Add(item)
	}
	for i := 0; i < n; i++ {
		if v, _ := arrayList.Get(i); v.(int) != arr[i] {
			panic("arrayList error")
		}
		if v, _ := linkedList.Get(i); v.(int) != arr[i] {
			panic("linkedList error")
		}
	}
	for i := len(arr) >> 1; i >= 0; i-- {
		randIndex := rand.Intn(len(arr))
		randItem := arr[randIndex]
		arr = append(arr[:randIndex], arr[randIndex+1:]...)
		arrayList.Remove(randItem)
		linkedList.Remove(randItem)
	}
	for i := 0; i < n>>1; i++ {
		item := rand.Intn(n << 8)
		arr = append(arr, item)
		arrayList.Add(item)
		linkedList.Add(item)
	}
	for i := len(arr) >> 1; i >= 0; i-- {
		randIndex := rand.Intn(len(arr))
		randItem := arr[randIndex]
		arr = append(arr[:randIndex], arr[randIndex+1:]...)
		arrayList.Remove(randItem)
		linkedList.Remove(randItem)
	}
	if len(arr) != arrayList.Len() {
		panic("arrayList error")
	}
	if len(arr) != linkedList.Len() {
		panic("linkedList error")
	}
	fmt.Printf("arr: %v\n", arr)
	var helper = func(list List) {
		fmt.Print("arr: ")
		it := list.Iterator()
		fmt.Print("[")
		for it.HasNext() {
			fmt.Printf("%v ", it.Next())
		}
		fmt.Println("]")
	}
	helper(arrayList)
	helper(linkedList)
}
