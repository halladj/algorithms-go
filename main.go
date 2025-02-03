package main

import (
	linkedlists "github.com/halladj/algorithms/linked-lists"
)

func main() {

	//---> linkedList using Arrays
	//l := linkedlists.NewList()

	//l.Add(0, 1)
	//l.Add(1, 2)
	//l.Add(2, 3)
	//l.Add(3, 4)
	//l.Add(4, 5)
	//l.Add(5, 6)

	//fmt.Printf("%v\n", l.GetArray())

	//l.Delete(3)
	//fmt.Printf("%v\n", l.GetArray())

	// ---> linkedList using Pointers
	hamza := linkedlists.NewLinkedList()

	hamza.Add(1, 10)
	hamza.Add(2, 7)
	hamza.Add(3, 34)
	hamza.Add(4, 87)
	hamza.Add(5, 100)
	hamza.Add(6, 763)

	hamza.Update(1, 99)

	hamza.ListPrint()

	hamza.Delete(3)
	hamza.ListPrint()
}
