package graph

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	g := new(Graph)
	a := new(Node)
	b := new(Node)
	c := new(Node)

	a.value = 1
	b.value = 2
	c.value = 3

	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)

	g.AddEdge(a, b)
	g.AddEdge(b, c)
	g.AddEdge(c, a)

	g.BFS(func(node *Node) {
		fmt.Printf("[Current Traverse Node]: %v\n", node)
	})

}

type Value struct {
	value int
}

func TestPointer(t *testing.T) {
	fmt.Println(new(Value))
	fmt.Println(*new(Value))

	mp := make(map[string]*Value)
	str := new(Value)
	str.value = 100
	mp["key"] = str
	fmt.Println(mp)
}

func TestPointer2(t *testing.T) {
	v := Value{value: 100}
	fmt.Println(v)
	v1 := change1(v)
	fmt.Println(v1)
	fmt.Println(v)

	v2 := change2(&v)
	fmt.Println(v2)
	fmt.Println(v)

}

func change1(v Value) Value {
	v.value = 200
	return v
}

func change2(v *Value) *Value {
	v.value = 300
	return v
}
