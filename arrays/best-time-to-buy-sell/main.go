package main

func maxProfit(prices []int) int {

	l := 0
	r := 1
	maxP := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			maxP = max(maxP, prices[r]-prices[l])
		} else {
			l = r

		}

		r = r + 1
	}

	return maxP

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	var arr []int = []int{1, 2, 3, 4, 5}

	maxProfit(arr)
}
