package leetcode

import "fmt"

func FirstUniqChar(s string) int {
	var ch = make([]int, 26)
	for _, v := range s {
		ch[v-'a']++
	}
	for i := 0; i < len(ch); i++ {
		if ch[i] == 1 {
			return i
		}
	}
	return -1

}
func FirstUniqCharTest() {
	s := "loveleetcode"
	pos := FirstUniqChar(s)
	fmt.Printf("Position non repeating of char in the string %s is  %d ", s, pos)
}
