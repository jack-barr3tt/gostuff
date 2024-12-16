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
	g := NewGraph([]string{"a", "b", "c"}, map[string][]Edge{
		"a": {{Node: "b", Cost: 1}, {Node: "c", Cost: 2}},
		"b": {{Node: "a", Cost: 1}, {Node: "c", Cost: 2}},
		"c": {{Node: "a", Cost: 2}, {Node: "b", Cost: 2}},
	})

	test.AssertEqual(t, g.nodeIds["a"].Name, "a")
	test.AssertEqual(t, g.nodeIds["b"].Name, "b")
	test.AssertEqual(t, g.nodeIds["c"].Name, "c")
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

	g = NewGraph([]string{"a", "b", "c"}, map[string][]Edge{
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
}
