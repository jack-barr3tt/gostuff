package union_find

type UnionFind[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}

func New[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
	}
}

func (uf *UnionFind[T]) MakeSet(x T) {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x
		uf.rank[x] = 0
	}
}

func (uf *UnionFind[T]) Find(x T) T {
	if _, exists := uf.parent[x]; !exists {
		uf.MakeSet(x)
		return x
	}

	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind[T]) Union(x, y T) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	return true
}

func (uf *UnionFind[T]) Connected(x, y T) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind[T]) Size() int {
	return len(uf.parent)
}

func (uf *UnionFind[T]) Count() int {
	roots := make(map[T]bool)
	for element := range uf.parent {
		roots[uf.Find(element)] = true
	}
	return len(roots)
}
