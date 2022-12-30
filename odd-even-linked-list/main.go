package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	evenHead := even
	for ; even != nil && even.Next != nil; {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// Test
func createNode(val int) ListNode {
	node := ListNode{
		Val: val,
		Next: nil,
	}
	return node
}

func addLastNode(head *ListNode, val int) *ListNode {
	if head == nil {
		node := createNode(val)
		head = &node
		return head
	}
	tmp := head
	for ; tmp.Next != nil; {
		tmp = tmp.Next
	}
	node := createNode(val)
	tmp.Next = &node
	return head
}

func arrToList(head *ListNode, arr []int) *ListNode {
	for _, v := range arr {
		head = addLastNode(head, v)
	}
	return head
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var head *ListNode
	head = arrToList(head, arr)
	head = oddEvenList(head)
	for tmp := head; tmp != nil; tmp = tmp.Next {
		fmt.Println(tmp.Val)
	}
}