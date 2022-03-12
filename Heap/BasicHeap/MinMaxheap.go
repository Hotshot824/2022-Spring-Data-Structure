package BasicHeap

import "fmt"

func checklevel(index int) int {
	count := 1
	for index != 1 {
		index /= 2
		count++
	}
	return count
}

func pushdonw_min(arr *[Max]int, start int, end int) {
	dad := start
	son := dad*2 + 1
	grandson := son * 2
	for son <= end {
		if grandson <= end {
			for grandson+1 < end && arr[grandson] > arr[grandson+1] {
				grandson++
			}
			if arr[grandson] < arr[dad] {
				arr[grandson], arr[dad] = arr[dad], arr[grandson]
			}
			if arr[grandson] > arr[son] {
				arr[grandson], arr[son] = arr[son], arr[grandson]
			}
			pushdonw_min(arr, grandson, end)
		} else {
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
	// fmt.Println("dad = ", dad, "\tson = ", son, "\tgrandson = ", grandson, "\t")
}

func pushdonw_max(arr *[Max]int, start int, end int) {
	dad := start
	son := dad*2 + 1
	grandson := son * 2
	for son <= end {
		if grandson <= end {
			for grandson+1 < end && arr[grandson] < arr[grandson + 1] {
				grandson++
			}
			if arr[grandson] > arr[dad] {
				arr[grandson], arr[dad] = arr[dad], arr[grandson]
			}
			if arr[grandson] < arr[son] {
				arr[grandson], arr[son] = arr[son], arr[grandson]
			}
			pushdonw_min(arr, grandson, end)
		} else {
			if son+1 <= end && arr[son] < arr[son+1] {
				son++
			}
			if arr[dad] < arr[son] {
				return
			} else {
				arr[dad], arr[son] = arr[son], arr[dad]
				dad = son
				son = dad*2 + 1
			}
		}
	}
	// fmt.Println("dad = ", dad, "\tson = ", son, "\tgrandson = ", grandson, "\t")
}

func (data *Heap) pushdown(index int) {
	if checklevel(index+1)%2 != 0 { //Min level
		fmt.Print("Min level\n")
		pushdonw_min(&data.Array, index, data.Last_index-1)
	} else {
		fmt.Print("Max level\n")
	}
}

func (data *Heap) MinMaxheap() {
	for i := data.Last_index/2 - 1; i >= 0; i-- {
		data.pushdown(i)
	}
}
