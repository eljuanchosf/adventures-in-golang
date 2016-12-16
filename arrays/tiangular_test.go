package arrays

import "testing"

/*
Triangular arrays, test case base on this triangular array of distances:
     A   B   C   D
A    0  150 270 390
B   150  0  210 600
C   270 210  0  150
D   390 600 150  0
*/

func Test_New(t *testing.T) {
	n := 4
	ta := TriSlice{}
	ta.Initialize(n)
	want := 6
	if got := len(ta.Values); got != want {
		t.Errorf("ta.New(), expected lenght of %d, got %d", want, got)
	}
}

func Test_Add(t *testing.T) {
	ta := TriSlice{}
	ta.Initialize(4)
	ta.Add(0, 1, 150)
	ta.Add(0, 2, 270)
	ta.Add(0, 3, 390)

	//Count non zero elements
	nz := 0
	for _, val := range ta.Values {
		if val != 0 {
			nz++
		}
	}

	want := 3
	got := nz
	if want != got {
		t.Errorf("Expected %d non zero elements, got %d", want, got)
	}
}

func Test_Get(t *testing.T) {
	type args struct {
		row int
		col int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"First value", args{0, 1, 150}, 150},
		{"Second value", args{1, 2, 390}, 390},
		{"First value", args{0, 2, 270}, 270},
	}

	ta := TriSlice{}
	ta.Initialize(4)

	for _, test := range tests {
		ta.Add(test.args.row, test.args.col, test.args.val)
		if got := ta.Get(test.args.row, test.args.col); got != test.want {
			t.Errorf("%q, Get(), expected %d, got %d", test.name, test.want, got)
		}
	}

}
