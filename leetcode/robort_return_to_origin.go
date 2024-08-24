package leetcode

import (
	"fmt"
)

type Stack []int

func RobotMoves(moves string) bool {
	x := 0
	y := 0
	for _, v := range moves {
		switch string(v) {
		case "U":
			x++
		case "D":
			x--
		case "L":
			y--
		case "R":
			y++
		default:
			return false
		}

	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}
func JudgeCircle() {
	moves := "UD"
	if RobotMoves(moves) {
		fmt.Println("Success")
	} else {
		fmt.Println("failed")
	}

}
