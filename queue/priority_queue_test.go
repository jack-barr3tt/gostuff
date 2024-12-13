package queue

import (
	"container/heap"
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestPriorityQueue(t *testing.T) {
	pq := make(PriorityQueue[string], 0)
	heap.Init(&pq)

	item1 := &Item[string]{Value: "foo", Priority: 3}
	item2 := &Item[string]{Value: "bar", Priority: 2}
	item3 := &Item[string]{Value: "baz", Priority: 1}

	heap.Push(&pq, item1)
	heap.Push(&pq, item2)
	heap.Push(&pq, item3)

	test.AssertEqual(t, "baz", pq[0].Value)
	test.AssertEqual(t, 1, pq[0].Priority)

	item := heap.Pop(&pq).(*Item[string])
	test.AssertEqual(t, "baz", item.Value)
	test.AssertEqual(t, 1, item.Priority)

	item = heap.Pop(&pq).(*Item[string])
	test.AssertEqual(t, "bar", item.Value)
	test.AssertEqual(t, 2, item.Priority)

	item = heap.Pop(&pq).(*Item[string])
	test.AssertEqual(t, "foo", item.Value)
	test.AssertEqual(t, 3, item.Priority)
}

func TestPriorityQueueUpdate(t *testing.T) {
	pq := make(PriorityQueue[string], 0)
	heap.Init(&pq)

	item1 := &Item[string]{Value: "baz", Priority: 5}
	item2 := &Item[string]{Value: "bar", Priority: 4}
	item3 := &Item[string]{Value: "foo", Priority: 3}

	heap.Push(&pq, item1)
	heap.Push(&pq, item2)
	heap.Push(&pq, item3)

	test.AssertEqual(t, "foo", pq[0].Value)
	test.AssertEqual(t, 3, pq[0].Priority)

	pq.update(item1, "foo-updated", 1)
	
	test.AssertEqual(t, "foo-updated", pq[0].Value)
	test.AssertEqual(t, 1, pq[0].Priority)
}
