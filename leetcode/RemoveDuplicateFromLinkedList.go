package leetcode

import "fmt"

func RemoveDuplicatedFromList(list *Node) {
	if list == nil {
		return
	}
	cur := list
	for cur != nil && cur.next != nil {
		if cur.val == cur.next.val {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}

	}
}
func RemoveDuplicateTest() {
	myList := Node{}
	myList.AddNode(10)
	myList.AddNode(10)
	myList.AddNode(20)
	myList.AddNode(20)
	myList.AddNode(20)
	myList.AddNode(20)
	PrintList(&myList)
	fmt.Println()
	RemoveDuplicatedFromList(&myList)
	PrintList(&myList)

}
