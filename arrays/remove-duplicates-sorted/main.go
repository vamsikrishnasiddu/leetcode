package main

import "fmt"

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		temp := nums[j]

		if nums[i] == temp {
			continue
		}

		j++
		nums[j] = nums[i]
	}

	return j + 1

}

func main() {

	var arr []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//var arr []int = []int{1, 1}
	//var arr []int = []int{}

	num := removeDuplicates(arr)
	fmt.Println(num)

	for i := 0; i < num; i++ {

		fmt.Println(arr[i])

	}

}
