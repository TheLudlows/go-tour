package graph

type Stack struct {
	nodes []Node
}

func (stack *Stack) push(node Node) {
	stack.nodes = append(stack.nodes, node)
}

func (stack *Stack) pop() *Node {
	l := len(stack.nodes)
	node := stack.nodes[l-1]
	stack.nodes = stack.nodes[:l-1]
	return &node
}

func NewStack() Stack {
	stack := Stack{}
	stack.nodes = []Node{}
	return stack
}
func (stack *Stack) isEmpty() bool {
	return len(stack.nodes) == 0
}
