package main

func findPoisonedDuration(timeSeries []int, duration int) int {

	var a, b, sum int

	b = -1

	for i := 0; i < len(timeSeries); i++ {

		if timeSeries[i] <= b {
			sum = 0
			b = timeSeries[i] + duration - 1
		} else {
			a = timeSeries[i]
			b = timeSeries[i] + duration - 1
		}

		sum = sum + b - a + 1
	}

	return sum

}

func main() {

}
