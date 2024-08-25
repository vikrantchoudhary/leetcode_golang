package leetcode

import (
	"fmt"
	"slices"
)

func SquareArray(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		arr[i] *= arr[i]
	}
	slices.Sort(arr)
	fmt.Println(arr)
	return arr
}
