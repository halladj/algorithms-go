package queue

type ArrayQueue struct {
	tab    []interface{}
	tete   int
	queue  int
	pleint bool
}

// Dequeue implements Queue.
func (a *ArrayQueue) Dequeue() (interface{}, error) {
	panic("unimplemented")
}

// Enqueue implements Queue.
func (a *ArrayQueue) Enqueue(interface{}) error {
	panic("unimplemented")
}

// Queue implements Queue.
func (a *ArrayQueue) Queue() interface{} {
	panic("unimplemented")
}

// QueuePrint implements Queue.
func (a *ArrayQueue) QueuePrint() {
	panic("unimplemented")
}

// Tete implements Queue.
func (a *ArrayQueue) Tete() interface{} {
	panic("unimplemented")
}

var _ Queue = (*ArrayQueue)(nil)
