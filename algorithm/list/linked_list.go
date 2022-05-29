package list

import "errors"

type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList() *LinkedList {
	dummy := &Node{
		next: nil,
	}
	return &LinkedList{
		head: dummy,
		tail: dummy,
	}
}

type LinkedListIterator struct {
	al  *LinkedList
	cur *Node
}

func (ll *LinkedList) Iterator() Iterator {
	return &LinkedListIterator{
		al:  ll,
		cur: ll.head,
	}
}

func (lli *LinkedListIterator) HasNext() bool {
	return lli.cur.next != nil
}

func (lli *LinkedListIterator) Next() interface{} {
	v := lli.cur.next
	lli.cur = lli.cur.next
	return v.val
}

func (ll *LinkedList) Len() int {
	return ll.length
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.length == 0
}

func (ll *LinkedList) Cap() int {
	return ll.length
}

func (ll *LinkedList) Add(x interface{}) {
	ll.tail.next = &Node{
		val: x,
	}
	ll.tail = ll.tail.next
	ll.length++
}

func (ll *LinkedList) FindIndex(x interface{}) int {
	var index int
	for item := ll.head.next; item != nil; item = item.next {
		if item.val == x {
			return index
		}
		index++
	}
	return -1
}

func (ll *LinkedList) Remove(x interface{}) bool {
	var index int
	if index = ll.FindIndex(x); index < 0 {
		return false
	}
	var pre = ll.head
	for index > 0 {
		pre = pre.next
		index--
	}
	pre.next = pre.next.next
	ll.length--
	return true
}

func (ll *LinkedList) Containers(x interface{}) bool {
	return ll.FindIndex(x) >= 0
}

func (ll *LinkedList) Clear() {
	ll.head.next = nil
	ll.tail = ll.head
	ll.length = 0
}

func (ll *LinkedList) GetNode(i int) (node *Node, e error) {
	if i < 0 || i >= ll.Len() {
		e = errors.New("index out of range")
		return
	}
	node = ll.head.next
	for n := 0; n < i; n++ {
		node = node.next
	}
	return
}

func (ll *LinkedList) Get(i int) (x interface{}, e error) {
	var n *Node
	if n, e = ll.GetNode(i); e == nil {
		x = n.val
	}
	return
}
