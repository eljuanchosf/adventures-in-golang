package doubleLinkedlist

//Cell represents the basic structure of a cell
type Cell struct {
	Value int
	Next  *Cell
	Prev  *Cell
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
	current.Next.Prev = current
}

//InsertAfter inserts a value after the item
func (l *LinkedList) InsertAfter(item, value int) {
	newCell := &Cell{Value: value}
	current := &l.Top
	for current.Next != nil {
		if current.Value == item {
			newCell.Next = current.Next
			newCell.Prev = current
			current.Next = newCell
			newCell.Next.Prev = newCell
			break
		}
		current = current.Next
	}
}

//InsertBefore inserts a value after the item
func (l *LinkedList) InsertBefore(item, value int) {
	newCell := &Cell{Value: value}
	current := &l.Top
	for current.Next != nil {
		if current.Next.Value == item {
			newCell.Next = current.Next
			newCell.Prev = current
			current.Next = newCell
			newCell.Next.Prev = newCell
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
			current.Next.Prev = current
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
