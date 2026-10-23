package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	errNotIntegersInHeap    = errors.New("expects a container with integers")
	errPreferenceOutOfRange = errors.New("preference out of range")
)

type MaxHeap []int

func (h *MaxHeap) Push(x any) {
	val, isInt := x.(int)
	if !isInt {
		panic("MaxHeap: push expects an int")
	}

	*h = append(*h, val)
}

func (h *MaxHeap) Pop() any {
	endVal := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return endVal
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func getPreferredDish(maxHeap *MaxHeap, preference int) (int, error) {
	if preference <= maxHeap.Len() {
		var result int

		for range preference {
			var isInt bool
			result, isInt = heap.Pop(maxHeap).(int)

			if !isInt {
				return 0, errNotIntegersInHeap
			}
		}

		return result, nil
	}

	return 0, errPreferenceOutOfRange
}

func scanPreferences(numberOfDishes int) (*MaxHeap, error) {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for range numberOfDishes {
		var priority int
		if _, err := fmt.Scan(&priority); err != nil {
			return maxHeap, fmt.Errorf("error while reading the priority: %w", err)
		}

		heap.Push(maxHeap, priority)
	}

	return maxHeap, nil
}

func main() {
	var numberOfDishes int
	if _, err := fmt.Scanln(&numberOfDishes); err != nil {
		fmt.Printf("error when reading the number of dishes: %v\n", err)

		return
	}

	maxHeap, err := scanPreferences(numberOfDishes)
	if err != nil {
		fmt.Printf("error when reading the all preferences: %v\n", err)

		return
	}

	var preference int
	if _, err := fmt.Scanln(&preference); err != nil {
		fmt.Printf("error when reading the preference: %v\n", err)

		return
	}

	result, err := getPreferredDish(maxHeap, preference)
	if err != nil {
		fmt.Printf("error when get preferred dish: %v\n", err)

		return
	}

	fmt.Println(result)
}
