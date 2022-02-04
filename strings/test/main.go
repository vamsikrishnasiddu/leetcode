package main

import (
	"fmt"
)

// func isValid(s string) bool {

// 	strArr := strings.Split(s, "")

// 	results := make([]bool, 3)
// 	var countTrues int

// 	if strArr[0] == "(" || strArr[0] == "{" || strArr[0] == "[" && len(strArr)%2 != 0 {

// 		for i := 0; i < len(strArr); i++ {
// 			switch strArr[i] {
// 			case "(":
// 				for j := i + 1; j < len(strArr); j++ {
// 					if strArr[j] == ")" {
// 						results[0] = true
// 						countTrues++
// 					}
// 				}

// 			case "{":
// 				for j := i + 1; j < len(strArr); j++ {
// 					if strArr[j] == "}" {
// 						results[1] = true
// 						countTrues++
// 					}
// 				}

// 			case "[":
// 				for j := i + 1; j < len(strArr); j++ {
// 					if strArr[j] == "]" {
// 						results[2] = true
// 						countTrues++
// 					}
// 				}

// 			}
// 		}

// 		var count = 0

// 		for i := 0; i < len(results); i++ {
// 			if results[i] == true {
// 				count++
// 			}
// 		}

// 		if count == countTrues {
// 			return true
// 		}

// 		return false

// 	} else {
// 		return false
// 	}

//}

func isValid(s string) bool {
	var q []rune
	for _, c := range s {

		switch c {
		case '(':
			q = append(q, ')')
		case '{':
			q = append(q, '}')
		case '[':
			q = append(q, ']')
		default:
			if len(q) == 0 || c != q[len(q)-1] {
				return false
			}

			q = q[:len(q)-1]

		}

	}

	return len(q) == 0

}

func main() {

	fmt.Println(isValid("([])"))

}
