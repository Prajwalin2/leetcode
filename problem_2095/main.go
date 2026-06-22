package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next

	return head
}
func main() {
	var arr = []int{1, 3, 4, 7, 1, 2, 6}
	list := readLinkedList(arr)
	list = deleteMiddle(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

func readLinkedList(arr []int) *ListNode {
	head := &ListNode{}
	node := head
	for _, n := range arr {
		node.Next = &ListNode{Val: n}
		node = node.Next
	}
	return head.Next
}
