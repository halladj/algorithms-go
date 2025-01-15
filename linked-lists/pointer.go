package linkedlists

type Node struct {
	info interface{}
	suiv *Node
}

// Access implements Nodeist.
func (l *Node) Access(position int) (interface{}, error) {
	panic("unimplemented")
}

// Add implements Nodeist.
func (l *Node) Add(position int, element interface{}) error {
	panic("unimplemented")
}

// Delete implements Nodeist.
func (l *Node) Delete(element interface{}) error {
	panic("unimplemented")
}

// EmptyList implements Nodeist.
func (l *Node) EmptyList() bool {
	panic("unimplemented")
}

var _ LinkedList = (*Node)(nil)
