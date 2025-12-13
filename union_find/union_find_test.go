package union_find

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestNew(t *testing.T) {
	uf := New[int]()
	if uf == nil {
		t.Fatal("New() returned nil")
	}
	test.AssertEqual(t, uf.Size(), 0)
	test.AssertEqual(t, uf.Count(), 0)
}

func TestMakeSet(t *testing.T) {
	uf := New[string]()

	uf.MakeSet("a")
	test.AssertEqual(t, uf.Size(), 1)
	test.AssertEqual(t, uf.Count(), 1)

	uf.MakeSet("a")
	test.AssertEqual(t, uf.Size(), 1)

	uf.MakeSet("b")
	test.AssertEqual(t, uf.Size(), 2)
	test.AssertEqual(t, uf.Count(), 2)
}

func TestFind(t *testing.T) {
	uf := New[int]()

	// Find on non-existent element should create it
	root := uf.Find(1)
	test.AssertEqual(t, root, 1)
	test.AssertEqual(t, uf.Size(), 1)

	// Find again should return the same root
	root = uf.Find(1)
	test.AssertEqual(t, root, 1)
}

func TestUnion(t *testing.T) {
	uf := New[int]()

	merged := uf.Union(1, 2)
	test.AssertEqual(t, merged, true)
	test.AssertEqual(t, uf.Size(), 2)
	test.AssertEqual(t, uf.Count(), 1)

	// Union again should return false
	merged = uf.Union(1, 2)
	test.AssertEqual(t, merged, false)

	// Add a third element and union it
	merged = uf.Union(2, 3)
	test.AssertEqual(t, merged, true)
	test.AssertEqual(t, uf.Size(), 3)
	test.AssertEqual(t, uf.Count(), 1)

	// All three should be in the same set
	test.AssertEqual(t, uf.Connected(1, 2), true)
	test.AssertEqual(t, uf.Connected(2, 3), true)
	test.AssertEqual(t, uf.Connected(1, 3), true)
}

func TestConnected(t *testing.T) {
	uf := New[string]()

	uf.MakeSet("a")
	uf.MakeSet("b")
	uf.MakeSet("c")

	// Initially, no elements are connected
	test.AssertEqual(t, uf.Connected("a", "b"), false)
	test.AssertEqual(t, uf.Connected("b", "c"), false)

	// Connect a and b
	uf.Union("a", "b")
	test.AssertEqual(t, uf.Connected("a", "b"), true)
	test.AssertEqual(t, uf.Connected("b", "c"), false)

	// Connect b and c (which connects all three)
	uf.Union("b", "c")
	test.AssertEqual(t, uf.Connected("a", "b"), true)
	test.AssertEqual(t, uf.Connected("b", "c"), true)
	test.AssertEqual(t, uf.Connected("a", "c"), true)
}

func TestCount(t *testing.T) {
	uf := New[int]()

	// Add 5 separate elements
	for i := 0; i < 5; i++ {
		uf.MakeSet(i)
	}

	test.AssertEqual(t, uf.Count(), 5)

	uf.Union(0, 1)
	test.AssertEqual(t, uf.Count(), 4)

	uf.Union(2, 3)
	test.AssertEqual(t, uf.Count(), 3)

	uf.Union(0, 2)
	test.AssertEqual(t, uf.Count(), 2)

	// Union all elements into one set
	uf.Union(0, 4)
	test.AssertEqual(t, uf.Count(), 1)
}

func TestGenericWithStruct(t *testing.T) {
	type Point struct {
		X, Y int
	}

	uf := New[Point]()

	p1 := Point{1, 2}
	p2 := Point{3, 4}
	p3 := Point{5, 6}

	uf.MakeSet(p1)
	uf.MakeSet(p2)
	uf.MakeSet(p3)

	test.AssertEqual(t, uf.Size(), 3)

	uf.Union(p1, p2)
	test.AssertEqual(t, uf.Connected(p1, p2), true)
	test.AssertEqual(t, uf.Connected(p1, p3), false)

	uf.Union(p2, p3)
	test.AssertEqual(t, uf.Connected(p1, p3), true)
}
