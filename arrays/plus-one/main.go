package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	var sum int
	var carry int = 1

	for i := 1; i < len(digits)+1; i++ {
		sum = digits[len(digits)-i] + carry

		if sum/10 == 0 {
			carry = 0
			digits[len(digits)-i]++
			break
		} else {
			carry = sum / 10
			digits[len(digits)-i] = 0
			sum = 0
		}
	}
	fmt.Println(carry)

	if carry == 1 {
		arr := make([]int, len(digits)+1)
		arr[0] = 1
		return arr
	}

	return digits

}

func main() {
	var arr []int = []int{8, 9, 9, 9}

	a := plusOne(arr)

	fmt.Println(a)

}
