package tree

// import "fmt"

type BST struct {
	Root *Node
	Balanced bool
	numNodes int
}

type Node struct {
	Id string
	Value int
	/* Left => X < Value. Right: X >= Value */
	Left, Right *Node
}

func (tree *BST) Insert(node Node) {
	tree.numNodes++
	if tree.Root == nil {
		tree.Root = &node
		return
	}

	targetNode := findClosest(tree.Root, node.Value)
	if node.Value > targetNode.Value {
		targetNode.Right = &node
	} else {
		targetNode.Left = &node
	}
}

func (tree *BST) InorderTraversal() []int {
	if (tree.Root == nil) {
		return []int {}
	}

	var inorderTraversal func(node *Node, store []int, counter *int)
	inorderTraversal = func(node *Node, store []int, counter *int) {
		if (node == nil) {
			return
		}
	
		inorderTraversal(node.Left, store, counter)
		store[*counter] = node.Value
		*counter++
		inorderTraversal(node.Right, store, counter)
	
		return
	}

	traversedValues := make([]int, tree.numNodes)
	counter := 0
	inorderTraversal(tree.Root, traversedValues, &counter)

	return traversedValues
}

func (tree *BST) PreorderTraversal() []int {
	if (tree.Root == nil) {
		return []int {}
	}

	var preorderTraversal func(node *Node, store []int, counter *int)
	preorderTraversal = func(node *Node, store []int, counter *int) {
		if (node == nil) {
			return
		}
	
		store[*counter] = node.Value
		*counter++
		preorderTraversal(node.Left, store, counter)
		preorderTraversal(node.Right, store, counter)
	
		return
	}

	traversedValues := make([]int, tree.numNodes)
	counter := 0
	preorderTraversal(tree.Root, traversedValues, &counter)

	return traversedValues
}

func (tree *BST) PostorderTraversal() []int {
	if (tree.Root == nil) {
		return []int {}
	}

	var postorderTraversal func(node *Node, store []int, counter *int)
	postorderTraversal = func(node *Node, store []int, counter *int) {
		if (node == nil) {
			return
		}
	
		postorderTraversal(node.Left, store, counter)
		postorderTraversal(node.Right, store, counter)
		store[*counter] = node.Value
		*counter++
	
		return
	}

	traversedValues := make([]int, tree.numNodes)
	counter := 0
	postorderTraversal(tree.Root, traversedValues, &counter)

	return traversedValues
}

func (tree *BST) Search(value int) bool {
	if (tree.Root == nil) {
		return false
	}

	var searchNode func(root *Node, value int) bool
	searchNode = func(root *Node, value int) bool {
		if root == nil {
			return false
		}

		if value < root.Value {
			return searchNode(root.Left, value)
		} else if value > root.Value {
			return searchNode(root.Right, value)
		} else {
			return true
		}
	}

	return searchNode(tree.Root, value)
}

func findClosest(root *Node, value int) *Node {
	var closest *Node

	if value < root.Value {
		if root.Left == nil {
			closest = root
		} else {
			closest = findClosest(root.Left, value)
		}
	} else {
		if root.Right == nil {
			closest = root
		} else {
			closest = findClosest(root.Right, value)
		}
	}

	return closest
}
