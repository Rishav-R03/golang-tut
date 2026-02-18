package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func CreateTree(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	root := &Node{Val: nums[0]}
	queue := []*Node{root}

	i := 1

	for i < len(nums) {
		cur := queue[0]
		queue := queue[1:]

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

func MaxDepthOfBT(root *Node) int {
	if root == nil {
		return 0
	}

	var leftTree int = MaxDepthOfBT(root.Left)
	var rightTree int = MaxDepthOfBT(root.Right)

	return 1 + max(leftTree, rightTree)
}

func main() {
	nums := []int{1, 2, 3, 4, -1, -1}
	root := CreateTree(nums)

	ans := MaxDepthOfBT(root)
	fmt.Println("We have max depth as \n", ans)
}
