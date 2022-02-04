package main

import "fmt"

func twoSum(nums []int, target int) []int {

	var indices []int = []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				indices = append(indices, i, j)

				return indices

			}

		}
	}
	return indices
}

func main() {
	arr := make([]int, 4)
	arr = []int{2, 7, 11, 15}

	brr := twoSum(arr, 9)

	fmt.Println(brr)
}
