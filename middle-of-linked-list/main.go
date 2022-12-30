package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func lenList(head *ListNode) int {
	len := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		len++
	}
	return len
}

func middleNode(head *ListNode) *ListNode {
	len := lenList(head)
	var middleIndex int
	if (len % 2 == 1) {
		middleIndex = (len - 1) / 2
	} else {
		middleIndex = len / 2
	}
	tmp := head
	for i := 1; i <= middleIndex; i++ {
		tmp = tmp.Next
	}
	return tmp
}

func main() {
	var one, two, three, four ListNode
	one = ListNode{
		Val: 1,
		Next: &two,
	}
	two = ListNode{
		Val: 2,
		Next: &three,
	}
	three = ListNode{
		Val: 3,
		Next: &four,
	}
	four = ListNode{
		Val: 4,
		Next: nil,
	}
	fmt.Println("len of list :", lenList(&one))
	fmt.Println("middle node :", middleNode(&one))
}