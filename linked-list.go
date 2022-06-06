package dsa

type Node struct {
	Value int
	Next  *Node
}

//

type LinkedList struct {
	Header *Node
	Len    int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Header: nil,
		Len:    0,
	}
}

func (ll *LinkedList) Index(i int) (value int) {
	return 0
}

func (ll *LinkedList) Insert(i int, value int) {
}

func (ll *LinkedList) Append(value int) {
}

func (ll *LinkedList) RemoveByIndex(i int) {
}

func (ll *LinkedList) RemoveByValue(i int) {
}
