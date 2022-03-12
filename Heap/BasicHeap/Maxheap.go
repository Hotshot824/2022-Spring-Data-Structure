package BasicHeap

func max_heapify(arr *[Max]int, start int, end int) {
	dad := start
	son := dad * 2 + 1
	for son <= end {
		if son + 1 <= end && arr[son] < arr[son + 1] {
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
