package tree

import (
	"algorithms/math"
	"fmt"
)

// BST : The tree
type BST struct {
	Root        *treeNode
	AutoBalance bool
	numNodes    int
}

type treeNode struct {
	ID     string
	Value  int
	height int
	/* if multiple values present */
	count int
	/* AVL balance factor = <left height> - <right height> */
	factor int
	/* Left => X < Value. Right: X > Value */
	Left, Right *treeNode
}

func (node treeNode) String() string {
	return fmt.Sprintf("{V:%d, C:%d, H:%d, F:%d}", node.Value, node.count, node.height, node.factor)
}

// Insert the value into the tree
func (tree *BST) Insert(value int) {
	if tree.Root == nil {
		tree.Root = newTreeNode(value)
		tree.numNodes = 1
		return
	}

	tree.numNodes++

	var findAndInsert func(root *treeNode, value int) *treeNode
	findAndInsert = func(root *treeNode, value int) *treeNode {
		if root.Value == value {
			root.count++
			return root
		}

		if value < root.Value {
			if root.Left == nil {
				root.Left = newTreeNode(value)
			} else {
				root.Left = findAndInsert(root.Left, value)
			}
		} else {
			if root.Right == nil {
				root.Right = newTreeNode(value)
			} else {
				root.Right = findAndInsert(root.Right, value)
			}
		}

		return checkAndRebalance(root, tree.AutoBalance)
	}

	tree.Root = findAndInsert(tree.Root, value)
}

/*
 * Removes the given value from the BST if present.
 *
 * Returns: false if element is not present
 *          true if element is present
 */
func (tree *BST) Remove(value int) bool {
	if tree.Root == nil {
		tree.numNodes = 0
		return false
	}

	var removeFromNode func(root *treeNode, value int) (bool, *treeNode)
	removeFromNode = func(root *treeNode, value int) (bool, *treeNode) {
		if root == nil {
			return false, root
		}

		var isRemoved bool
		finalNode := root
		if value == root.Value {
			if root.count > 1 {
				root.count--
			} else {
				finalNode = merge(root.Left, root.Right, tree.AutoBalance)
			}
			isRemoved = true
		} else if value < root.Value {
			isRemoved, root.Left = removeFromNode(root.Left, value)
		} else {
			isRemoved, root.Right = removeFromNode(root.Right, value)
		}

		return isRemoved, checkAndRebalance(finalNode, tree.AutoBalance)
	}

	var result bool
	result, tree.Root = removeFromNode(tree.Root, value)
	if result {
		tree.numNodes--
	}

	return result
}

// InorderTraversal recursively <left><root><right>. Returns a sorted array
func (tree *BST) InorderTraversal() []int {
	if tree.Root == nil {
		return []int{}
	}

	var inorderTraversal func(node *treeNode, store []int, counter *int)
	inorderTraversal = func(node *treeNode, store []int, counter *int) {
		if node == nil {
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

// PreorderTraversal recursively <root><left><right>
func (tree *BST) PreorderTraversal() []int {
	if tree.Root == nil {
		return []int{}
	}

	var preorderTraversal func(node *treeNode, store []int, counter *int)
	preorderTraversal = func(node *treeNode, store []int, counter *int) {
		if node == nil {
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

// PostorderTraversal recursively <left><right><root>
func (tree *BST) PostorderTraversal() []int {
	if tree.Root == nil {
		return []int{}
	}

	var postorderTraversal func(node *treeNode, store []int, counter *int)
	postorderTraversal = func(node *treeNode, store []int, counter *int) {
		if node == nil {
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

// Search the value and returns true if the value is present in bst else false
func (tree *BST) Search(value int) bool {
	var searchNode func(root *treeNode, value int) bool
	searchNode = func(root *treeNode, value int) bool {
		if root == nil {
			return false
		}

		if value == root.Value {
			return true
		} else if value < root.Value {
			return searchNode(root.Left, value)
		} else {
			return searchNode(root.Right, value)
		}
	}

	return searchNode(tree.Root, value)
}

// PrintNodes prints the tree in level order
func (tree *BST) PrintNodes() {
	if tree.Root == nil {
		return
	}

	printNodes(tree.Root)
}

// PrintNodes prints the tree in level order
func printNodes(root *treeNode) {
	fmt.Println("##### Printing tree Start #####")
	var queue []*treeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		formattedStr := ""
		var stage []*treeNode
		for _, node := range queue {
			formattedStr = fmt.Sprintf("%v, %v", formattedStr, node)
			if node != nil {
				stage = append(stage, node.Left)
				stage = append(stage, node.Right)
			}
		}
		queue = stage

		fmt.Println(formattedStr)
	}
	fmt.Println("##### Printing tree End #####")
}

func merge(left *treeNode, right *treeNode, autoBalance bool) *treeNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	var leftNode *treeNode
	var rightNode *treeNode
	if left.Value > right.Value {
		leftNode = right
		rightNode = left
	} else {
		leftNode = left
		rightNode = right
	}

	var mergedNode *treeNode
	if leftNode.height >= rightNode.height {
		leftNode.Right = merge(leftNode.Right, rightNode, autoBalance)
		mergedNode = leftNode
	} else {
		rightNode.Left = merge(rightNode.Left, leftNode, autoBalance)
		mergedNode = rightNode
	}

	return checkAndRebalance(mergedNode, autoBalance)
}

func checkAndRebalance(root *treeNode, autoBalance bool) *treeNode {
	updateNodeHeightInfo(root)

	if !autoBalance || root == nil || root.factor >= -1 && root.factor <= 1 {
		return root
	}

	finalRoot := root
	if root.factor < -1 {
		if root.Right != nil && root.Right.factor > 0 {
			finalRoot.Right = rRotate(root.Right)
		}

		finalRoot = lRotate(root)
	} else if root.factor > 1 {
		if root.Left != nil && root.Left.factor < 0 {
			finalRoot.Left = lRotate(root.Left)
		}

		finalRoot = rRotate(root)
	}

	return finalRoot
}

func lRotate(root *treeNode) *treeNode {
	if root == nil || root.Right == nil {
		return root
	}

	right := root.Right
	root.Right = right.Left
	right.Left = root

	updateNodeHeightInfo(root)
	updateNodeHeightInfo(right)

	return right
}

func rRotate(root *treeNode) *treeNode {
	if root == nil || root.Left == nil {
		return root
	}

	left := root.Left
	root.Left = left.Right
	left.Right = root

	updateNodeHeightInfo(root)
	updateNodeHeightInfo(left)

	return left
}

func updateNodeHeightInfo(root *treeNode) {
	if root == nil {
		return
	}

	leftHeight := 0
	rightHeight := 0
	if root.Left != nil {
		leftHeight = root.Left.height
	}
	if root.Right != nil {
		rightHeight = root.Right.height
	}

	root.height = int(math.Max(leftHeight, rightHeight) + 1)
	root.factor = leftHeight - rightHeight
}

func newTreeNode(value int) *treeNode {
	return &treeNode{
		Value:  value,
		count:  1,
		height: 1,
		factor: 0,
	}
}
