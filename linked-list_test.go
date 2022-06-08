package dsa

import (
	"fmt"
	"testing"
)

var list *MyLinkedList

func TestMain(m *testing.M) {

	l := Constructor()
	list = &l

	list.AddAtTail(1)
	list.AddAtTail(2)
	list.AddAtTail(3)
	list.AddAtTail(4) // idnex 3

	m.Run()
}

func TestIndex(t *testing.T) {

	if 4 != list.Get(3) {
		t.Error("simple append and index fail")
	}
}

func TestRemoveByValue(t *testing.T) {

	list.RemoveByValue(1)
	fmt.Println(list)
}

func TestAddAtTail(t *testing.T) {

	list.AddAtTail(5)
	fmt.Println(list)
}

func TestAddAtHead(t *testing.T) {

	list.AddAtHead(1)
	fmt.Println(list)
}

func TestCase1(t *testing.T) {
	// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
	// [[],[1],[3],[1,2],[1],[1],[1]]
	l := Constructor()
	list := &l
	list.AddAtHead(1)
	list.AddAtTail(3)
	fmt.Println(list)
	list.AddAtIndex(1, 2)
	fmt.Println(list)
	list.Get(1)
	fmt.Println(list)
	list.DeleteAtIndex(1)
	fmt.Println(list)
	list.Get(1)
	fmt.Println(list)
}

func TestCase2(t *testing.T) {
	// ["MyLinkedList" , "addAtHead" , "get" , "addAtHead" , "addAtHead" , "deleteAtIndex" , "addAtHead" , "get" , "get" , "get" , "addAtHead" , "deleteAtIndex"]
	// [[]             , [4]         , [1]   , [1]         , [5]         , [3]             , [7]         , [3]   , [3]   , [3]   , [1]         , [4]]
	l := Constructor()
	list := &l
	list.AddAtHead(4)
	list.AddAtHead(1)
	list.AddAtHead(5) // 514
	list.DeleteAtIndex(3)
	list.AddAtHead(7)
	list.AddAtHead(1) // 17514
	list.DeleteAtIndex(4)
	fmt.Println(list)
}
