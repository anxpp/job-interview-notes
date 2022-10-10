package main

import (
	"fmt"
	"testing"
)

func TestLc0707(t *testing.T) {
	//["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
	//[[],[1],[3],[1,2],[1],[1],[1]]
	ll := Constructor()
	ll.AddAtHead(1)
	ll.Print()
	ll.AddAtTail(3)
	ll.Print()
	ll.AddAtIndex(1, 2)
	ll.Print()
	fmt.Println(ll.Get(1))
	ll.DeleteAtIndex(1)
	ll.Print()
	fmt.Println(ll.Get(1))
	fmt.Println("-----------------------")
	//["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
	//[[],				[7],		[2],		[1],		[3,0],			[2],			[6],		[4],	[4],	[4],		[5,0],			[6]]
	ll = Constructor()
	ll.AddAtHead(7)
	ll.Print()
	ll.AddAtHead(2)
	ll.Print()
	ll.AddAtHead(1)
	ll.Print()
	ll.AddAtIndex(3, 0)
	ll.Print()
	ll.DeleteAtIndex(2)
	ll.Print()
	ll.AddAtHead(6)
	ll.Print()
	ll.AddAtTail(4)
	ll.Print()
	fmt.Println(ll.Get(4))
	ll.AddAtHead(4)
	ll.Print()
	ll.AddAtIndex(5, 0)
	ll.Print()
	ll.AddAtHead(6)
	ll.Print()
	//["MyLinkedList","addAtHead","addAtTail","addAtHead","addAtTail","addAtHead","addAtHead","get","addAtHead","get","get","addAtTail"]
	//[[],				[7],		[7],			[9],		[8],		[6],		[0],	[5],	[0],	 [2],  [5],     [4]]
	ll = Constructor()
	ll.AddAtHead(7)
	ll.Print()
	ll.AddAtTail(7)
	ll.Print()
	ll.AddAtHead(9)
	ll.Print()
	ll.AddAtTail(8)
	ll.Print()
	ll.AddAtHead(6)
	ll.Print()
	ll.AddAtHead(0)
	ll.Print()
}

type MyLinkedList struct {
	head, tail *Node
	cnt        int
}

type Node struct {
	val  int
	prev *Node
	next *Node
}

func Constructor() MyLinkedList {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return MyLinkedList{head: head, tail: tail}
}

func (this *MyLinkedList) Print() {
	head := this.head
	for head != nil {
		fmt.Print(head.val)
		head = head.next
	}
	fmt.Println()
}

func insert(prev, next, cur *Node) {
	prev.next, next.prev = cur, cur
	cur.prev, cur.next = prev, next
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.cnt {
		return -1
	}
	if index <= this.cnt>>1 {
		// 左边找
		cur := this.head.next
		for index > 0 {
			cur = cur.next
			index--
		}
		return cur.val
	}
	// 右边找
	index = this.cnt - index
	cur := this.tail.prev
	for index > 1 {
		cur = cur.prev
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	insert(this.head, this.head.next, &Node{val: val})
	this.cnt++
}

func (this *MyLinkedList) AddAtTail(val int) {
	insert(this.tail.prev, this.tail, &Node{val: val})
	this.cnt++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.cnt {
		this.AddAtTail(val)
		return
	}
	if index > this.cnt {
		return
	}
	node := &Node{val: val}
	if index <= this.cnt>>1 {
		prev := this.head
		for index > 0 {
			prev = prev.next
			index--
		}
		insert(prev, prev.next, node)
	} else {
		next := this.tail
		index = this.cnt - index
		for index > 1 {
			next = next.prev
			index--
		}
		insert(next, next.prev, node)
	}
	this.cnt++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.cnt {
		return
	}
	if index <= this.cnt>>1 {
		prev := this.head
		for index > 0 {
			prev = prev.next
			index--
		}
		next := prev.next.next
		prev.next, next.prev = next, prev
	} else {
		next := this.tail
		index = this.cnt - index
		for index > 1 {
			next = next.prev
			index--
		}
		prev := next.prev.prev
		prev.next, next.prev = next, prev
	}
	this.cnt--
}
