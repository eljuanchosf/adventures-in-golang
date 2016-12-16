package arrays

//TriSlice defines a triangular slice
type TriSlice struct {
	NumRows int
	Values  []int
}

//Initialize initializes a new triangular slice
func (ta *TriSlice) Initialize(rows int) {
	elemCount := rows * (rows - 1) / 2
	ta.NumRows = rows
	ta.Values = make([]int, elemCount)
}

//Add adds an element to the array
func (ta *TriSlice) Add(row, col, value int) {
	ta.Values[mapRowCol(row, col)] = value
}

//Get returns a value mapped for the array
func (ta *TriSlice) Get(row, col int) (val int) {
	return ta.Values[mapRowCol(row, col)]
}

//Size returns the number of elements in the array
func (ta *TriSlice) Size() int {
	return len(ta.Values)
}

func mapRowCol(row, col int) (idx int) {
	if row < col {
		tr := row
		row = col
		col = tr
	}
	return row*(row-1)/2 + col
}
