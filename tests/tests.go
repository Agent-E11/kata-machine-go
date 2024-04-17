package tests

import (
	"github.com/agent-e11/kata-machine-go"
)

// -------- Graphs --------

//	 3-(1) -1- (4) -2- (5)
//	/   |       5      /|
//
// (0)   4 ------|--18-- 1
//
//	\   |/      |       |
//	 1-(2) -7- (3) -1- (6)
var GraphList1 kata.WeightedAdjacencyList = [][]kata.GraphEdge{
	{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	},
	{
		{To: 0, Weight: 3},
		{To: 2, Weight: 4},
		{To: 4, Weight: 1},
	},
	{
		{To: 1, Weight: 4},
		{To: 3, Weight: 7},
		{To: 0, Weight: 1},
		{To: 5, Weight: 18},
	},
	{
		{To: 2, Weight: 7},
		{To: 4, Weight: 5},
		{To: 6, Weight: 1},
	},
	{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	},
	{
		{To: 6, Weight: 1},
		{To: 4, Weight: 2},
		{To: 2, Weight: 18},
	},
	{
		{To: 3, Weight: 1},
		{To: 5, Weight: 1},
	},
}

//	 3->(1)<1-1>(4)---2>(5)
//	/      /-----|------/|
//
// (0)   18      5        1
//
//	\    v       v        v
//	 1->(2)---7>(3)<1---(6)
var GraphList2 kata.WeightedAdjacencyList = [][]kata.GraphEdge{
	{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	},
	{
		{To: 4, Weight: 1},
	},
	{
		{To: 3, Weight: 7},
	},
	{},
	{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	},
	{
		{To: 2, Weight: 18},
		{To: 6, Weight: 1},
	},
	{
		{To: 3, Weight: 1},
	},
}

//	 3->(1)<1-1>(4)---2>(5)
//	/      /-----|------/|
//
// (0)   18      5       1
//
//	\    v       v       v
//	 1->(2)---7>(3)<1---(6)
var GraphMatrix2 kata.WeightedAdjacencyMatrix = [][]int{
	{0, 3, 1, 0, 0, 0, 0}, // Node 0
	{0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 7, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 5, 0, 2, 0},
	{0, 0, 18, 0, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 0},
}

// -------- Trees --------

// This tree is a Binary Search Tree
//
// [Binary Search Tree]
//
// [Binary Search Tree]: https://en.wikipedia.org/wiki/Binary_search_tree
var Tree1 kata.BinaryNode[int] = kata.BinaryNode[int]{
	Value: 20,
	Right: &kata.BinaryNode[int]{
		Value: 50,
		Right: &kata.BinaryNode[int]{
			Value: 100,
			Right: nil,
			Left:  nil,
		},
		Left: &kata.BinaryNode[int]{
			Value: 30,
			Right: &kata.BinaryNode[int]{
				Value: 45,
				Right: nil,
				Left:  nil,
			},
			Left: &kata.BinaryNode[int]{
				Value: 29,
				Right: nil,
				Left:  nil,
			},
		},
	},
	Left: &kata.BinaryNode[int]{
		Value: 10,
		Right: &kata.BinaryNode[int]{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &kata.BinaryNode[int]{
			Value: 5,
			Right: &kata.BinaryNode[int]{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}

var Tree2 kata.BinaryNode[int] = kata.BinaryNode[int]{
	Value: 20,
	Right: &kata.BinaryNode[int]{
		Value: 50,
		Right: nil,
		Left: &kata.BinaryNode[int]{
			Value: 30,
			Right: &kata.BinaryNode[int]{
				Value: 45,
				Right: &kata.BinaryNode[int]{
					Value: 49,
					Left:  nil,
					Right: nil,
				},
				Left: nil,
			},
			Left: &kata.BinaryNode[int]{
				Value: 29,
				Right: nil,
				Left: &kata.BinaryNode[int]{
					Value: 21,
					Right: nil,
					Left:  nil,
				},
			},
		},
	},
	Left: &kata.BinaryNode[int]{
		Value: 10,
		Right: &kata.BinaryNode[int]{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &kata.BinaryNode[int]{
			Value: 5,
			Right: &kata.BinaryNode[int]{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}
