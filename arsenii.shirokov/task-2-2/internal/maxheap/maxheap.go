package maxheap

type MaxHeap []int

func (h *MaxHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic("MaxHeap: push expects an int")
	}

	*h = append(*h, val)
}

func (h *MaxHeap) Pop() any {
	if len(*h) == 0 {
		return nil
	}

	endVal := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return endVal
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	if i >= len(*h) || j >= len(*h) {
		panic("MaxHeap: index out of range")
	}

	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	if i >= len(*h) || j >= len(*h) {
		panic("MaxHeap: index out of range")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
