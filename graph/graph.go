package graph

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

//    3-(1) -1- (4) -2- (5)
//   /   |       5      /|
// (0)   4 ------|--18-- 1
//   \   |/      |       |
//    1-(2) -7- (3) -1- (6)
var List1 WeightedAdjacencyList = [][]GraphEdge{
	{
		{ To: 1, Weight: 3 },
		{ To: 2, Weight: 1 },
	},
	{
		{ To: 0, Weight: 3 },
		{ To: 2, Weight: 4 },
		{ To: 4, Weight: 1 },
	},
	{
		{ To: 1, Weight: 4 },
		{ To: 3, Weight: 7 },
		{ To: 0, Weight: 1 }, // FIXME: Node 2 should have a connection to node 5 with a weight of 18
	},
	{
		{ To: 2, Weight: 7 },
		{ To: 4, Weight: 5 },
		{ To: 6, Weight: 1 },
	},
	{
		{ To: 1, Weight: 1 },
		{ To: 3, Weight: 5 },
		{ To: 5, Weight: 2 },
	},
	{
		{ To: 6, Weight: 1 },
		{ To: 4, Weight: 2 },
		{ To: 2, Weight: 18 },
	},
	{
		{ To: 3, Weight: 1 },
		{ To: 5, Weight: 1 },
	},
}

//   3->(1)<1-1>(4)---2>(5)
//  /      /-----|------/|
// (0)   18      5        1
//  \    v       v        v
//   1->(2)---7>(3)<1---(6)
var List2 WeightedAdjacencyList = [][]GraphEdge{
	{
		{ To: 1, Weight: 3 },
		{ To: 2, Weight: 1 },
	},
	{
		{ To: 4, Weight: 1 },
	},
	{
		{ To: 3, Weight: 7 },
	},
	{ },
	{
		{ To: 1, Weight: 1 },
		{ To: 3, Weight: 5 },
		{ To: 5, Weight: 2 },
	},
	{
		{ To: 2, Weight: 18 },
		{ To: 6, Weight: 1 },
	},
	{
		{ To: 3, Weight: 1 },
	},
}

//   3->(1)<1-1>(4)---2>(5)
//  /      /-----|------/|
// (0)   18      5        1
//  \    v       v        v
//   1->(2)---7>(3)<1---(6)
var Matrix2 WeightedAdjacencyMatrix = [][]int{
    {0, 3, 1,  0, 0, 0, 0}, // Node 0
    {0, 0, 0,  0, 1, 0, 0},
	{0, 0, 7,  0, 0, 0, 0}, // FIXME: I think that the 7 should be one over (index 3)
    {0, 0, 0,  0, 0, 0, 0},
    {0, 1, 0,  5, 0, 2, 0},
    {0, 0, 18, 0, 0, 0, 1},
	{0, 0, 0,  1, 0, 0, 1}, // FIXME: This node shouldn't have a connection to itself
}
