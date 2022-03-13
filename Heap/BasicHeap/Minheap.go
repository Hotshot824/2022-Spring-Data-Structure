package BasicHeap

func min_heapify(arr *[Max]int, start int, end int) {
	dad := start
	son := dad*2 + 1
	for son <= end {
		if son+1 <= end && arr[son] > arr[son + 1] {
			son++
		}
		if arr[dad] < arr[son] {
			return
		} else {
			swap(arr ,dad, son)
			dad = son
			son = dad*2 + 1
		}
	}
}

func (data *Heap) Minheap() {
	for i := data.Last_index/2 - 1; i >= 0; i-- {
		min_heapify(&data.Array, i, data.Last_index-1)
	}
}