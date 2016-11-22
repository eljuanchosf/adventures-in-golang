package main

//BubleSort sorts an array of integers by swapping two elements at once, iteratively
func BubleSort(array []int) []int {
	arrayLenght := len(array) - 1
	for {
		swapped := false
		for elementA := 0; elementA < arrayLenght; elementA++ {
			elementB := elementA + 1
			if array[elementA] > array[elementB] {
				swapped = true
				swapElement := array[elementA]
				array[elementA] = array[elementB]
				array[elementB] = swapElement
			}
		}
		if !swapped {
			break
		}
	}
	return array
}
