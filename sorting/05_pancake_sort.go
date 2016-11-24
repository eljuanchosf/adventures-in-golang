package sorting

//PancakeSort sorts an array of integers by iteratively flipping a subarray of elements
func PancakeSort(array []int) []int {
	elementCount := len(array) - 1
	for i := elementCount; i > 0; i-- {
		maxPos := findBiggestElement(array, 0, i)
		if maxPos != i {
			if maxPos != 0 {
				flipArray(array, 0, maxPos)
			}
			flipArray(array, 0, i)
		}
	}
	return array
}
