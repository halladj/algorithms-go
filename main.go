package main

import (
	"fmt"

	linkedlists "github.com/halladj/algorithms/linked-lists"
)

func main() {

	hamza := linkedlists.NewList()

	hamza.Add(0, 10)
	hamza.Add(1, 2)
	hamza.Add(2, 4)
	hamza.Add(3, 88)
	hamza.Add(4, 8)

	fmt.Printf("%v\n", hamza.GetArray())

	hamza.Delete(3)
	fmt.Printf("%v\n", hamza.GetArray())
}
