package tree

import "github.com/agent-e11/kata-machine-go"

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
            Left: nil,
        },
        Left: &kata.BinaryNode[int]{
            Value: 30,
            Right: &kata.BinaryNode[int]{
                Value: 45,
                Right: nil,
                Left: nil,
            },
            Left: &kata.BinaryNode[int]{
                Value: 29,
                Right: nil,
                Left: nil,
            },
        },
    },
    Left: &kata.BinaryNode[int]{
        Value: 10,
        Right: &kata.BinaryNode[int]{
            Value: 15,
            Right: nil,
            Left: nil,
        },
        Left: &kata.BinaryNode[int]{
            Value: 5,
            Right: &kata.BinaryNode[int]{
                Value: 7,
                Right: nil,
                Left: nil,
            },
            Left: nil,
        },
    },
};

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
                    Left: nil,
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
                    Left: nil,
                },
            },
        },
    },
    Left: &kata.BinaryNode[int]{
        Value: 10,
        Right: &kata.BinaryNode[int]{
            Value: 15,
            Right: nil,
            Left: nil,
        },
        Left: &kata.BinaryNode[int]{
            Value: 5,
            Right: &kata.BinaryNode[int]{
                Value: 7,
                Right: nil,
                Left: nil,
            },
            Left: nil,
        },
    },
}
