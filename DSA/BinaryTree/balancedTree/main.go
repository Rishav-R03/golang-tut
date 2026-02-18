package main

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func IsBalanced(root *Node) bool {
	return dfs(root) != -1
}

func dfs(node *Node) int {
	if node == nil {
		return 0
	}
	lh := dfs(node.Left)
	if lh == -1 {
		return -1
	}
	rh := dfs(node.Right)
	if rh == -1 {
		return -1
	}
	if math.Abs((float64(lh - rh))) > 1 {
		return -1
	}
	return max(lh, rh) + 1

}
