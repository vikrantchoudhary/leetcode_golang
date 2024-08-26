package leetcode

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverseArray(nums, 0, n-1)
	reverseArray(nums, k, n-1)
	reverseArray(nums, 0, k-1)
	fmt.Println(nums)

}
func reverseArray(nums []int, start int, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
func rotateExtraSpace(nums []int, k int) {
	n := len(nums)
	temp := make([]int, k)
	for i := 0; i < k; i++ {
		temp[i] = nums[n+i-k]
		fmt.Println(temp[i])
	}
	//fmt.Println(temp)
	//j := 0
	for i := n - 1; i >= 0; i-- {
		if i >= n-k {
			nums[i] = nums[i-n+k]
			fmt.Println(i)
			fmt.Println(i - n + k + 1)
		}
		/*else {
			nums[i] = temp[k-j]
			j++
		}*/

	}
	fmt.Println(nums)

}
func RotateArrayTest() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
}
