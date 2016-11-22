package sorting

//InsertionSort sors an array by picking one array element at a time and inserting it into already sorted subarray.
func InsertionSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	elementCount := len(array)
	for element := 1; element < elementCount; element++ {
		tmp := array[element]
		subElement := element - 1
		for {
			if subElement >= 0 && tmp < array[subElement] {
				array[subElement+1] = array[subElement]
				subElement--
			} else {
				break
			}
		}
		array[subElement+1] = tmp
	}
	return array
}
