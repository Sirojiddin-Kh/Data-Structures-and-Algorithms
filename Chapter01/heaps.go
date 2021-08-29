package main

import (

	"fmt"
)


// MaxHeap struct has a slice that holds the array

type MaxHeap struct {

	array []int
}

// Insert method adds an element to the heap

func (h *MaxHeap) Insert(key int) {

	h.array = append(h.array, key)
	h.maxHeapifUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap

func (h *MaxHeap) Extract() int {

	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {

		fmt.Println("Cannot be extracted because array length is 0")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapiyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifUp(index int) {

	for h.array[parent(index)] < h.array[index] {

		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifDown will heapify from top to botton

func (h *MaxHeap) maxHeapiyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0 

	// loop while index has at least one child

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {

			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {

			return
		}
	}
}


// get the parent index

func parent(i int) int {

	return (i -1) / 2
}

// get the left child index 

func left(i int) int {

	return 2 * i + 1
}

// get the right child index

func right(i int) int {

	return 2 * i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {

	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {

	m :=&MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	for _, value := range buildHeap {
		m.Insert(value)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
