package linkedlists

type LinkedL struct {
	info interface{}
	suiv *LinkedL
}

// Access implements LinkedList.
func (l *LinkedL) Access(position int) (interface{}, error) {
	panic("unimplemented")
}

// Add implements LinkedList.
func (l *LinkedL) Add(position int, element interface{}) error {
	panic("unimplemented")
}

// Delete implements LinkedList.
func (l *LinkedL) Delete(element interface{}) error {
	panic("unimplemented")
}

// EmptyList implements LinkedList.
func (l *LinkedL) EmptyList() bool {
	panic("unimplemented")
}

var _ LinkedList = (*LinkedL)(nil)
