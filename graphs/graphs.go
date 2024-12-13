package graphs

import (
	"container/heap"

	"github.com/jack-barr3tt/gostuff/queue"
)

type Edge struct {
	Node string
	Cost int
}

type Node struct {
	Name string
	Adj  []Edge
}

type VirtualGraph struct {
	nodeIds map[string]*Node
	gen     func(n *Node) []Edge
}

func NewVirtualGraph(nodeGenerator func(n *Node) []Edge, origin string) VirtualGraph {
	node := &Node{Name: origin}

	nodeIds := make(map[string]*Node)
	nodeIds[origin] = node

	return VirtualGraph{nodeIds: nodeIds, gen: nodeGenerator}
}

func (g VirtualGraph) At(name string) (*Node, bool) {
	n, ok := g.nodeIds[name]
	if ok && len(n.Adj) == 0 {
		n.Adj = g.gen(n)
		for _, edge := range n.Adj {
			if _, ok := g.nodeIds[edge.Node]; !ok {
				g.nodeIds[edge.Node] = &Node{Name: edge.Node}
			}
		}
	}
	return n, ok
}

func (g VirtualGraph) ShortestPath(source, target string, heuristic func(n Node) int) ([]string, int) {
	// check the source node exists
	if _, ok := g.At(source); !ok {
		return nil, -1
	}

	// make the priority queue
	pq := make(queue.PriorityQueue[string], 0)
	heap.Init(&pq)

	// add the source node to the queue
	startItem := &queue.Item[string]{Value: source, Priority: 0}
	heap.Push(&pq, startItem)

	cameFrom := make(map[string]string)
	costSoFar := make(map[string]int)
	costSoFar[source] = 0

	costRemaining := make(map[string]int)
	costRemaining[source] = heuristic(*g.nodeIds[source])

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*queue.Item[string]).Value
		if curr == target {
			return g.reconstructPath(cameFrom, target)
		}

		// assume nodes in the queue are always valid so ignore the ok value
		currNode, _ := g.At(curr)

		for _, edge := range currNode.Adj {
			newCost := costSoFar[curr] + edge.Cost

			if _, ok := costSoFar[edge.Node]; !ok || newCost < costSoFar[edge.Node] {
				cameFrom[edge.Node] = curr
				costSoFar[edge.Node] = newCost
				costRemaining[edge.Node] = newCost + heuristic(*g.nodeIds[edge.Node])

				if !pq.Has(edge.Node) {
					item := &queue.Item[string]{Value: edge.Node, Priority: costRemaining[edge.Node]}
					heap.Push(&pq, item)
				}
			}
		}
	}

	return nil, -1
}

func (g VirtualGraph) reconstructPath(cameFrom map[string]string, current string) ([]string, int) {
	totalPath := []string{current}
	ok := true
	for {
		current, ok = cameFrom[current]
		if !ok {
			break
		}
		totalPath = append([]string{current}, totalPath...)
	}

	cost := 0
	for i := 0; i < len(totalPath)-1; i++ {
		currNode, _ := g.At(totalPath[i])
		for _, edge := range currNode.Adj {
			if edge.Node == totalPath[i+1] {
				cost += edge.Cost
				break
			}
		}
	}

	return totalPath, cost
}
