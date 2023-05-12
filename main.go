package main

import (
	"fmt"
)

// MaxHeap struct has a slice that holds the array

type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.MaxHeapifyUp(len(h.array) - 1)
}

//Extract returns the largest key, and removes it from the heap

//	Parent Index
func Parent(i int) int {
	return (i - 1) / 2
}

// Left Index
func left(i int) int {
	return (2 * i) + 1
}

//	Right Index
func right(i int) int {
	return (2 * i) + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// Extract.

func (h *MaxHeap) Extract() int {
	last := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("Unable to extract. Array is empty")
		return -1
	}
	extracted := h.array[0]
	h.array[0] = h.array[last]
	h.array = h.array[:last]
	h.MaxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	//loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { //when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { //when right child is larger
			childToCompare = r
		}
		//Compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) MaxHeapifyUp(index int) {
	for h.array[Parent(index)] < h.array[index] {
		h.swap(Parent(index), index)
		index = Parent(index)
	}
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 22, 80, 9, 101, 45, 42, 44}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
