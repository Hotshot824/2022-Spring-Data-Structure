package BasicHeap

import "fmt"

const (
	Max = 100
)

type Heap struct {
	Array      [Max]int
	Last_index int
}

func NewHeap() Heap {
	return Heap{
		Last_index: 0,
	}
}

func (data *Heap) Append(Newdata int) {
	data.Array[data.Last_index] = Newdata
	data.Last_index += 1
}

func (data *Heap) Showdata() {
	for i := 0; i < data.Last_index; i++ {
		fmt.Print(data.Array[i], "\t")
		if (i > 0) && (i % 10 == 0){
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
}

func max_heapify(arr *[Max]int, start int, end int) {
	dad := start
	son := dad*2 + 1
	for son <= end {
		if son+1 <= end && arr[son] < arr[son+1] {
			son++
		}
		if arr[dad] > arr[son] {
			return
		} else {
			arr[dad], arr[son] = arr[son], arr[dad]
			dad = son
			son = dad*2 + 1
		}
	}
}

func (data *Heap) Maxheap() {
	for i := data.Last_index/2 - 1; i >= 0; i-- {
		max_heapify(&data.Array, i, data.Last_index-1)
	}
}
