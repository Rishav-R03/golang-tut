package main

import (
	"fmt"
	"math"
	"strconv"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func InsertIntoTree(nums []int) *Node {
	///Base case check
	if len(nums) == 0 {
		return nil
	}

	// Form first node
	root := &Node{Val: nums[0]}
	queue := []*Node{root}

	i := 1
	for i < len(nums) {
		cur := queue[0]
		queue = queue[1:]

		if i < len(nums) {
			if nums[i] != -1 {
				cur.Left = &Node{Val: nums[i]}
				queue = append(queue, cur.Left)
			}
			i++
		}

		if i < len(nums) {
			if nums[i] != -1 {
				cur.Right = &Node{Val: nums[i]}
				queue = append(queue, cur.Right)
			}
			i++
		}
	}
	return root
}

func (root *Node) InorderTraversal() {
	if root == nil {
		return
	}

	root.Left.InorderTraversal()
	fmt.Print(strconv.Itoa(root.Val) + " ")
	root.Right.InorderTraversal()
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
	return int(math.Max(float64(lh), float64(rh))) + 1

}

func main() {
	nums := []int{1, -1, 2, -1, 3, -1, 4} // unbalanced
	root := InsertIntoTree(nums)
	root.InorderTraversal()
	if IsBalanced(root) {
		fmt.Println("\n The tree is balanced")
	} else {
		fmt.Println("\n The tree is not balanced")
	}

}
