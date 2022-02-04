package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var indices []int = []int{}
	var m = make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {

		if _, ok := m[target-nums[i]]; ok {

			indices = append(indices, m[target-nums[i]], i)

		} else {
			m[nums[i]] = i
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
