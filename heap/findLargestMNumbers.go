package main

import (
	"fmt"
)


type Heap[T any] struct {
	data []T
	comp func(a, b T) bool
}

func NewHeap[T any](comp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),
		comp:  comp,
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.Size() == 0 {
		var val T
		return val, false
	}
	return h.data[0], true
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.heapifyUp((len(h.data) - 1))
}

func (h *Heap[T]) heapifyUp(i int) {
	for h.comp(h.data[parentIndex(i)], h.data[i]) {
		h.swap(i, parentIndex(i))
		i = parentIndex(i)
	}
}

func lefChildIndex(i int) int {
	return 2*i + 1
}

func rightChildIndex(i int) int {
	return 2*i + 2
}

func parentIndex(i int) int {
	return (i - 1) / 2
}



func findLargetM(nums []int, m int) []int {
	// use max-heap
	heap := NewHeap(func(a, b int) bool { return a > b })
	for _, v := range nums {
		heap.Push(v)
	}
	fmt.Println(heap)
	return heap.data
}

func main() {
	findLargestM([]int{1,5,4,2,3}, 3)
}	
