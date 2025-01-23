package linkedlists

import "errors"

type Node struct {
	info interface{}
	suiv *Node
}

// Access implements Nodeist.
func (l *Node) Access(position int) (interface{}, error) {

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
	if position < 0 {
		return errors.New("index out of iounds")
	}

	n, ok := element.(int)
	if !ok {
		return errors.New("invalid type")
	}

	temp := l
	for i := 0; i < position; i++ {

		temp = temp.suiv
	}

	k := NewLinkedList()
	k.info = n
	k.suiv = temp.suiv
	temp = k
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

var _ LinkedList = (*Node)(nil)
