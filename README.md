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

Make sure you have [Node.js](https://nodejs.org/en/) and yarn installed: `npm install --global yarn`

clone the repo and install the dependencies

```bash
yarn install
```

edit the `ligma.config.js` file
```javascript
module.exports = {
    dsa: [
        "InsertionSort",
        "MergeSort",
        "Queue",
        "Stack",
        "QuickSort",
        "DijkstraList",
        "PrimsList",
    ],
}
```

create a day of katas, this will use the list in the `ligma.config.js`.
```bash
yarn generate
```

this will progressively create folders named

```
src/day1
src/day2
...
```

`yarn generate` will also update the `tsconfig.json` and `jest.config` to point
the latest `day` folder via tspaths.  This allows us to avoid updating anything
for testing each day.

### Testing
```
yarn test
```

I have yet to create a testing strategy for next sets of algorithms, but we
will get there when i cross that bridge.

## Help wanted
A simple way to specify test, thinking something like `tests.json` and `cat
test.json 2> /dev/null` to specify the tests to run.  tests.json wouldn't be
committed.
