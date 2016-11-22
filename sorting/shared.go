package sorting

func swapInt(array []int, positionA int, positionB int) {
	swapElement := array[positionA]
	array[positionA] = array[positionB]
	array[positionB] = swapElement
}

func findSmallestElement(array []int, startIndex int) (position int) {
	value := array[startIndex]
	position = startIndex
	for index := startIndex; index < len(array); index++ {
		if array[index] < value {
			value = array[index]
			position = index
		}
	}
	return
}
