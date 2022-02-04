package main

import (
	"fmt"
	"math"
)

func maxSubArray(arr []int) int {

	max := math.MinInt16
	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max

}
func main() {
	var arr []int = []int{-2, 1, -1, -3, 1, -4}

	max := maxSubArray(arr)

	fmt.Println(max)
}
