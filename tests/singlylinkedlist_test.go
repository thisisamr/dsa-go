package tests

import (
	sll "github.com/thisisamr/dsa-go/linkedlists/singlylinkedlist"
	"testing"
)

func TestPrepend(t *testing.T) {
	list := sll.List[int]{}
	list.Prepend(1)
	if list.Head.Value != 1 || list.Length() != 1 {
		t.Error("Prepend failed: head should have value 1 and size should be 1")
	}
}

func TestAppend(t *testing.T) {
	list := sll.List[int]{}
	list.Append(2)
	if list.Tail.Value != 2 || list.Length() != 1 {
		t.Error("Append failed: tail should have value 2 and size should be 1")
	}
}

func TestInsertAt(t *testing.T) {
	list := sll.List[int]{}
	list.Prepend(1)
	list.Append(3)
	err := list.InsertAt(2, 2)
	if err != nil || list.Head.Next.Value != 2 {
		t.Error("InsertAt failed at position 2")
	}
	err = list.InsertAt(4, 5) // Out of bounds
	if err == nil {
		t.Error("Expected error for out-of-bounds insert")
	}
}

func TestPop(t *testing.T) {
	list := sll.List[int]{}
	list.Append(1)
	list.Append(2)
	val, err := list.Pop()
	if err != nil || val != 2 || list.Length() != 1 {
		t.Error("Pop failed: expected 2")
	}
	val, err = list.Pop()
	if err != nil || val != 1 || list.Length() != 0 {
		t.Error("Pop failed: expected 1")
	}
	_, err = list.Pop() // Empty list pop
	if err == nil {
		t.Error("Expected error when popping empty list")
	}
}

func TestDeleteFirst(t *testing.T) {
	list := sll.List[int]{}
	list.Append(1)
	list.Append(2)
	val, err := list.DeleteFirst()
	if err != nil || val != 1 || list.Length() != 1 {
		t.Error("DeleteFirst failed: expected 1")
	}
	val, err = list.DeleteFirst()
	if err != nil || val != 2 || list.Length() != 0 {
		t.Error("DeleteFirst failed: expected 2")
	}
	_, err = list.DeleteFirst() // Empty list delete
	if err == nil {
		t.Error("Expected error when deleting first of empty list")
	}
}

func TestDeleteAt(t *testing.T) {
	list := sll.List[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	val, err := list.DeleteAt(2)
	if err != nil || val != 2 || list.Length() != 2 {
		t.Error("DeleteAt failed: expected 2")
	}
	val, err = list.DeleteAt(1)
	if err != nil || val != 1 || list.Length() != 1 {
		t.Error("DeleteAt failed: expected 1")
	}
	_, err = list.DeleteAt(5) // Out of bounds
	if err == nil {
		t.Error("Expected error for out-of-bounds delete")
	}
}

func TestDisplayEmptyList(t *testing.T) {
	list := sll.List[int]{}
	err := list.Display()
	if err == nil {
		t.Error("Expected error when displaying empty list")
	}
}
