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

func TestVirtualAt(t *testing.T) {
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
}

func TestVirtualShortestPath(t *testing.T) {
	// test case inspired by advent of code 2024 day 13
	g := NewVirtualGraph(func(n *Node) []Edge {
		pos := stringstuff.GetNums(n.Name)

		return []Edge{
			{Node: fmt.Sprintf("%d,%d", pos[0]+94, pos[1]+34), Cost: 3},
			{Node: fmt.Sprintf("%d,%d", pos[0]+22, pos[1]+67), Cost: 1},
		}
	}, "0,0")

	_, length := g.ShortestPath("0,0", "8400,5400", func(n Node) int {
		pos := stringstuff.GetNums(n.Name)
		if pos[0] > 8400 || pos[1] > 5400 {
			return 1 << 31
		}
		return 8400 - pos[0] + 5400 - pos[1]
	})

	test.AssertEqual(t, length, 280)
}
