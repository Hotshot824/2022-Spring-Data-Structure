package main

import (
	"Heap/BasicHeap"
)

func main() {
	Heap := BasicHeap.NewHeap()

	data := []int{20, 8, 28, 10, 4, 5, 40, 55}
	for i := 0; i < len(data); i++ {
		Heap.Append(data[i])
	}

	Heap.Showdata()

	Heap.MinMaxheap()

	Heap.Showdata()
}
