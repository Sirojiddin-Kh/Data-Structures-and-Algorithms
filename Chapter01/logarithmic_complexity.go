package main 

import (

	"fmt"
)

type Tree struct {

	LeftNode *Tree
	Value int
	RightNode *Tree
}

func (tree *Tree) Insert(m int) {

	if tree != nil {

		if tree.LeftNode == nil {

			tree.LeftNode = &Tree{nil, m, nil}
		} else {

			if tree.RightNode == nil {

				tree.RightNode = &Tree{nil, m, nil}
			} else {

				if tree.LeftNode != nil {

					tree.LeftNode.Insert(m)
				} else {
					tree.RightNode.Insert(m)
				}
			}
		}
	} else {
		tree = &Tree{nil, m, nil}
	}
}

func print(tree *Tree) {

	if tree != nil {

		fmt.Println("Value", tree.Value)
		fmt.Println("Tree Node Left")
		print(tree.LeftNode)
		fmt.Println("Tree Node Right")
		print(tree.RightNode)
	} else {

		fmt.Println("Nil\n")
	}
}

func main() {

	var tree *Tree = &Tree{nil, 1, nil}
	print(tree)
	tree.Insert(3)
	print(tree)
	tree.Insert(5)
	print(tree)
	tree.Insert(7)
	print(tree)
}