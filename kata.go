package kata

// TODO: I think that the BinaryNode should be defined in binarytree

// [Binary Tree]
//
// [Binary Tree]: https://en.wikipedia.org/wiki/Binary_tree
type BinaryNode[T comparable] struct {
	Value T
	Left *BinaryNode[T]
	Right *BinaryNode[T]
}

// [Adjacency Matrix]
//
// [Adjacency Matrix]: https://en.wikipedia.org/wiki/Adjacency_matrix
type WeightedAdjacencyMatrix = [][]int

type GraphEdge struct {
	To int
	Weight int
}

// [Adjacency List]
//
// [Adjacency List]: https://en.wikipedia.org/wiki/Adjacency_list
type WeightedAdjacencyList = [][]GraphEdge

type Point struct {
	X int
	Y int
}

