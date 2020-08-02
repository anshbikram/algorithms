package main

import (
	"algorithms/tree"
	"fmt"
)

func main() {
	t := tree.BST{
		AutoBalance: true,
	}
	t.Insert(25)
	t.Insert(50)
	t.Insert(75)
	t.Insert(100)
	t.Insert(125)
	t.Insert(150)
	t.Insert(175)

	fmt.Println("inorder:", t.InorderTraversal())
	fmt.Println("preorder:", t.PreorderTraversal())
	fmt.Println("postorder:", t.PostorderTraversal())
	fmt.Println("Search 5 =>", t.Search(5))
	fmt.Println("Search 17 =>", t.Search(175))

	t.PrintNodes()
	t.Remove(100)
	t.Remove(150)
	t.Remove(175)
	t.Remove(175)
	t.PrintNodes()
}
