package main

import "fmt"

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {

		if nums[i] != val {
			nums[j] = nums[i]
			j++
		} else {
			continue

		}

	}

	return j

}

func main() {

	//var arr []int = []int{0, 1, 2, 2, 3, 0, 4, 2}
	var arr []int = []int{3, 2, 2, 3}

	val := 3

	j := removeElement(arr, val)

	for i := 0; i < j; i++ {
		fmt.Println(arr[i])
	}

}
