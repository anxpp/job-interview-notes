package main

import "testing"

func Test207(t *testing.T) {
	println(canFinish(2, [][]int{{0, 1}, {1, 0}}))
}

type Node struct {
	val  int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) Put(v int) {
	node := &Node{val: v}
	if q.head == nil {
		q.head, q.tail = node, node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue) Remove(v int) {
	var pre, cur *Node = nil, q.head
	for cur.val != v {
		pre, cur = cur, cur.next
	}
	if pre == nil {
		q.head = cur.next
	} else {
		pre.next = cur.next
	}
}

func (q Queue) IsEmpty() bool {
	return q.head == nil
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// key依赖哪些课程
	m := map[int]*Queue{}
	// 哪些课程依赖key
	mTo := map[int]*Queue{}
	for _, item := range prerequisites {
		if m[item[0]] == nil {
			m[item[0]] = &Queue{}
		}
		if mTo[item[1]] == nil {
			mTo[item[1]] = &Queue{}
		}
		m[item[0]].Put(item[1])
		mTo[item[1]].Put(item[0])
	}
	var list Queue
	for i := 0; i < numCourses; i++ {
		list.Put(i)
	}
	var changed = true
	for changed {
		changed = false
		var pre *Node
		for item := list.head; item != nil; pre, item = item, item.next {
			n := item.val
			// n依赖的课程数为0
			if m[n].IsEmpty() {
				// 依赖n的课程中移出n
				for head := mTo[n].head; head != nil; head = head.next {
					t := m[head.val]
					t.Remove(n)
				}
				changed = true
				delete(mTo, n)
				if pre == nil {
					list.head = item.next
				} else {
					pre.next = item.next
				}
			}
		}
	}
	return list.IsEmpty()
}
