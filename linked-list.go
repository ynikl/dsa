package dsa

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

// [LeetCode题](https://leetcode.cn/problems/design-linked-list)
type MyLinkedList struct {
	Header *Node
	// 增加 Tail 字段可以快速在链表尾部增加节点
	Tail *Node
	Len  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Header: nil,
		Tail:   nil,
		Len:    0,
	}
}

func (ll *MyLinkedList) Get(i int) (value int) {

	if i > ll.Len-1 {
		panic("index out of range")
	}

	return ll.getByIndex(i).Value
}

func (ll *MyLinkedList) AddAtHead(value int) {

	node := newNode(value)

	node.Next = ll.Header
	ll.Header = node
	if ll.Tail == nil {
		ll.Tail = node
	}
	ll.Len++

}

func (ll *MyLinkedList) getByIndex(i int) *Node {

	dummy := ll.Header
	for index := 0; index < i; index++ {
		dummy = dummy.Next
	}

	return dummy
}

func (ll *MyLinkedList) Insert(i int, value int) {

	if i > ll.Len {
		panic("index out of range")
	}

	// just as append action
	if i == ll.Len {
		ll.AddAtTail(value)
		return
	}

	node := newNode(value)
	preNode := ll.getByIndex(i - 1)

	node.Next = preNode.Next
	preNode.Next = node
	ll.Len++
}

func (ll *MyLinkedList) AddAtTail(value int) {

	node := newNode(value)

	if ll.Len == 0 {
		ll.Header = node
	}

	if ll.Tail != nil {
		ll.Tail.Next = node
	}

	ll.Tail = node
	ll.Len++
}

func (ll *MyLinkedList) AddAtIndex(index, value int) {

	if index > ll.Len {
		return
	}

	node := newNode(value)

	if index <= 0 {
		node.Next = ll.Header
		ll.Header = node

		// append of last
	} else if index == ll.Len {

		if ll.Tail != nil {
			ll.Tail.Next = node
		}
		ll.Tail = node

	} else {

		// at middle
		prev := ll.getByIndex(index - 1)
		node.Next = prev.Next
		prev.Next = node
		fmt.Println(ll)
	}

	ll.Len++
}

func (ll *MyLinkedList) DeleteAtIndex(i int) {

	if i >= ll.Len {
		return
	}

	preNode := ll.getByIndex(i - 1)
	if i == ll.Len-1 {
		ll.Tail = preNode
		preNode.Next = nil
	} else if i == 0 {
		ll.Header = ll.Header.Next
	} else {
		preNode.Next = preNode.Next.Next
	}

	ll.Len--
}

func (ll *MyLinkedList) RemoveByValue(value int) {

	dummy := ll.Header
	var preNode *Node
	for index := 0; dummy != nil && index < ll.Len; index++ {

		if dummy.Value != value {
			continue
		}

		// remove the first element
		if preNode == nil {
			ll.Header = ll.Header.Next
		} else if dummy.Next == nil {
			ll.Tail = preNode
			preNode.Next = nil
		}

		preNode = dummy
		dummy = dummy.Next
		ll.Len--
	}
}

func (ll *MyLinkedList) String() string {

	sb := strings.Builder{}
	dummy := ll.Header
	for dummy != nil {
		sb.WriteString(strconv.Itoa(dummy.Value))

		dummy = dummy.Next
	}

	return sb.String()
}

func newNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}
