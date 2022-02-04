// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	var sum int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == 1 {
				sum = sum + 4

				if i-1 >= 0 {
					if grid[i-1][j] == 1 {
						sum = sum - 1

					}
				}

				if j-1 >= 0 {

					if grid[i][j-1] == 1 {
						sum = sum - 1

					}

				}

				if i+1 <= len(grid)-1 {

					if grid[i+1][j] == 1 {
						sum = sum - 1
					}

				}

				if j+1 <= len(grid[i])-1 {

					if grid[i][j+1] == 1 {
						sum = sum - 1
					}

				}

			}
		}
	}

	return sum

}

func main() {
	var grid [][]int = [][]int{{1, 1}}
	fmt.Println(islandPerimeter(grid))

}
