package queue

type ChainedQueue struct {
	tab    []interface{}
	tete   int
	queue  int
	pleint bool
}

// Dequeue implements Queue.
func (l *ChainedQueue) Dequeue() (interface{}, error) {
	panic("unimplemented")
}

// Enqueue implements Queue.
func (l *ChainedQueue) Enqueue(interface{}) error {
	panic("unimplemented")
}

// Queue implements Queue.
func (l *ChainedQueue) Queue() interface{} {
	panic("unimplemented")
}

// QueuePrint implements Queue.
func (l *ChainedQueue) QueuePrint() {
	panic("unimplemented")
}

// Tete implements Queue.
func (l *ChainedQueue) Tete() interface{} {
	panic("unimplemented")
}

var _ Queue = (*ChainedQueue)(nil)
