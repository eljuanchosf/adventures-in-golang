package shared

import "math"

//TestCase defines the basic structure of a test case
type TestCase struct {
	Name  string
	Array []int
	Want  []int
}

//SwapInt swaps two positions in an array of integers
func SwapInt(array []int, positionA, positionB int) {
	swapElement := array[positionA]
	array[positionA] = array[positionB]
	array[positionB] = swapElement
}

//FindSmallestElement finds the smallest element in an array of integers
func FindSmallestElement(array []int, startIndex, endIndex int) (position int) {
	value := array[startIndex]
	position = startIndex
	for index := startIndex; index <= endIndex; index++ {
		if array[index] < value {
			value = array[index]
			position = index
		}
	}
	return
}

//FindBiggestElement finds the biggest element in an array of integers
func FindBiggestElement(array []int, startIndex, endIndex int) (position int) {
	value := array[startIndex]
	position = startIndex
	for index := startIndex; index <= endIndex; index++ {
		if array[index] > value {
			value = array[index]
			position = index
		}
	}
	return
}

//CompareArrays compares two arrays to see if they match element by element
func CompareArrays(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//FlipArray flips an array staring in position startIndex up to position endIndex
func FlipArray(array []int, startIndex, endIndex int) {
	i := startIndex
	j := endIndex
	for i < j {
		SwapInt(array, i, j)
		i++
		j--
	}
}

//CopyArray copies an array of integers into another
func CopyArray(inputArray []int) (array []int) {
	array = make([]int, len(inputArray))
	copy(array, inputArray)
	return
}

//TestCases returns an array of test cases for sorting methods
func TestCases() []TestCase {
	return []TestCase{
		{"First and last swapped", []int{6, 2, 3, 4, 5, 1}, []int{1, 2, 3, 4, 5, 6}},
		{"All scrambled", []int{4, 2, 1, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
		{"Only two", []int{2, 1}, []int{1, 2}},
		{"Only one", []int{1}, []int{1}},
	}
}

//Avg returns a float32 that represents the average value for an array of integers
func Avg(array []int) (avg float32) {
	sum := 0
	for _, element := range array {
		sum += element
	}
	avg = float32(sum) / float32(len(array))
	return
}

//Variance returns a float32 that represents the variance value for an array of integers
func Variance(array []int) (variance float32) {
	var sum float32
	mean := Avg(array)
	sum = 0.0
	for _, element := range array {
		sum += (float32(element) - mean) * (float32(element) - mean)
	}
	variance = sum / float32(len(array))
	return
}

//StdDev returns a float32 that represents the standard deviation value for an array of integers
func StdDev(array []int) (stdDev float32) {
	return float32(math.Sqrt(float64(Variance(array))))
}
