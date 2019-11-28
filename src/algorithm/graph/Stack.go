package graph

type Stack struct {
	nodes []Node
}

func (stack Stack) push(node Node) {
	stack.nodes = append(stack.nodes, node)
}

func (stack Stack) pop() (node *Node) {
	l := len(stack.nodes)
	if l == 0 {
		return node
	}
	node = &stack.nodes[l-1]
	stack.nodes = stack.nodes[:l-1]
	return node
}
