package leetcode

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	if len(nums) < 1 {
		return 0
	}
	j := 0
	n := len(nums) - 1
	for i := 1; i <= n; i++ {
		if nums[j] == nums[i] {
			count++
		} else {
			j++
			nums[j] = nums[i]

		}
	}

	/*if j != 0 && nums[j] != nums[n] {
		nums[j] = nums[n]
		count++
	}*/
	for i := 1; i < n-j+1; i++ {
		nums[j+i] = 0
	}
	fmt.Println(nums)

	return count
}

func RemoveDuplicatesArrayTest() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
