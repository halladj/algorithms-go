package linkedlists

type LinkedList interface {
	EmptyList() bool
	Add(position int, element interface{}) error
	Delete(position int) error
	Access(position int) (interface{}, error)
}
