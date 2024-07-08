# Things I need to do

<!-- TODO:-->

- [ ] Rework the README file
- [ ] Remove unnecessary JavaScript/TypeScript files
- [ ] Add missing algorithms (If there is an algorithm on an Adjacency List, there should be one for an Adjacency Matrix)
- [x] Check to make sure all APIs are consistent
    - [x] All Get methods should return (value T, ok bool)
    - [x] All Remove methods should be consistent
    - [x] Store single characters as runes, instead of strings
        - Except in the case of the Trie. The public API for Trie should work only with strings
- [x] Create script to generate a day of katas
- [x] Move all type definitions to central package (maybe?) (like tree.BinaryNode, and graph.WeightedAdjacencyList)
- [x] Remove unnecessary suffixes like list in binarysearchlist
- [x] Package imports and aliasing should be consistent
- [x] Add links to Wikipedia articles describing data structures and algorithms

# Unplanned
- [ ] Create simpler testing framework, like jest (maybe?)
    - This might be hard, because of slice (and nested slice) comparison
