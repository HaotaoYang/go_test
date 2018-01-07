package my_sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for i := 1; i < data.Len(); i++ {
		for j := 0; j < data.Len()-i; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}

type IntArray []int

func (array IntArray) Len() int {
	return len(array)
}

func (array IntArray) Less(i, j int) bool {
	return array[i] > array[j]
}

func (array IntArray) Swap(i, j int) {
	// temp := array[i]
	// array[i] = array[j]
	// array[j] = temp
	array[i], array[j] = array[j], array[i]
}
