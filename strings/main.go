// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {

	str := strings.Split(s, "")

	sum := 0

	m := make(map[string]int)

	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	for i := 0; i < len(str); i++ {

		if (i+1 < len(str)) && ((str[i] == "I" && (str[i+1] == "V" || str[i+1] == "X")) || (str[i] == "X" && (str[i+1] == "L" || str[i+1] == "C")) || (str[i] == "C" && (str[i+1] == "D" || str[i+1] == "M"))) {
			sum = sum + (m[str[i+1]] - m[str[i]])
			i = i + 1
		} else {
			sum = sum + m[str[i]]
		}

	}

	return sum

}

func main() {
	s := "MDLXX"
	fmt.Println(romanToInt(s))
}
