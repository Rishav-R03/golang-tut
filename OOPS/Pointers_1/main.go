package main

import "fmt"

// var x int = 10 stores value 10
// var p *int = &x stores the address where 10 is sitting

// &: (Address-of) => Where is this variable located?
// *: (Dereferencing) => Go to this address and tell me what value is there

type Node struct {
	Val int
}

// This WON'T change the original node
func (n Node) UpdateValue(v int) {
	n.Val = v
}

// This is will change
func (n *Node) RealChange(v int) {
	n.Val = v
}
func GetAdd(num *int) int {
	return *num
}

func DoubleValue(n *int) {
	*n = (*n) * 2
	// *n *=2
}

func main() {
	var x int = 10
	fmt.Println("X has value: ", x)

	var y int = 5
	var p *int = &x
	fmt.Println("P is the pointer to address of value of X", p)

	fmt.Println(GetAdd(p))

	var p2 *int = &y

	var z = *p + *p2
	fmt.Println("The sum of X and Y ", z)

	var node = Node{Val: 2}

	node.UpdateValue(5)
	fmt.Println("This is unchanged as 2", node.Val)

	node.RealChange(6)
	fmt.Println("This is updated: ", node.Val)

	fmt.Println("Before doubling ", x)
	DoubleValue(p)
	fmt.Print("After doubling ", x)

}
