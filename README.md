# Kata Machine (Go)

## What is this?

This repo is a version of
[ThePrimeagen's kata-machine](https://github.com/ThePrimeagen/kata-machine)
repo ported to Go. You can use it to go through his
[Data Structures & Algorithms course](https://frontendmasters.com/courses/algorithms/introduction)
using Go instead of TypeScript. There are also a few other niceties, like
Wikipedia links for the descriptions of different data structures and
algorithms.

## Supported Algorithms
- Insertion sort
- Merge sort
- Quick sort
- Bubble sort
- Prim's Minimum Spanning Tree
    - Adjacency List
- Dijkstra's Shortest Path
    - Adjacency List
- Linear search
- Binary search
- Binary Tree Traversal
    - Pre-Order
    - In-Order
    - Post-Order
- Depth-first search on Binary Search Tree
- Depth-first search on graph
    - Adjacency List
- Breadth-first search on graph
    - Adjacency List
    - Adjacency Matrix
- Maze solver
- Two Crystal Balls problem

## Supported Data Structures
- Singly linked list
- Doubly linked list
- Array list
- Ring buffer
- Queue
- Stack
- Graph
    - Adjacency List
    - Adjacency Matrix
- Binary tree
- Min-Heap
- Trie
- Hashmap
- Least Recently Used cache

## How It Works

Make sure you have Go and Python installed:

```sh
go version
python3 -V
```

Clone the repo:

```sh
git clone https://github.com/Agent-E11/kata-machine-go
```

Create a day of katas:

```sh
python3 ./scripts/generate.py
```

This will progressively create folders named like this:

```
./day0
./day1
...
```

You can now edit the Go files in the (for example) `./day0/` folder, and follow
along with the course.

### Testing

To test all of your code (e.g. the day0 folder), run:

```sh
go test -v ./day0/...
```

Or, if you want to test a specific algorithm/data structure (e.g. arraylist):

```sh
go test -v ./day0/arraylist
```

