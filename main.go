package main

import (
	"fmt"

	linkedlists "github.com/halladj/algorithms/linked-lists"
)

func main() {

	// ---> linkedList using Arrays
	//hamza := linkedlists.NewList()

	//hamza.Add(0, 10)
	//hamza.Add(1, 2)
	//hamza.Add(2, 4)
	//hamza.Add(3, 88)
	//hamza.Add(4, 8)

	//fmt.Printf("%v\n", hamza.GetArray())

	//hamza.Delete(3)
	//fmt.Printf("%v\n", hamza.GetArray())

	// ---> linkedList using Pointers
	hamza := linkedlists.NewLinkedList()

	err := hamza.Add(1, 10)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	err = hamza.Add(2, 7)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	hamza.ListPrint()

	v, err := hamza.Access(2)
	n, ok := v.(int)
	if err != nil || ok != true {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Printf("now this is the element %v", n)
}
