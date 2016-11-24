package sorting

type testCase struct {
	name  string
	array []int
	want  []int
}

func swapInt(array []int, positionA int, positionB int) {
	swapElement := array[positionA]
	array[positionA] = array[positionB]
	array[positionB] = swapElement
}

func findSmallestElement(array []int, startIndex, endIndex int) (position int) {
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

func findBiggestElement(array []int, startIndex, endIndex int) (position int) {
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

func compareArrays(a, b []int) bool {
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

func flipArray(array []int, startIndex, endIndex int) {
	i := startIndex
	j := endIndex
	for i < j {
		swapInt(array, i, j)
		i++
		j--
	}
}

func testCases() []testCase {
	return []testCase{
		{"First and last swapped", []int{6, 2, 3, 4, 5, 1}, []int{1, 2, 3, 4, 5, 6}},
		{"All scrambled", []int{4, 2, 1, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
		{"Only two", []int{2, 1}, []int{1, 2}},
		{"Only one", []int{1}, []int{1}},
	}
}
