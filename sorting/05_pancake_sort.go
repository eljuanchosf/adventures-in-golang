package sorting

import "github.com/eljuanchosf/adventures-in-golang/shared"

//PancakeSort sorts an array of integers by iteratively flipping a subarray of elements
func PancakeSort(array []int) []int {
	elementCount := len(array) - 1
	for i := elementCount; i > 0; i-- {
		maxPos := shared.FindBiggestElement(array, 0, i)
		if maxPos != i {
			if maxPos != 0 {
				shared.FlipArray(array, 0, maxPos)
			}
			shared.FlipArray(array, 0, i)
		}
	}
	return array
}
