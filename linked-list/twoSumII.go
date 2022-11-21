package main

import (
	"fmt"
	"strconv"
	"math"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumber(l1, l2 *ListNode) *ListNode {
	return listNodeFromInteger(integerFromListNode(l1) + integerFromListNode(l2))
}

func integerFromListNode(l *ListNode) int {
	var n int
	count := 0
	size := l.getSize()
	for l != nil {
		n += (l.Val * int(math.Pow(10, float64(size - count))))
		l = l.Next
		count++
	}
	return n
}

func (head *ListNode) getSize() int {
	var count int
	for head.Next != nil {
		count++
		head = head.Next
	}
	return count
}

func listNodeFromInteger(n int) *ListNode {
	// get reminder of dividing n by 10
	// create ListNode
	tail := &ListNode{ Val: n % 10 }
	var curr *ListNode
	n /= 10
	// repeat and set new ListNode's next value to the previously made
	for n >= 10 {
		curr = &ListNode{ Val: n % 10 }
		curr.Next = tail
		tail = curr
		n /= 10
	}
	curr = &ListNode{ Val: n }
	curr.Next = tail
	return curr
}

func (head *ListNode) insert(l *ListNode) {
	if head.Next == nil {
		head.Next = l
		return
	}
	head.Next.insert(l)
}

func (head *ListNode) String() string {
	var s string
	for head.Next != nil {
		s += strconv.Itoa(head.Val) + "->"
		head = head.Next
	}
	return s + strconv.Itoa(head.Val)
}

func main() {
	l := &ListNode{ Val: 7 }
	l.insert(&ListNode{ Val: 2 })
	l.insert(&ListNode{ Val: 4 })
	l.insert(&ListNode{ Val: 3 })
	fmt.Println(l)
	l2 := &ListNode{ Val: 5 }
	l2.insert(&ListNode{ Val: 6 })
	l2.insert(&ListNode{ Val: 4})
	fmt.Printf("RESULT: %v", addTwoNumber(l, l2))
}
