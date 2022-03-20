package main

import (
	"Heap/Heap"
	"fmt"
)

func setData(data []int, Heap *Heap.Heap, len int){
	for i := 0; i < len; i++ {
		Heap.Append(data[i])
	}
}

func main() {
	Heap := Heap.NewHeap()
	data := []int{27, 7, 80, 5, 67, 18, 62, 24, 58, 25}
	setData(data, &Heap, len(data))

	fmt.Println("MaxHeap :")
	Heap.Showdata()
	Heap.Maxheap()
	Heap.Showdata()

	Heap = Heap.NewHeap()
	data = []int{40, 23, 10, 15, 8}
	setData(data, &Heap, len(data))

	fmt.Println("MaxHeap Add and Remove :")
	Heap.Showdata()
	Heap.Maxheap()
	Heap.Showdata()

	fmt.Println("Add : 60")
	Heap.Append(60)
	Heap.Maxheap()
	Heap.Showdata()

	fmt.Println("Add : 20")
	Heap.Append(20)
	Heap.Maxheap()
	Heap.Showdata()

	fmt.Println("Remove : 60")
	Heap.RmData(60)
	Heap.Maxheap()
	Heap.Showdata()

	fmt.Println("Remove : 23")
	Heap.RmData(23)
	Heap.Maxheap()
	Heap.Showdata()

	Heap = Heap.NewHeap()
	data = []int{20, 30, 10, 50, 60, 40, 45, 5, 15, 25}
	setData(data, &Heap, len(data))

	fmt.Println("MinHeap :")
	Heap.Showdata()
	Heap.Minheap()
	Heap.Showdata()

	Heap = Heap.NewHeap()
	data = []int{20, 8, 28, 10, 4, 5, 40, 55}
	setData(data, &Heap, len(data))

	fmt.Println("Min-MaxHeap :")
	Heap.Showdata()
	Heap.MinMaxheap()
	Heap.Showdata()

	fmt.Println("Min-MaxHeap Add and Remove :")
	fmt.Println("Add : 2")
	Heap.Append(2)
	Heap.MinMaxheap()
	Heap.Showdata()

	fmt.Println("Remove : 40")
	Heap.RmData(40)
	Heap.MinMaxheap()
	Heap.Showdata()
}
