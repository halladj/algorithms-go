package linkedlists

type LinkedList interface {
	EmptyList() bool
	Add(position int, element interface{}) error
	Delete(element interface{}) error
	Find(position int) error
}
