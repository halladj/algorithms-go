package linkedlists

import (
	"errors"
	"fmt"
)

type Node struct {
	info interface{}
	suiv *Node
}

// Access implements Nodeist.
func (l *Node) Access(position int) (interface{}, error) {

	// TODO: fix implementation cz it is hot trash.
	if l == nil {
		return nil, errors.New("didn't find element")
	} else {

		if l.info == position {
			return l.info, nil
		} else {
			return l.suiv.Access(position)
		}
	}
}

// Add implements Nodeist.
func (l *Node) Add(position int, element interface{}) error {
	n, ok := element.(int)
	if !ok {
		return errors.New("invalid type")
	}

	if position <= 0 || position > l.ListLength()+1 {
		return errors.New("index out of Bounds")
	}

	k := &Node{info: n, suiv: nil}

	temp := l
	for i := 1; i < position; i++ {
		temp = temp.suiv
	}

	temp.suiv = k

	return nil
}

// Delete implements Nodeist.
func (l *Node) Delete(element interface{}) error {
	panic("unimplemented")
}

// EmptyList implements Nodeist.
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
	for temp != nil {
		fmt.Printf("%d", temp.info)
		fmt.Printf(" ,")
		temp = temp.suiv
	}
	fmt.Printf("]")
}

func (l *Node) ListLength() int {
	temp := l
	length := 0

	for temp != nil {
		temp = temp.suiv
		length++
	}

	return length
}

var _ LinkedList = (*Node)(nil)
