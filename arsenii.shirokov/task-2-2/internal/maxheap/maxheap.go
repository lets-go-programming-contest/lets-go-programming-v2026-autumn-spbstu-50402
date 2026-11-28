package maxheap

import (
	"container/heap"
	"errors"
)

var (
	errNotIntegersInHeap = errors.New("expects a container with integers")
	errEmptyHeap         = errors.New("maxheap: heap is empty")
)

type maxHeap []int

func (h *maxHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic("MaxHeap: push expects an int")
	}

	*h = append(*h, val)
}

func (h *maxHeap) Pop() any {
	if len(*h) == 0 {
		return nil
	}

	endVal := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return endVal
}

func (h *maxHeap) Len() int {
	return len(*h)
}

func (h *maxHeap) Less(i, j int) bool {
	if i >= len(*h) || j >= len(*h) || i < 0 || j < 0 {
		panic("MaxHeap: index out of range")
	}

	return (*h)[i] > (*h)[j]
}

func (h *maxHeap) Swap(i, j int) {
	if i >= len(*h) || j >= len(*h) || i < 0 || j < 0 {
		panic("MaxHeap: index out of range")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

type IntMaxHeap struct {
	h maxHeap
}

func NewIntMaxHeap() IntMaxHeap {
	h := maxHeap{}
	heap.Init(&h)

	return IntMaxHeap{
		h: h,
	}
}

func (mh *IntMaxHeap) Push(x int) {
	heap.Push(&mh.h, x)
}

func (mh *IntMaxHeap) Pop() (int, error) {
	if mh.h.Len() == 0 {
		return 0, errEmptyHeap
	}

	val, ok := heap.Pop(&mh.h).(int)
	if !ok {
		return 0, errNotIntegersInHeap
	}

	return val, nil
}

func (mh *IntMaxHeap) Len() int {
	return mh.h.Len()
}
