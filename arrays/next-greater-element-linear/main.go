package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() int {

	return s.items[len(s.items)-1]

}

func (s *Stack) Empty() bool {
	if len(s.items) == 0 {
		return true
	}

	return false
}

func nextGreaterElementLinear(arr []int) []int {

	var brr []int = []int{}

	for i := 0; i < len(arr); i++ {
		large := -1

		for j := i + 1; j < len(arr); j++ {

			if arr[j] > arr[i] {
				large = arr[j]
				break
			}

		}

		brr = append(brr, large)

	}

	return brr
}

func nextGreaterElementLinearOn(arr []int) []int {

	var brr []int = make([]int, len(arr))

	var st Stack

	for i := len(arr) - 1; i >= 0; i-- {

		for !st.Empty() && st.Top() <= arr[i] {
			st.Pop()
		}

		if !st.Empty() {

			brr[i] = st.Top()

		} else {
			brr[i] = -1
		}

		st.Push(arr[i])

	}

	return brr
}

func main() {

	var arr []int = []int{1, 2, 3, 4}

	fmt.Println(nextGreaterElementLinearOn(arr))

}
