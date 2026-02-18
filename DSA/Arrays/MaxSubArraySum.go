package main

import (
	"fmt"
	"math"
)

func maxSubArraySum(nums []int) int {
	cur := nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		cur = int(math.Max(float64(nums[i]+cur), float64(nums[i])))
		max = int(math.Max(float64(max), float64(cur)))
	}
	return max
}

func main() {
	var nums = []int{1, -2, 3, 5, -6, -5}
	var ans = maxSubArraySum(nums)
	fmt.Println(ans)
}
