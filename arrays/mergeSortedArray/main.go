package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	nL := m
	nR := n

	nums3 := make([]int, m+n)

	i := 0
	j := 0
	k := 0

	for i < nL && j < nR {

		if nums1[i] <= nums2[j] {
			nums3[k] = nums1[i]
			i++
		} else if nums2[j] <= nums1[i] {
			nums3[k] = nums2[j]
			j++
		}

		k++
	}

	for i < nL {
		nums3[k] = nums1[i]
		i++
		k++
	}

	for j < nR {
		nums3[k] = nums2[j]
		j++
		k++
	}

	copy(nums1, nums3)

}

func main() {

	//nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3

	var nums1 []int = []int{1, 2, 3, 0, 0, 0}
	var nums2 []int = []int{2, 5, 6}
	var m int = 3
	var n int = 3

	merge(nums1, m, nums2, n)

	fmt.Println(nums1)

}
