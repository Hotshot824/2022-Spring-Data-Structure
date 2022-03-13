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

func (data *Heap) RmData(RmData int) {
	for i := 0; i < data.Last_index; i++{
		if (RmData == data.Array[i]){
			swap(&data.Array, i, data.Last_index - 1)
			data.Array[data.Last_index - 1] = 0
			data.Last_index--
		}
	}
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

func swap(arr *[Max]int, a int, b int){
	fmt.Printf("swap\t%d\t%d\n", arr[a], arr[b])
	arr[a], arr[b] = arr[b], arr[a]
}