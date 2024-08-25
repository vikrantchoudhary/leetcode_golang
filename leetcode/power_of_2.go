package leetcode

import "fmt"

func IsPowerOfTwo(n int) bool {
	for n > 0 {
		if n == 1 {
			return true
		}
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}
func IsPowerOfTwoTest() {
	n := 16
	if IsPowerOfTwo(n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("NO")
	}
}
