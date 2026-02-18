package main

import "fmt"

// design a tree struct
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// building a tree from slice
//most common way to turn slice into tree
//is level order insertion

func insertLevelOrder(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	root := &Node{Val: arr[0]}
	queue := []*Node{root}

	i := 1
	for i < len(arr) {
		cur := queue[0]
		queue = queue[1:]

		//insert left child
		if i < len(arr) {
			cur.Left = &Node{Val: arr[i]}
			queue = append(queue, cur.Left)
			i++
		}
		//insert right child
		if i < len(arr) {
			cur.Right = &Node{Val: arr[i]}
			queue = append(queue, cur.Right)
			i++
		}
	}
	return root
}

func (n *Node) Inorder() {
	if n == nil {
		return
	}
	n.Left.Inorder()
	fmt.Printf("%d", n.Val)
	n.Right.Inorder()
}

func main() {
	nodes := []int{10, 20, 30, 40, 50}
	fmt.Println("Input Slice:", nodes)
	root := insertLevelOrder(nodes)

	fmt.Print("Inorder Traversal: ")
	root.Inorder()
	fmt.Println()
}
