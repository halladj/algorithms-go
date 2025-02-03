package linkedlists

import (
	"errors"
)

type List struct {
	tab  []int
	size int
}

// Add implements LinkedList.
func (l *List) Add(position int, element interface{}) error {
	if position < 0 && position > l.size {
		return errors.New("index out of iounds")
	}

	if l.size > 0 {
		for i := l.size - 1; i < position && l.size != position; i-- {
			l.tab[i+1] = l.tab[i]
		}
	}

	n, ok := element.(int)
	if !ok {
		return errors.New("invalid type")
	}

	l.tab[position] = n
	l.size = l.size + 1

	return nil
}

// Delete implements LinkedList.
func (l *List) Delete(position int) error {

	if position < 0 || position >= l.size {
		return errors.New("element out of bound")
	}
	// we are assuming that element is an index
	// and we gonna over-ride it.
	for i := position; i < l.size; i++ {
		l.tab[i] = l.tab[i+1]
	}

	l.size = l.size - 1
	return nil
}

// EmptyList implements LinkedList.
func (l *List) EmptyList() bool {
	return l.size == 0
}

// Find implements LinkedList.
func (l *List) Access(position int) (interface{}, error) {
	if position < 0 || position > l.size {
		return 0, errors.New("index out of iounds")
	}

	return l.tab[position], nil
}

func (l *List) GetArray() []int {
	return l.tab
}

func NewList() *List {
	t := make([]int, 50)
	return &List{tab: t, size: 0}
}

// TODO: write a toturial about this one....
/*
	var _ InterfaceName = TypeName{}
	// or
	var _ InterfaceName = (*TypeName)(nil)
*/
var _ LinkedList = (*List)(nil)

// TODO: write a toturial about  "type assertion" allows you to declare
// an interface value contains a certain concrete type or that its concrete
// type satisfies another interface.

// https://go.dev/ref/spec#Type_assertions
