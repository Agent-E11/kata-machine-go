# Things I need to do

<!-- TODO:-->

- [ ] Check to make sure all APIs are consistent
    - [ ] All Get methods should return (value T, ok bool)
    - [ ] All Remove methods should be consistent
    - [x] Store single characters as runes, instead of strings
        - Except in the case of the Trie, the public API should work only with strings
- [ ] Move all type definitions to central package (maybe?) (like tree.BinaryNode, and graph.WeightedAdjacencyList)
- [ ] Create simpler testing framework, like jest (maybe?)
- [x] Add links to wikipedia articles describing data structures and algorithms
