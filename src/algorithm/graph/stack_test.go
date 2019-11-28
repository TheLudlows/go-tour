package graph

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {

	stack := new(Stack)
	stack.nodes = make([]Node, 0)
	fmt.Println(stack)

	stack.push(Node{1})
	stack.push(Node{2})
	stack.push(Node{3})
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())

}
