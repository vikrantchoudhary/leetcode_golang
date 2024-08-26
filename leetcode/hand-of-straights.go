package leetcode

import (
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)
	if n == 0 || n%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	dmap := map[int]int{}
	for i := 0; i < n; i++ {
		dmap[hand[i]]++
	}
	for _, nums := range hand {
		if dmap[nums] == 0 {
			continue
		}
		for j := 0; j < groupSize; j++ {
			if dmap[nums+j] != 0 {
				dmap[nums+j]--
			} else {
				return false
			}
		}

	}

	return true
}

func IsNStraightHandTest() {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	groupSize := 3
	if isNStraightHand(hand, groupSize) {
		fmt.Println("Hand of Straights are possible")
	} else {
		fmt.Println("not possible")
	}
}
