package queue

// First in First out [ Gets in from tail, goes out from the head ]
type Queue interface {
	// EnQueue the Tail
	Enqueue(interface{}) error
	// DeQueue the head
	Dequeue() (interface{}, error)

	QueuePrint()
}
