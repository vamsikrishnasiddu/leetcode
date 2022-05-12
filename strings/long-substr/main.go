// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func findSubstrings(str string) int {

	s := []string{}
	var strs string

	for i := 0; i < len(str); i++ {
		strs = ""

		for j := i; j < len(str); j++ {
			strs = strs + string(str[j])
			s = append(s, strs)

		}

	}

	var max = math.MinInt32
	var check bool

	m := make(map[byte]int)

	//s = []string{"vdf"}

	for _, v := range s {
		check = false

		for j := 0; j < len(v); j++ {

			if _, ok := m[v[j]]; !ok {

				m[v[j]] = j
			} else {

				check = true
				break
			}

		}

		for k := range m {
			delete(m, k)
		}

		if !check && len(v) > max {

			max = len(v)

		}

	}

	return max

}

func main() {
	str := "vdvdfdvdd"
	fmt.Println(findSubstrings(str))

}
