package graph

import (
	"fmt"
	"sync"
)

type Node struct {
	value int
}

type Graph struct {
	nodes  []*Node          // 节点集
	edges  map[Node][]*Node // 邻接表表示的无向图
	lock   sync.RWMutex     // 保证线程安全
	matrix [][]int          // 邻接矩阵表示带权重无向图
	max    int              // 最大结点个数
}

// 增加节点
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if len(g.nodes) >= g.max {
		panic(fmt.Sprint("more than max node ", g.max))
	}
	g.nodes = append(g.nodes, n)
}

// 增加边
func (g *Graph) AddEdge(from, to *Node, weight ...int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	// 首次建立图
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
		g.matrix = make([][]int, g.max)
		for i := 0; i < len(g.matrix); i++ {
			g.matrix[i] = make([]int, g.max)
		}
	}
	if len(weight) == 1 {
		x, y := IndexOf(g.nodes, from, to)
		if x < 0 || y < 0 {
			panic(fmt.Sprint("edg error from:{},to:{}", x, y))
		}
		g.matrix[x][y] = weight[0]
	}
	g.edges[*from] = append(g.edges[*from], to)
	g.edges[*to] = append(g.edges[*to], from)
}

// 实现 BFS 遍历
func (g *Graph) BFS(f func(node *Node)) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	// 初始化队列
	q := NewNodeQueue()
	// 取图的第一个节点入队列
	head := g.nodes[0]
	q.Enqueue(*head)
	// 标识节点是否已经被访问过
	visited := make(map[*Node]bool)
	visited[head] = true
	// 遍历所有节点直到队列为空
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		nexts := g.edges[*node]
		// 将所有未访问过的邻接节点入队列
		for _, next := range nexts {
			// 如果节点已被访问过
			if visited[next] {
				continue
			}
			q.Enqueue(*next)
			visited[next] = true
		}
		if f != nil {
			f(node)
		}
	}
}

func (graph *Graph) DFS(f func(node *Node)) {
	graph.lock.Lock()
	defer graph.lock.Unlock()

	stack := NewStack()
	head := graph.nodes[0]
	stack.push(*head)
	visited := make(map[*Node]bool)
	visited[head] = true
	f(head)
	for {
		if stack.isEmpty() {
			break
		}
		node := stack.pop()
		visited[node] = true

		around := graph.edges[*node]
		for _, next := range around {
			if !visited[next] {
				stack.push(*node)
				stack.push(*next)
				visited[next] = true
				if f != nil {
					f(next)
				}
				break
			}
		}

	}
}

//find index of node in Node arr,if not exist return -1
func IndexOf(nodes []*Node, from, to *Node) (f, t int) {

	f = -1
	t = -1
	for index, node := range nodes {
		if node == from {
			f = index
		} else if node == to {
			t = index
		}
	}
	return f, t
}
