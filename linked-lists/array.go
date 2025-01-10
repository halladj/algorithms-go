package linkedlists

type List struct {
	tab  []int
	size int
}

// Add implements LinkedList.
func (l List) Add(position int, element interface{}) error {
	panic("unimplemented")
}

// Delete implements LinkedList.
func (l List) Delete(element interface{}) error {
	panic("unimplemented")
}

// EmptyList implements LinkedList.
func (List) EmptyList(l LinkedList) bool {
	panic("unimplemented")
}

// Find implements LinkedList.
func (l List) Find(position int) error {
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
var _ LinkedList = List{}
