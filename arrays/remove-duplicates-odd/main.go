package main

import "fmt"

func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {

		if nums[i]%2 == 0 {
			nums[j] = nums[i]
			j++
		} else {
			temp := nums[j]
			if i == 0 || nums[i-1] != nums[i] {
				nums[j] = nums[i]
				j++
			}

			if temp == nums[i] {
				continue
			}

		}

	}

	return j

}

func main() {

	var arr []int = []int{0, 0, 3, 3, 5, 5, 4, 4, 1, 1, 5}

	num := removeDuplicates(arr)

	for i := 0; i < num; i++ {
		fmt.Println(arr[i])
	}

}
