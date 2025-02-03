package linkedlists

import (
	"errors"
	"fmt"
)

type Node struct {
	info interface{}
	suiv *Node
}

// Access implements NodeList.
func (l *Node) Access(position int) (interface{}, error) {
	if l == nil {
		return nil, errors.New("didn't find element")
	}
	if position <= 0 || position >= l.ListLength() {
		return 0, errors.New("index out of iounds")
	}

	// can not find a recursive algo !!!!
	//if l.info == position {
	//return l.info, nil
	//} else {
	//return l.suiv.Access(position)
	//}
	temp := l
	for i := 1; i < position; i++ {
		temp = temp.suiv
	}
	return temp.info, nil

}

// Add implements NodeList.
func (l *Node) Add(position int, element interface{}) error {
	n, ok := element.(int)
	if !ok {
		return errors.New("invalid type")
	}

	if position <= 0 || position > l.ListLength()+1 {
		return errors.New("index out of Bounds")
	}

	temp := l
	for i := 1; i < position; i++ {
		temp = temp.suiv
	}

	k := &Node{info: n, suiv: nil}
	k.suiv = temp.suiv
	temp.suiv = k

	return nil
}

// Delete implements NodeList.
func (l *Node) Delete(element interface{}) error {
	panic("unimplemented")
}

// EmptyList implements NodeList.
func (l *Node) EmptyList() bool {
	return l == nil
}

func NewLinkedList() *Node {
	return &Node{
		info: 0,
		suiv: nil,
	}
}

func (l *Node) ListPrint() {
	var temp = l

	fmt.Printf("[")
	for temp.suiv != nil {
		fmt.Printf("%d", temp.info)
		fmt.Printf(" ,")
		temp = temp.suiv
	}
	// Last element, so it does not have a trailing ",".
	fmt.Printf("%d", temp.info)
	fmt.Printf("]\n")
}

func (l *Node) ListLength() int {
	temp := l
	length := 1

	for temp != nil {
		temp = temp.suiv
		length++
	}

	return length
}

var _ LinkedList = (*Node)(nil)
