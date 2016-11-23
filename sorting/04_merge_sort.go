package sorting

//MergeSort sorts an array of integers by iteratively splitting the array in two subarrays and ordering
//the subarrays and then merging the subarrays back into the main one.
func MergeSort(array []int) []int {
	elementCount := len(array)
	if elementCount <= 1 {
		return array
	}
	midway := (elementCount / 2)
	leftSide := array[:midway]
	rightSide := array[midway:]

	leftSide = MergeSort(leftSide)
	rightSide = MergeSort(rightSide)

	return merge(leftSide, rightSide)
}

func merge(leftSide []int, rightSide []int) (array []int) {
	for len(leftSide) > 0 && len(rightSide) > 0 {
		if leftSide[0] > rightSide[0] {
			array = append(array, rightSide[0])
			rightSide = rightSide[1:]
		} else {
			array = append(array, leftSide[0])
			leftSide = leftSide[1:]
		}
	}

	for len(leftSide) > 0 {
		array = append(array, leftSide[0])
		leftSide = leftSide[1:]
	}

	for len(rightSide) > 0 {
		array = append(array, rightSide[0])
		rightSide = rightSide[1:]
	}

	return array
}
