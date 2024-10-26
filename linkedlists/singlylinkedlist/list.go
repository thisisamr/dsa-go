package singlylinkedlist

import "fmt"

type Node[T interface{}] struct {
	Value T
	Next  *Node[T]
}

func new_node[T interface{}](v T) *Node[T] {
	return &Node[T]{
		Value: v,
		Next:  nil,
	}
}

type List[T interface{}] struct {
	Head *Node[T]
	Tail *Node[T]
	size int
}

func (l *List[T]) Length() int {
	return l.size
}

func (l *List[T]) Display() error {
	if l.Length() == 0 {
		return fmt.Errorf("error : List is empty")
	}
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
	return nil
}

func (l *List[T]) Prepend(value T) {
	nn := new_node(value)
	if l.Length() == 0 {
		l.Head = nn
		l.Head.Next = nil
		l.Tail = l.Head
		l.size++
	} else {
		temp := l.Head
		l.Head = nn
		l.Head.Next = temp
		l.size++
	}
}

func (l *List[T]) Append(value T) {
	//we can do i either by using the tail or just be traversing until node.next->nil
	nn := new_node(value)
	if l.Length() == 0 {

		l.Head = nn
		l.Head.Next = nil
		l.Tail = l.Head
	} else {
		l.Tail.Next = nn
		l.Tail = nn
	}
	l.size++
}

func (l *List[T]) InsertAt(value T, position int) error {
	//chaeck for position out of bounds
	if position < 1 || position > l.size+1 {
		return fmt.Errorf("position out of bounds")
	}
	nn := new_node(value)
	var prev *Node[T]
	current := l.Head

	for position > 1 {
		prev = current
		current = current.Next
		position--
	}
	if prev != nil {

		prev.Next = nn
		nn.Next = current
	} else {
		nn.Next = current
		l.Head = nn
	}
	l.size++
	return nil

}

//Deletes the last node in the list
func (l *List[T]) Pop() (interface{}, error) {
	if l.size == 0 {
		return nil, fmt.Errorf("cannot pop an empty list")
	}
	var prev *Node[T]
	current := l.Head
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	if prev != nil {
		prev.Next = nil
		l.Tail = prev
	} else {
		l.Head = nil
		l.Tail = nil
	}
	l.size--
	return current.Value, nil
}

func (l *List[T]) DeleteFirst() (interface{}, error) {
	if l.Length() == 0 {
		return nil, fmt.Errorf("Cannot preform delete on an empty list")
	}
	temp := l.Head
	if l.Head.Next == nil {
		l.Head = nil
		l.Tail = nil
	} else {
		l.Head = l.Head.Next
	}
	l.size--
	return temp.Value, nil
}

func (l *List[T]) DeleteAt(position int) (interface{}, error) {
	if position < 1 || position > l.Length() {
		return nil, fmt.Errorf("POsition out of bounds")
	}
	var prev *Node[T]
	current := l.Head
	for position > 1 {
		prev = current
		current = current.Next
		position--
	}
	if prev == nil {
		l.Head = nil
		l.Tail = nil
	} else {
		prev.Next = current.Next
	}
	l.size--
	return current.Value, nil

}
