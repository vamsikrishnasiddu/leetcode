package main

import "fmt"

func check(val int, arr []int) bool {

	element := arr[val]

	if element <= val {
		return true
	}

	return false
}

func binarySearch(arr []int, target int) int {

	l := 0
	r := len(arr) - 1

	var mid int = -1

	//var check bool

	for l < r {

		mid = (l + r) / 2

		if check(mid, arr) {
			r = mid
		} else {
			l = mid + 1
		}

	}

	if arr[l] == target {
		return l

	}
	return -1
}

func main() {

	var arr []int = []int{3, 5, 7, 9, 10}

	target := 8

	val := binarySearch(arr, target)
	if val == -1 {
		fmt.Println("element not found.")
	} else {
		fmt.Println(val)
	}
}
