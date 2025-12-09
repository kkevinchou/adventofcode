package utils

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}

	return &UnionFind{
		parent: parent,
		size:   size,
	}
}

func (u *UnionFind) Find(a int) int {
	for a != u.parent[a] {
		// half compression
		u.parent[a] = u.parent[u.parent[a]]
		a = u.parent[a]
	}
	return a
}

func (u *UnionFind) Union(a, b int) {
	aSet := u.Find(a)
	bSet := u.Find(b)
	if aSet == bSet {
		return
	}

	// swap, aSet will be the bigger set and bSet will be the smaller set
	if u.size[aSet] < u.size[bSet] {
		aSet, bSet = bSet, aSet
	}

	u.parent[bSet] = aSet
	u.size[aSet] += u.size[bSet]
}

func (u *UnionFind) Size(a int) int {
	root := u.Find(a)
	return u.size[root]
}
