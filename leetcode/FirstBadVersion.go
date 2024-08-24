package leetcode

import "fmt"

func firstBadVersion(n int) int {
	start := 0
	end := n
	for start < end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}
func isBadVersion(version int) bool {
	if version > 10 {
		return true
	}
	return false
}
func VersionTest() {
	n := 100
	//bad := 12
	fmt.Printf("first Bad version %d", firstBadVersion(n))

}
