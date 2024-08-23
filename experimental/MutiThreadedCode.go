package experimental

import (
	"fmt"
	"time"
)

func PrintAlphabets() {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, v := range letters {
		fmt.Printf("%c", v)
		time.Sleep(5 * time.Millisecond)
	}
}
func PrintDigits() {
	for i := 1; i < 20; i++ {
		fmt.Printf("%d", i)
		time.Sleep(5 * time.Millisecond)
	}
}
func SingleThreadedPrint() {
	PrintAlphabets()
	fmt.Println()
	PrintDigits()
}
func MutiThreadePrint() {
	go PrintAlphabets()
	go PrintDigits()
	time.Sleep(2 * time.Second)
}
