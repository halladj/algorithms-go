package linkedlists

type LinkedList interface {
	EmptyList(l LinkedList) bool
	Add(position int, element interface{}) error
	Delete(element interface{}) error
	Find(position int) error
}
