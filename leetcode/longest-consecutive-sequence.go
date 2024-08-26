package leetcode

import "fmt"

func longestConsecutive(nums []int) int {

	count := 0

	//set := make(map[int]int)
	set := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		fmt.Print(" \t ", nums[i])
		set[nums[i]] = true
	}
	fmt.Println()
	fmt.Println(set)

	for n := range nums {
		if set[n+1] {
			continue
		}
		cur_count := 0
		for cur := n; set[cur]; cur-- {
			cur_count++
		}
		if cur_count > count {
			count = cur_count
		}
	}
	return count
}

func LongestConsecutive() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	//nums2 := []int{100, 4, 200, 1, 3, 2}
	fmt.Printf("Longest consecutive seq = %d ", longestConsecutive(nums))
}
