package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	var arr []int = nums
	firstMax := math.MinInt64
	secondMax := math.MinInt64
	thirdMaxx := math.MinInt64

	for i := 0; i < len(arr); i++ {
		if arr[i] > firstMax {
			firstMax = arr[i]
		}

		for j := 0; j <= i; j++ {
			if arr[j] < firstMax && arr[j] > secondMax {
				secondMax = arr[j]
			}
		}

		for k := 0; k <= i; k++ {

			if arr[k] < secondMax && arr[k] > thirdMaxx {
				thirdMaxx = arr[k]
			}

		}

	}

	if thirdMaxx == math.MinInt64 {
		return firstMax
	}

	return thirdMaxx
}

func main() {

	var arr []int = []int{1, 2}

	fmt.Println(thirdMax(arr))

}
