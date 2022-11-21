package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) print() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func getListNodeFromInteger(n int) *ListNode {
	l := &ListNode{}
	p := l
	for n >= 10 {
		p.Val = n % 10
		p.Next = &ListNode{}
		p = p.Next
		n /= 10
	}
	p.Val = n
	return l
}

func getIntegerFromListNode(l *ListNode, depth float64) int {
	if l.Next == nil {
		return (int(math.Pow(10, depth)) * l.Val)
	}
	return (int(math.Pow(10, depth)) * l.Val) + getIntegerFromListNode(l.Next, depth + 1)
}

func twoSum(l1 *ListNode, l2 *ListNode) *ListNode {
	return getListNodeFromInteger(getIntegerFromListNode(l1, 0) + getIntegerFromListNode(l2, 0))
}

func main() {
	l1 := &ListNode{Val: 7 }
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next = &ListNode{Val: 3}
	l2 := &ListNode{Val:5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	twoSum(l1, l2).print()
}
