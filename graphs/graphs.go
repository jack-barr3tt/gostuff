package graphs

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/jack-barr3tt/gostuff/maps"
	"github.com/jack-barr3tt/gostuff/queue"
	"github.com/jack-barr3tt/gostuff/slices"
)

type Edge struct {
	Node string
	Cost int
}

type Node struct {
	Name string
	Adj  []Edge
}

type Graph struct {
	nodeIds map[string]*Node
	gen     func(n *Node) []Edge
}

func NewVirtualGraph(nodeGenerator func(n *Node) []Edge, origin string) Graph {
	node := &Node{Name: origin}

	nodeIds := make(map[string]*Node)
	nodeIds[origin] = node

	return Graph{nodeIds: nodeIds, gen: nodeGenerator}
}

func NewGraph(nodes []string, edges map[string][]Edge) (Graph, error) {
	nodeIds := make(map[string]*Node)
	for _, name := range nodes {
		nodeIds[name] = &Node{Name: name, Adj: edges[name]}
	}

	for nodeName, nodeEdges := range edges {
		if _, exists := nodeIds[nodeName]; !exists {
			return Graph{}, fmt.Errorf("edges defined for non-existent node: %s", nodeName)
		}
		for _, edge := range nodeEdges {
			if _, exists := nodeIds[edge.Node]; !exists {
				return Graph{}, fmt.Errorf("edge from %s references non-existent node: %s", nodeName, edge.Node)
			}
		}
	}

	return Graph{nodeIds: nodeIds}, nil
}

func (g Graph) At(name string) (*Node, bool) {
	n, ok := g.nodeIds[name]
	if g.gen == nil {
		return n, ok
	}
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

// ShortestPath returns the shortest path from source to target using the A* algorithm.
// The heuristic function should return -1 if the node is not reachable.
// The heuristic function should return lower values for nodes that are more favorable.
func (g Graph) ShortestPath(source, target string, heuristic func(n Node) int) ([]string, int) {
	if _, ok := g.At(source); !ok {
		return nil, -1
	}

	pq := make(queue.PriorityQueue[string], 0)
	heap.Init(&pq)
	heap.Push(&pq, &queue.Item[string]{Value: source, Priority: 0})

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

		currNode, _ := g.At(curr)
		for _, edge := range currNode.Adj {
			newCost := costSoFar[curr] + edge.Cost
			if _, ok := costSoFar[edge.Node]; (!ok || newCost < costSoFar[edge.Node]) && heuristic(*g.nodeIds[edge.Node]) != -1 {
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

func (g Graph) reconstructPath(cameFrom map[string]string, current string) ([]string, int) {
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

// AllShortestPaths returns all shortest paths from start to goal.
func (g Graph) AllShortestPaths(source, target string, heuristic func(n Node) int) ([][]string, int) {
	if _, ok := g.At(source); !ok {
		return nil, -1
	}

	pq := make(queue.PriorityQueue[string], 0)
	heap.Init(&pq)
	heap.Push(&pq, &queue.Item[string]{Value: source, Priority: 0})

	cameFrom := make(map[string][]string)
	costSoFar := make(map[string]int)
	costSoFar[source] = 0

	var allPaths [][]string
	minCost := -1

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*queue.Item[string]).Value
		if curr == target {
			if minCost == -1 {
				minCost = costSoFar[curr]
			}
			if costSoFar[curr] == minCost {
				allPaths = append(allPaths, g.reconstructAllPaths(cameFrom, target)...)
			}
			continue
		}

		currNode, _ := g.At(curr)
		for _, edge := range currNode.Adj {
			newCost := costSoFar[curr] + edge.Cost
			if _, ok := costSoFar[edge.Node]; (!ok || newCost < costSoFar[edge.Node]) && heuristic(*g.nodeIds[edge.Node]) != -1 {
				cameFrom[edge.Node] = []string{curr}
				costSoFar[edge.Node] = newCost
				priority := newCost + heuristic(*g.nodeIds[edge.Node])
				heap.Push(&pq, &queue.Item[string]{Value: edge.Node, Priority: priority})
			} else if newCost == costSoFar[edge.Node] {
				cameFrom[edge.Node] = append(cameFrom[edge.Node], curr)
			}
		}
	}

	allUniquePaths := make(map[string]bool)
	for _, path := range allPaths {
		allUniquePaths[strings.Join(path, "_")] = true
	}

	return slices.Map(func(v string) []string { return strings.Split(v, "_") }, maps.Keys(allUniquePaths)), minCost
}

func (g Graph) reconstructAllPaths(cameFrom map[string][]string, current string) [][]string {
	var paths [][]string
	var dfs func(path []string, node string)
	dfs = func(path []string, node string) {
		if len(cameFrom[node]) == 0 {
			paths = append(paths, append([]string{node}, path...))
			return
		}
		for _, prev := range cameFrom[node] {
			dfs(append([]string{node}, path...), prev)
		}
	}
	dfs([]string{}, current)

	return paths
}

func (g Graph) DFT(start string, visit func(n Node)) {
	visited := make(map[string]bool)
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if visited[n.Name] {
			return
		}
		visited[n.Name] = true
		visit(*n)
		for _, edge := range n.Adj {
			nextNode, _ := g.At(edge.Node)
			dfs(nextNode)
		}
	}
	startNode, ok := g.At(start)
	if !ok {
		return
	}
	dfs(startNode)
}

