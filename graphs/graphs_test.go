package graphs

import (
	"fmt"
	"testing"

	stringstuff "github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/test"
)

func TestNewVirtualGraph(t *testing.T) {
	origin := "a"

	g := NewVirtualGraph(func(n *Node) []Edge {
		return []Edge{
			{Node: n.Name + "a", Cost: 1},
			{Node: n.Name + "b", Cost: 2},
		}
	}, origin)

	test.AssertEqual(t, g.nodeIds[origin].Name, origin)
}

func TestNewGraph(t *testing.T) {
	// Test valid graph creation
	g, err := NewGraph([]string{"a", "b", "c"}, map[string][]Edge{
		"a": {{Node: "b", Cost: 1}, {Node: "c", Cost: 2}},
		"b": {{Node: "a", Cost: 1}, {Node: "c", Cost: 2}},
		"c": {{Node: "a", Cost: 2}, {Node: "b", Cost: 2}},
	})

	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, g.nodeIds["a"].Name, "a")
	test.AssertEqual(t, g.nodeIds["b"].Name, "b")
	test.AssertEqual(t, g.nodeIds["c"].Name, "c")

	// Test error case: edges reference non-existent nodes
	_, err = NewGraph([]string{"a", "b"}, map[string][]Edge{
		"a": {{Node: "c", Cost: 1}},
	})
	test.AssertEqual(t, err != nil, true)

	// Test error case: edges defined for non-existent node
	_, err = NewGraph([]string{"a", "b"}, map[string][]Edge{
		"a": {{Node: "b", Cost: 1}},
		"c": {{Node: "a", Cost: 1}},
	})
	test.AssertEqual(t, err != nil, true)
}

func TestNewEmptyGraph(t *testing.T) {
	g := NewEmptyGraph()
	test.AssertEqual(t, len(g.nodeIds), 0)
}

func TestAddNode(t *testing.T) {
	g, _ := NewGraph([]string{"a", "b"}, map[string][]Edge{
		"a": {{Node: "b", Cost: 1}},
		"b": {{Node: "a", Cost: 1}},
	})

	g.AddNode("c", []Edge{{Node: "a", Cost: 2}})

	nodeC, ok := g.At("c")
	test.AssertEqual(t, ok, true)
	test.AssertEqual(t, nodeC.Name, "c")
	test.AssertEqual(t, len(nodeC.Adj), 1)
	test.AssertEqual(t, nodeC.Adj[0].Node, "a")
	test.AssertEqual(t, nodeC.Adj[0].Cost, 2)
}

func TestAddEdge(t *testing.T) {
	g, _ := NewGraph([]string{"a", "b"}, map[string][]Edge{
		"a": {},
		"b": {},
	})

	err := g.AddEdge("a", "b", 3)
	test.AssertEqual(t, err, nil)

	nodeA, _ := g.At("a")
	test.AssertEqual(t, len(nodeA.Adj), 1)
	test.AssertEqual(t, nodeA.Adj[0].Node, "b")
	test.AssertEqual(t, nodeA.Adj[0].Cost, 3)

	// Test error case: adding edge between non-existent nodes
	err = g.AddEdge("a", "c", 2)
	test.AssertEqual(t, err != nil, true)
}

func TestAt(t *testing.T) {
	origin := "a"

	g := NewVirtualGraph(func(n *Node) []Edge {
		return []Edge{
			{Node: n.Name + "a", Cost: 1},
			{Node: n.Name + "b", Cost: 2},
		}
	}, origin)

	node, ok := g.At(origin)
	test.AssertEqual(t, node.Name, origin)
	test.AssertEqual(t, ok, true)

	node, ok = g.At(origin + "a")
	test.AssertEqual(t, node.Name, origin+"a")
	test.AssertEqual(t, ok, true)

	node, ok = g.At(origin + "b")
	test.AssertEqual(t, node.Name, origin+"b")
	test.AssertEqual(t, ok, true)

	g, _ = NewGraph([]string{"a", "b", "c"}, map[string][]Edge{
		"a": {{Node: "b", Cost: 1}, {Node: "c", Cost: 2}},
		"b": {{Node: "a", Cost: 1}, {Node: "c", Cost: 2}},
		"c": {{Node: "a", Cost: 2}, {Node: "b", Cost: 2}},
	})

	a, ok := g.At("a")
	test.AssertEqual(t, ok, true)
	test.AssertEqual(t, a.Name, "a")
	test.AssertEqual(t, len(a.Adj), 2)
}

func TestShortestPath(t *testing.T) {
	// test case inspired by advent of code 2024 day 13
	g1 := NewVirtualGraph(func(n *Node) []Edge {
		pos := stringstuff.GetNums(n.Name)

		return []Edge{
			{Node: fmt.Sprintf("%d,%d", pos[0]+94, pos[1]+34), Cost: 3},
			{Node: fmt.Sprintf("%d,%d", pos[0]+22, pos[1]+67), Cost: 1},
		}
	}, "0,0")

	_, length1 := g1.ShortestPath("0,0", "8400,5400", func(n Node) int {
		pos := stringstuff.GetNums(n.Name)
		if pos[0] > 8400 || pos[1] > 5400 {
			return -1
		}
		return 8400 - pos[0] + 5400 - pos[1]
	})

	test.AssertEqual(t, length1, 280)

	g2 := NewVirtualGraph(func(n *Node) []Edge {
		pos := stringstuff.GetNums(n.Name)

		return []Edge{
			{Node: fmt.Sprintf("%d,%d", pos[0]+26, pos[1]+66), Cost: 3},
			{Node: fmt.Sprintf("%d,%d", pos[0]+67, pos[1]+21), Cost: 1},
		}
	}, "0,0")

	_, length2 := g2.ShortestPath("0,0", "12748,12176", func(n Node) int {
		pos := stringstuff.GetNums(n.Name)
		if pos[0] > 12748 || pos[1] > 12176 {
			return -1
		}
		return 12748 - pos[0] + 12176 - pos[1]
	})

	test.AssertEqual(t, length2, -1)

	// test case inspired by pearson edexcel a level decision mathematics 1 textbook ISBN 9781292183299 page 66

	g3, _ := NewGraph([]string{"S", "A", "B", "C", "D", "T"}, map[string][]Edge{
		"S": {{Node: "A", Cost: 5}, {Node: "B", Cost: 6}, {Node: "C", Cost: 2}},
		"A": {{Node: "S", Cost: 5}, {Node: "D", Cost: 4}},
		"B": {{Node: "S", Cost: 6}, {Node: "D", Cost: 4}, {Node: "T", Cost: 8}, {Node: "C", Cost: 2}},
		"C": {{Node: "S", Cost: 2}, {Node: "B", Cost: 2}, {Node: "T", Cost: 12}},
		"D": {{Node: "A", Cost: 4}, {Node: "B", Cost: 4}, {Node: "T", Cost: 3}},
		"T": {{Node: "B", Cost: 8}, {Node: "C", Cost: 12}, {Node: "D", Cost: 3}},
	})

	path, length := g3.ShortestPath("S", "T", func(n Node) int {
		return 1
	})

	test.AssertEqual(t, path, []string{"S", "C", "B", "D", "T"})
	test.AssertEqual(t, length, 11)
}

func TestAllShortestPaths(t *testing.T) {
	g, _ := NewGraph([]string{"A", "B", "C", "D", "E"}, map[string][]Edge{
		"A": {{Node: "B", Cost: 1}, {Node: "C", Cost: 2}, {Node: "D", Cost: 1}},
		"B": {{Node: "A", Cost: 1}, {Node: "E", Cost: 2}},
		"C": {{Node: "A", Cost: 2}, {Node: "D", Cost: 1}, {Node: "E", Cost: 1}},
		"D": {{Node: "A", Cost: 1}, {Node: "C", Cost: 1}, {Node: "E", Cost: 3}},
		"E": {{Node: "B", Cost: 2}, {Node: "C", Cost: 1}, {Node: "D", Cost: 3}},
	})

	paths, length := g.AllShortestPaths("A", "E", func(n Node) int { return 1 })

	expected := [][]string{{"A", "B", "E"}, {"A", "C", "E"}, {"A", "D", "C", "E"}}

	test.AssertEqual(t, length, 3)
	test.AssertSlicesEqual(t, paths, expected)

	// test getting all shortest paths while using a heuristic function that makes a node unreachable

	pathsH, lengthH := g.AllShortestPaths("A", "E", func(n Node) int {
		if n.Name == "B" {
			return -1
		}
		return 1
	})

	expectedH := [][]string{{"A", "C", "E"}, {"A", "D", "C", "E"}}

	test.AssertEqual(t, lengthH, 3)
	test.AssertSlicesEqual(t, pathsH, expectedH)
}

func TestDFT(t *testing.T) {
	g, _ := NewGraph([]string{"A", "B", "C", "D", "E"}, map[string][]Edge{
		"A": {{Node: "B", Cost: 1}},
		"B": {{Node: "C", Cost: 1}, {Node: "A", Cost: 1}},
		"C": {{Node: "B", Cost: 1}},
		"D": {{Node: "E", Cost: 1}},
		"E": {{Node: "D", Cost: 1}},
	})

	visitedA := []string{}
	g.DFT("A", func(n Node) {
		visitedA = append(visitedA, n.Name)
	})

	test.AssertSlicesEqual(t, visitedA, []string{"A", "B", "C"})

	visitedD := []string{}
	g.DFT("D", func(n Node) {
		visitedD = append(visitedD, n.Name)
	})

	test.AssertSlicesEqual(t, visitedD, []string{"D", "E"})
}

func TestConnected(t *testing.T) {
	g, _ := NewGraph([]string{"A", "B", "C", "D", "E"}, map[string][]Edge{
		"A": {{Node: "B", Cost: 1}},
		"B": {{Node: "C", Cost: 1}, {Node: "A", Cost: 1}},
		"C": {{Node: "B", Cost: 1}},
		"D": {{Node: "E", Cost: 1}},
		"E": {{Node: "D", Cost: 1}},
	})

	test.AssertSlicesEqual(t, g.Connected("A"), []string{"A", "B", "C"})
	test.AssertSlicesEqual(t, g.Connected("B"), []string{"A", "B", "C"})
	test.AssertSlicesEqual(t, g.Connected("C"), []string{"A", "B", "C"})
	test.AssertSlicesEqual(t, g.Connected("D"), []string{"D", "E"})
	test.AssertSlicesEqual(t, g.Connected("E"), []string{"D", "E"})
}

func TestSubgraphs(t *testing.T) {
	g, _ := NewGraph([]string{"A", "B", "C", "D", "E"}, map[string][]Edge{
		"A": {{Node: "B", Cost: 1}},
		"B": {{Node: "C", Cost: 1}, {Node: "A", Cost: 1}},
		"C": {{Node: "B", Cost: 1}},
		"D": {{Node: "E", Cost: 1}},
		"E": {{Node: "D", Cost: 1}},
	})

	subgraphs := g.Subgraphs()
	expected := []string{"A", "D"}

	test.AssertSlicesEqual(t, subgraphs, expected)
}
