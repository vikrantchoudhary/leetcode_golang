package leetcode

import "fmt"

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	node := new(Node)
	result := node
	rem := 0
	carry := 0
	for l1 != nil || l2 != nil {
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val = l1.val
			l1 = l1.next
		}
		if l2 != nil {
			l2Val = l2.val
			l2 = l2.next
		}
		rem = (carry + l1Val + l2Val) % 10
		carry = (carry + l1Val + l2Val) / 10
		node.next = &Node{val: rem, next: nil}
		node = node.next

	}
	if carry != 0 {
		node.next = &Node{val: carry, next: nil}
	}

	return result.next

}
func TwoListSumTest() {
	myList := Node{}
	myList.AddNode(2)
	myList.AddNode(4)
	myList.AddNode(9)
	myList2 := Node{}
	myList2.AddNode(5)
	myList2.AddNode(6)
	myList2.AddNode(4)
	myList2.AddNode(9)
	PrintList(&myList)
	fmt.Println()
	PrintList(&myList2)
	result := addTwoNumbers(&myList, &myList2)
	fmt.Println()
	PrintList(result)

}
