package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (head *ListNode) reverseKGroup(k int) *ListNode {
	tail := head // 1
	for i:= 0; i < k; i++ {
		tail = tail.Next // 3
	}
	curr, newHead := head.reverseUntil(tail) // [2 1 345], [345]
	fmt.Printf("from curr: %v\n", curr)
	for curr.Next != nil {
		curr, _ = curr.reverseUntil(curr.getTail(k))
		fmt.Println("prev: ")
		//prev.Next = newHead
		//prev = curr
	}
	return newHead
}

func (head *ListNode) getTail(k int) *ListNode {
	tail := head // 1
        for i:= 0; i < k; i++ {
		if tail != nil {
                	tail = tail.Next // 2
		}
        }
	return tail
}

func (head *ListNode) String() string {
	var s string
	for head.Next != nil {
		s += strconv.Itoa(head.Val) + " -> "
		head = head.Next
	}
	return s + strconv.Itoa(head.Val) + "\n"
}

func (head *ListNode) reverseUntil(tail *ListNode) (*ListNode, *ListNode) {
	if tail == nil {
		return head, nil
	}
	curr := head
	var next *ListNode
	prev := tail
	for curr != tail {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	fmt.Println("-----------------")
	//curr.Next = tail
	return tail, prev
}

func main() {
	l := &ListNode{Val: 1}
	l.Next = &ListNode{Val: 2}
	l.Next.Next = &ListNode{Val: 3}
	l.Next.Next.Next = &ListNode{Val: 4}
	l.Next.Next.Next.Next = &ListNode{Val: 5}
	k := l.reverseKGroup(2)
	fmt.Printf("RESULT \n %v", k)
}
