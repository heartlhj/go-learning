package common

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

var size = 0
var stack = new(Node)

func Push(v interface{}) bool {
	if stack == nil {
		stack = &Node{v, nil}
		size = 1
		return true
	}

	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

func Pop(t *Node) (interface{}, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		size = 0
		stack = nil
		return t.Value, true
	}

	stack = stack.Next
	size--
	return t.Value, true
}

func Traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Stack!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}
