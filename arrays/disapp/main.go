// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func findDisappearedNumbers(nums []int) []int {

	var i, j int
	n := len(nums)

	var arr []int = make([]int, n)
	var brr []int = []int{}

	prev := -1

	for k := 0; k < n; k++ {
		arr[k] = k + 1
	}

	sort.Slice(nums, func(p, q int) bool {
		return nums[p] < nums[q]
	})

	fmt.Println(arr)

	for i < n && j < n {

		if nums[i] == arr[j] {

			i++
			j++

		} else {
			//prev = nums[i]
			if nums[i] == prev {
				i++
			} else {
				brr = append(brr, arr[j])
			}

		}
		if i < len(nums) {
			j++
			prev = nums[i]
		}

	}

	return brr

}

func main() {
	var arr []int = []int{4, 3, 2, 7, 8, 2, 3, 1}

	fmt.Println(findDisappearedNumbers(arr))
}
