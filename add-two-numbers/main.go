package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

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

func listToArray(head *ListNode) []int {
	var arr []int
	for tmp := head; tmp != nil; tmp = tmp.Next {
		arr = append(arr, tmp.Val)
	}
	return arr
}

func arrToList(head *ListNode, arr []int) *ListNode {
	for _, v := range arr {
		head = addLastNode(head, v)
	}
	return head
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	arr1 := listToArray(l1)
	arr2 := listToArray(l2)
	lenMax := 0
	if (len(arr1) >= len(arr2)) {
		lenMax = len(arr1)
		for i := len(arr2); i < lenMax; i++ {
			arr2 = append(arr2, 0)
		}
	} else {
		lenMax = len(arr2)
		for i := len(arr1); i < lenMax; i++ {
			arr1 = append(arr1, 0)
		}
	}
	fmt.Println(arr1, arr2)
	var arr3 []int
	carry := 0
	for i := 0; i < lenMax; i++ {
		arr3 = append(arr3, (arr1[i] + arr2[i] + carry) % 10)
		if (arr1[i] + arr2[i] + carry >= 10) {
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		arr3 = append(arr3, 1)
	}
	fmt.Println(arr3)
	var head *ListNode
	head = arrToList(head, arr3)
	return head
}

func main() {
	list1 := []int{9,9,9,9,9,9,9}
	list2 := []int{9,9,9,9}
	var l1, l2 *ListNode
	for _, v := range list1 {
		l1 = addLastNode(l1, v)
	}
	for _, v := range list2 {
		l2 = addLastNode(l2, v)
	}
	fmt.Println(addTwoNumbers(l1, l2))
}