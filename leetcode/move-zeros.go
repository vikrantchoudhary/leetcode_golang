package leetcode

import "fmt"

func moveZeroes(nums []int) {
	pos := 0 // known position to fill
	result := make([]int, len(nums))
	for i := range nums {
		if nums[i] != 0 {
			result[pos] = nums[i]
			pos++
		}
	}
	for i := range nums {
		if i >= pos {
			nums[i] = 0
		} else {
			nums[i] = result[i]
		}
	}

	fmt.Println(nums)
}

func MoveZeroesTest() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
