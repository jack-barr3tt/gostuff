package queue

import (
	"container/heap"
	"reflect"
)

// Priority queue implementation where lower priority values are dequeued first

type Item[T any] struct {
	Value    T
	Priority int
	Index    int
}

type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Has(item T) bool {
	for _, i := range *pq {
		if reflect.DeepEqual(i.Value, item) {
			return true
		}
	}
	return false
}

func (pq *PriorityQueue[T]) update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
