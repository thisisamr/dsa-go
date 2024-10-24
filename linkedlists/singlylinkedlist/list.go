package singlylinkedlist

import "fmt"

type Node[T interface{}] struct {
	Value T
	Next  *Node[T]
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

func new_node[T interface{}](v T) *Node[T] {
	return &Node[T]{
		Value: v,
		Next:  nil,
	}
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
