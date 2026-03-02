package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func hasPathSum(root *Node, target int) bool {
	if root == nil {
		return false
	}
	sumSlice, sum := []int{}, 0
	calculateSum(root, &sum, &sumSlice)
	for _, v := range sumSlice {
		if v == target {
			return true
		}
	}
	return false
}

func calculateSum(root *Node, sum *int, sumSlice *[]int) {
	if root == nil {
		return
	}
	*sum += root.Val
	if root.Left == nil && root.Right == nil {
		*sumSlice = append(*sumSlice, *sum)
		*sum -= root.Val
		return
	}
	calculateSum(root.Left, sum, sumSlice)
	calculateSum(root.Right, sum, sumSlice)
	*sum -= root.Val
}

func main() {
	fmt.Println("Path Sum in Binary Tree")
	root := &Node{Val: 5}
	root.Left = &Node{Val: 4}
	root.Right = &Node{Val: 8}
	root.Left.Left = &Node{Val: 11}
	root.Left.Left.Left = &Node{Val: 7}
	root.Left.Left.Right = &Node{Val: 2}
	root.Right.Left = &Node{Val: 13}
	root.Right.Right = &Node{Val: 4}
	root.Right.Right.Right = &Node{Val: 1}

	target := 22
	fmt.Printf("Does the binary tree have a path sum of %d? %v\n", target, hasPathSum(root, target))
}
