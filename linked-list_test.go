package dsa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
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

type LinkedListTestSuite struct {
	suite.Suite
	list *MyLinkedList
}

func (suite *LinkedListTestSuite) SetupTest() {
	list := Constructor()
	suite.list = &list
}

func (suite *LinkedListTestSuite) TestGet() {
	suite.Equal(-1, suite.list.Get(5), "not exist index return -1")
	suite.Equal(-1, suite.list.Get(-1), "not exist index return -1")
}

func (suite *LinkedListTestSuite) TestAddAtHead() {
	suite.list.AddAtHead(5)
	suite.Equal(5, suite.list.Get(0), "return add at head value")
	suite.list.AddAtHead(1)
	suite.Equal(1, suite.list.Get(0), "return add at head value")
}

func (suite *LinkedListTestSuite) TestInsert() {
	suite.list.Insert(2, 5)
	suite.Equal(-1, suite.list.Get(0), "insert out range nothing happened")
	suite.list.Insert(0, 1)
	suite.Equal(1, suite.list.Get(0), "insert at the header")
	suite.list.Insert(0, 2)
	suite.Equal(2, suite.list.Get(0), "insert at the header")

	suite.list.Insert(2, 3)
	suite.Equal(3, suite.list.Get(2), "insert at the tailer")
}

func TestLinkedList(t *testing.T) {
	suite.Run(t, new(LinkedListTestSuite))
}
