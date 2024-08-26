package leetcode

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 || nums[0] > nums[1] {
		return 0
	}
	n := len(nums) - 1

	low := 1
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid-1] {
			high = mid - 1
		} else if nums[mid] < nums[mid+1] {
			low = mid + 1
		}
	}
	return -1
}

func findPeakElement1(nums []int) int {
	n := len(nums)
	if n < 3 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid-1] {
			high = mid - 1
		} else if nums[mid] < nums[mid+1] {
			low = mid + 1
		}
	}

	return -1
}

func FindPeakElementTest() {
	nums := []int{6, 5, 4, 3, 2, 3, 2}
	fmt.Println(findPeakElement(nums))

}
