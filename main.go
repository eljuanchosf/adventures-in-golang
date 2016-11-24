package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"sort"
	"time"

	"github.com/eljuanchosf/adventures-in-golang/sorting"
)

var algorithms map[string]interface{}

const arraySize = 25000

func main() {
	tests := []struct {
		name  string
		array []int
	}{
		{"Random array", randomArray(arraySize)},
		{"Ordered array", orderedArray(arraySize)},
	}

	log.Printf("**** Array size is %d elements ****", arraySize)

	for _, test := range tests {
		for _, algorithm := range availableAlgorithms() {
			testAlgorithm(algorithm, test.array, test.name)
		}
		fmt.Println()
	}
}

func availableAlgorithms() []string {
	algorithms = make(map[string]interface{}, 6)
	algorithms["BubleSort"] = sorting.BubleSort
	algorithms["SelectiveSort"] = sorting.SelectiveSort
	algorithms["InsertionSort"] = sorting.InsertionSort
	algorithms["HeapSort"] = sorting.HeapSort
	algorithms["MergeSort"] = sorting.MergeSort
	algorithms["PancakeSort"] = sorting.PancakeSort

	sa := make([]string, 0, len(algorithms))
	for i := range algorithms {
		sa = append(sa, i)
	}
	sort.Strings(sa)
	return sa
}

func testAlgorithm(algorithm string, testArray []int, testType string) {
	start := time.Now()
	if _, err := call(algorithms, algorithm, testArray); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s - %13s = %s", testType, algorithm, time.Since(start))
}

func randomArray(elements int) (array []int) {
	array = []int{}
	for i := 0; i < elements; i++ {
		array = append(array, getRandom(0, elements))
	}
	return
}

func orderedArray(elements int) (array []int) {
	array = []int{}
	for i := 0; i < elements; i++ {
		array = append(array, i)
	}
	return
}

func getRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%13s took %s\n", name, elapsed)
}

func call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
