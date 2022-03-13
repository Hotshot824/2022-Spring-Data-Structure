package BasicHeap

func checklevel(index int) int {
	count := 1
	for index != 1 {
		index /= 2
		count++
	}
	return count
}

func max_grandson(arr *[Max]int, i int, end int) int {
	result := i
	for j := 1; j < 4; j++ {
		if i+j <= end {
			if arr[i+j] > arr[result] {
				result = i + j
			}
		}
	}
	return result
}

func min_grandson(arr *[Max]int, i int, end int) int {
	result := i
	for j := 1; j < 4; j++ {
		if i+j <= end {
			if arr[i+j] < arr[result] {
				result = i + j
			}
		}
	}
	return result
}

func max_son(arr *[Max]int, son int, end int) int {
	result := son
	if son+1 <= end && arr[son+1] > arr[son] {
		result++
	}
	return result
}

func min_son(arr *[Max]int, son int, end int) int {
	result := son
	if son+1 <= end && arr[son+1] < arr[son] {
		result++
	}
	return result
}

func parent(i int) int {
	if i%2 == 0 {
		return (i / 2) + 1
	} else {
		return i / 2
	}
}

func pushdonw_min(arr *[Max]int, start int, end int) {
	dad := start
	son := dad*2 + 1
	grandson := son*2 + 1
	if son <= end { //If has son node
		if grandson <= end { //If has grandson node
			grandson := min_grandson(arr, grandson, end) //Selection minimum grandson node and compare
			if arr[grandson] < arr[dad] {
				swap(arr, grandson, dad)
				son = max_son(arr, son, end)
				if arr[grandson] > arr[parent(grandson)] {
					swap(arr, grandson, parent(grandson))
				}
				pushdonw_min(arr, grandson, end) //Recursive by grandson node
			}
		} else { //Not exist grandson
			son = min_son(arr, son, end) 
			if arr[son] < arr[dad] {
				swap(arr, son, dad)
			}
		}
	}
}

func pushdonw_max(arr *[Max]int, start int, end int) {
	dad := start
	son := dad*2 + 1
	grandson := son*2 + 1
	if son <= end { //If has son node
		if grandson <= end { //If has grandson node
			grandson := max_grandson(arr, grandson, end) //Selection minimum grandson node and compare
			if arr[grandson] > arr[dad] {
				swap(arr, grandson, dad)
				if arr[grandson] < arr[parent(grandson)] {
					swap(arr, grandson, parent(grandson))
				}
				pushdonw_max(arr, grandson, end) //Recursive by grandson node
			}
		} else { //Not exist grandson
			son = max_son(arr, son, end)
			if arr[son] > arr[dad] {
				swap(arr, son, dad)
			}
		}
	}
}

func (data *Heap) pushdown(index int) {
	if checklevel(index+1)%2 != 0 { //if node on min level
		pushdonw_min(&data.Array, index, data.Last_index-1)
	} else {
		pushdonw_max(&data.Array, index, data.Last_index-1)
	}
}

func (data *Heap) MinMaxheap() { //Buttom up
	for i := data.Last_index/2 - 1; i >= 0; i-- {
		data.pushdown(i)
	}
}
