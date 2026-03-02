package main

import "math"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func createTree(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	root := &Node{Value: nums[0]}
	queue := []*Node{root}

	ind := 1

	for ind < len(nums) {
		curr := queue[0]
		queue = queue[1:]

		if ind < len(nums) {
			if nums[ind] != -1 {
				curr.Left = &Node{Value: nums[ind]}
				queue = append(queue, curr.Left)
			}
			ind++
		}

		if ind < len(nums) {
			if nums[ind] != -1 {
				curr.Right = &Node{Value: nums[ind]}
				queue = append(queue, curr.Right)
			}
			ind++
		}
	}
	return root
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))
}
func main() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := createTree(nums)
	println(maxDepth(root)) // Output: 3
}
