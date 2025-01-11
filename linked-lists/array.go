package linkedlists

import "errors"

type List struct {
	tab  []int
	size int
}

// Add implements LinkedList.
func (l *List) Add(position int, element interface{}) error {
	if l.size < 0 {
		return errors.New("index out of iounds")
	}

	for i := l.size - 1; i < position; i-- {
		l.tab[i] = l.tab[i-1]
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
func (l *List) Delete(element interface{}) error {
	n, ok := element.(int)
	if !ok {
		return errors.New("invalid type")
	}

	// TODO: find element and delete it.
	return nil
}

// EmptyList implements LinkedList.
func (l *List) EmptyList() bool {
	panic("unimplemented")
}

// Find implements LinkedList.
func (l *List) Find(position int) error {
	panic("unimplemented")
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
