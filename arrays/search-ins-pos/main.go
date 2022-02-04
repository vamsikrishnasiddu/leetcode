package main

import "fmt"

func searchInsert(nums []int, target int) int {
	var l int = 0
	var r int = len(nums) - 1
	var mid int

	for r >= l {
		mid = (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			l = mid + 1
		}

		if target < nums[mid] {
			r = mid - 1
		}
	}

	if target < nums[mid] {
		return mid
	}
	return mid + 1

}

func main() {

	var arr []int = []int{1, 2, 4, 6, 7}
	var target int = 3

	pos := searchInsert(arr, target)
	fmt.Println(pos)

}
