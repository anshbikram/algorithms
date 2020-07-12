package main

import (
	"fmt"
	"algorithms/tree"
)

func main() {  
	t := tree.BST{}
	t.Insert(tree.Node {
		Value: 10,
	})
	t.Insert(tree.Node {
		Value: 5,
	})
	t.Insert(tree.Node {
		Value: 15,
	})

	t.Insert(tree.Node {
		Value: 12,
	})

	fmt.Println("preorder:", t.InorderTraversal())
	fmt.Println("inorder:", t.PreorderTraversal())
	fmt.Println("postorder:", t.PostorderTraversal())
	fmt.Println("Is 5 present?", t.Search(5))
	fmt.Println("Is 17 present?", t.Search(17))
}
