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

func nextGreaterElements(nums []int) []int {
	h := []int{}
	for i := 0; i < len(nums); i++ {
		counter := 0
		j := i + 1
		for true {
			if j == len(nums) {
				j = 0
			}
			if nums[j] > nums[i] {
				h = append(h, nums[j])
				break
			}
			if counter == len(nums) { // whether there are no numbers greater than the current number
				h = append(h, -1)
				break
			}
			j++
			counter++
		}
	}
	return h
}

func nextGreaterElementsOn(nums []int) []int {

	nge := make([]int, len(nums))

	n := len(nums)

	var st Stack

	for i := 2*n - 1; i >= 0; i-- {

		for !st.Empty() && st.Top() <= nums[i%n] {
			st.Pop()
		}

		if i < n {
			if !st.Empty() {
				nge[i] = st.Top()
			} else {
				nge[i] = -1
			}
		}
		st.Push(nums[i%n])
	}

	return nge

}
func main() {

	var brr []int = []int{1, 2, 1}

	arr := nextGreaterElements(brr)

	fmt.Println(arr)

}
