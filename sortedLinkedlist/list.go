package sortedLinkedlist

//Cell represents the basic structure of a cell
type Cell struct {
	Value int
	Next  *Cell
}

//LinkedList represents a linkedList
type LinkedList struct {
	Top Cell
}

//Add adds a new value to the LinkedList
func (l *LinkedList) Add(value int) {
	newCell := &Cell{Value: value}
	current := &l.Top
	for current.Next != nil && current.Next.Value < newCell.Value {
		current = current.Next
	}
	newCell.Next = current.Next
	current.Next = newCell
}

//RemoveItem inserts a value after the item
func (l *LinkedList) RemoveItem(item int) {
	current := &l.Top
	for current.Next != nil {
		if current.Next.Value == item {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
}

//RemoveFromStart deletes the first item in the list
func (l *LinkedList) RemoveFromStart() {
	l.Top = *l.Top.Next
}

//ToArray flattens the list to an array
func (l *LinkedList) ToArray() (arr []int) {
	current := l.Top.Next
	arr = append(arr, current.Value)
	for current.Next != nil {
		current = current.Next
		arr = append(arr, current.Value)
	}
	return
}
