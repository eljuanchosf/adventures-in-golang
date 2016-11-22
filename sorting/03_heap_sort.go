package sorting

//HeapSort sorts an array of integers by swapping two elements at once, iteratively
func HeapSort(array []int) []int {
	elementCount := len(array)
	for index := (elementCount / 2) - 1; index >= 0; index-- {
		heapify(array, elementCount, index)
	}
	for index := elementCount - 1; index >= 0; index-- {
		swapInt(array, 0, index)
		heapify(array, index, 0)
	}
	return array
}

func heapify(array []int, heapSize int, index int) {
	largest := index
	l := 2*index + 1
	r := 2*index + 2

	if l < heapSize && array[l] > array[largest] {
		largest = l
	}
	if r < heapSize && array[r] > array[largest] {
		largest = r
	}

	if largest != index {
		swapInt(array, index, largest)
		heapify(array, heapSize, largest)
	}
}
