package sortedLinkedlist

import (
	"reflect"
	"sort"
	"testing"
)

var testArray []int
var expected []int

func buildLinkedList() (l LinkedList) {
	testArray = []int{23, 17, 8, 44}
	expected = testArray
	for _, a := range testArray {
		l.Add(a)
	}
	sort.Ints(expected)
	return
}

func Test_Add(t *testing.T) {
	l := buildLinkedList()
	want := 17
	if got := l.Top.Next.Next.Value; got != want {
		t.Errorf("LinkedList.Add(), expected %d, got %d", want, got)
	}
}

func Test_ToArray(t *testing.T) {
	l := buildLinkedList()
	if got := l.ToArray(); !reflect.DeepEqual(got, expected) {
		t.Errorf("LinkedList.ToArray(), expected %d, got %d", expected, got)
	}
}

func Test_RemoveFromStart(t *testing.T) {
	l := buildLinkedList()
	want := expected[1:]
	l.RemoveFromStart()
	if got := l.ToArray(); !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.RemoveFromStart(), expected %d, got %d", want, got)
	}
}

func Test_RemoveItem(t *testing.T) {
	l := buildLinkedList()
	first := expected[0:2]
	last := expected[3:]
	want := first
	want = append(want, last...)
	l.RemoveItem(23)
	if got := l.ToArray(); !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.RemoveItem(), expected %d, got %d", want, got)
	}
}
