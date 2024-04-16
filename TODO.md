# Things I need to do

<!-- TODO:-->

- [ ] Check to make sure all APIs are consistent
    - [x] All Get methods should return (value T, ok bool)
    - [x] All Remove methods should be consistent
    - [x] Store single characters as runes, instead of strings
        - Except in the case of the Trie. The public API for Trie should work only with strings
- [ ] Remove unnecessary prefixes like list in binarysearchlist
- [ ] Move all type definitions to central package (maybe?) (like tree.BinaryNode, and graph.WeightedAdjacencyList)
- [ ] Create simpler testing framework, like jest (maybe?)
    - This might be hard, because of slice (and nested slice) comparison
- [x] Add links to wikipedia articles describing data structures and algorithms
- [x] Package imports and aliasing should be consistent
