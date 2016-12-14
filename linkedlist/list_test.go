package linkedlist

import (
	"reflect"
	"testing"

	"github.com/eljuanchosf/adventures-in-golang/shared"
)

var testArray []int

func buildLinkedList() (l LinkedList) {
	testArray = []int{23, 17, 8, 44}
	for _, a := range testArray {
		l.Add(a)
	}
	return
}

func Test_Add(t *testing.T) {
	l := buildLinkedList()
	want := 8
	if got := l.Top.Next.Next.Next.Value; got != want {
		t.Errorf("LinkedList.Add(), expected %d, got %d", want, got)
	}
}

func Test_ToArray(t *testing.T) {
	want := testArray
	l := buildLinkedList()
	if got := l.ToArray(); !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.ToArray(), expected %d, got %d", want, got)
	}
}

func Test_RemoveFromStart(t *testing.T) {
	l := buildLinkedList()
	want := testArray[1:]
	l.RemoveFromStart()
	if got := l.ToArray(); !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.RemoveFromStart(), expected %d, got %d", want, got)
	}
}

func Test_InsertAfter(t *testing.T) {
	l := buildLinkedList()
	first := testArray[0:2]
	last := shared.CopyArray(testArray[2:])
	want := first
	want = append(want, 66)
	want = append(want, last...)
	l.InsertAfter(17, 66)
	if got := l.ToArray(); !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.InsertAfter(), expected %d, got %d", want, got)
	}
}

func Test_RemoveItem(t *testing.T) {
	l := buildLinkedList()
	first := testArray[0:2]
	last := testArray[3:]
	want := first
	want = append(want, last...)
	l.RemoveItem(8)
	if got := l.ToArray(); !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.RemoveItem(), expected %d, got %d", want, got)
	}
}

func Test_InsertBefore(t *testing.T) {
	l := buildLinkedList()
	first := testArray[0:2]
	last := shared.CopyArray(testArray[2:])
	want := first
	want = append(want, 66)
	want = append(want, last...)
	l.InsertBefore(8, 66)
	if got := l.ToArray(); !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.InsertBefore(), expected %d, got %d", want, got)
	}
}
