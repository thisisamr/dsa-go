package main

import (
	"fmt"

	"github.com/thisisamr/dsa-go/linkedlists/singlylinkedlist"
)

func main() {
	list := &singlylinkedlist.List[int]{Head: nil, Tail: nil}
	fmt.Print(list)
}
