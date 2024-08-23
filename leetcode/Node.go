package leetcode

import "fmt"

type Node struct {
	val  int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (list *Node) AddNode(data int) {
	node := &Node{val: data, next: nil}
	if list == nil {
		list = node
	}
	curr := list
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = node
	return
}

func PrintList(list *Node) {
	for list != nil {
		fmt.Print(list.val, " -->")
		list = list.next
	}
}

func PrintTest() {
	myList := Node{}

	myList.AddNode(10)
	myList.AddNode(20)
	myList.AddNode(30)
	myList.AddNode(40)
	myList.AddNode(40)
	myList.AddNode(50)
	myList.AddNode(60)
	myList.AddNode(60)
	myList.AddNode(60)
	PrintList(&myList)
}
