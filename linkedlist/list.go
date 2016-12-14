package linkedlist

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
	current := &l.Top
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Cell{Value: value}
}

//InsertAfter inserts a value after the item
func (l *LinkedList) InsertAfter(item, value int) {
	current := &l.Top
	for current.Next != nil {
		if current.Value == item {
			tempCell := current.Next
			current.Next = &Cell{value, tempCell}
			break
		}
		current = current.Next
	}
}

//InsertBefore inserts a value after the item
func (l *LinkedList) InsertBefore(item, value int) {
	current := &l.Top
	for current.Next != nil {
		if current.Next.Value == item {
			newCell := &Cell{value, current.Next}
			current.Next = newCell
			break
		}
		current = current.Next
	}
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
