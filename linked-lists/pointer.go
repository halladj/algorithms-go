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

// TODO: TO fix it we must start adding at the first so
// the first element would be the the last as follows
// [e1, e2, e3, MAX_VALUE], so then we dont have to deal with it.
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
func (l *Node) Delete(position int) error {
	if position <= 0 || position >= l.ListLength() {
		return errors.New("index out of Bounds")
	}

	temp := l
	for i := 1; i < position; i++ {
		temp = temp.suiv
	}

	temp.suiv = temp.suiv.suiv
	return nil
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

func (l *Node) Update(position int, value interface{}) error {
	n, ok := value.(int)
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

	temp.info = n
	return nil
}

var _ LinkedList = (*Node)(nil)
