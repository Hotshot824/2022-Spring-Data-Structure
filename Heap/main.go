package main

import (
	"Heap/BasicHeap"
)

func main() {
	Heap := BasicHeap.NewHeap()

	data := []int{27, 7, 80, 5, 67, 18, 62, 24, 58, 25}
	for i := 0; i < len(data); i++ {
		Heap.Append(data[i])
	}

	Heap.Showdata()

	Heap.Maxheap()
	
	Heap.Showdata()
}
