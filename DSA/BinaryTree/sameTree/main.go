package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func isSame(root1 *Node, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root1 != nil && root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}

func main() {
	// Create two binary trees
	root1 := &Node{Val: 1}
	root1.Left = &Node{Val: 2}
	root1.Right = &Node{Val: 3}

	root2 := &Node{Val: 1}
	root2.Left = &Node{Val: 2}
	root2.Right = &Node{Val: 3}

	println(isSame(root1, root2)) // Output: true
}
