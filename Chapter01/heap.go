package main

import (

	"fmt"
	"container/heap"
)

type IntegerHeap []int 

// IntegerHeap method gets the length of integerHeap

func (iheap IntegerHeap) Len() int {


	return len(iheap)
}

// IntegerHeap method -checks if element of i index is less the j index

func (iheap IntegerHeap) Less(i, j int) bool {

	return iheap[i] < iheap[j]
}

// IntegerHeap method - swap the element of i to j index

func (iheap IntegerHeap) Swap(i, j int) {

	iheap[i], iheap[j] = iheap[j], iheap[i]
}

// IntegerHeap has a Push method that pushes the item with the interface

func (iheap *IntegerHeap) Push(heapinf interface{}) {

	*iheap = append(*iheap, heapinf.(int))
}

// IntegerHeap method -pops the item from the heap

func (iheap *IntegerHeap) Pop() interface{} {

	var length int 
	var lastElement int 

	var previous IntegerHeap = *iheap 
	length = len(previous)
	lastElement = previous[length-1]

	*iheap = previous[0 : length - 1]

	return lastElement
}

func main() {

	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Pop(intHeap, 2)
	fmt.Println(intHeap)
}