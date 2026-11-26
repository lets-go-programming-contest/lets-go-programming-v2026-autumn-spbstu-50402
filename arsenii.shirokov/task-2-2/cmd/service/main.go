package main

import (
	"container/heap"
	"errors"
	"fmt"

	"choice-of-dish/internal/maxheap"
)

var (
	errNotIntegersInHeap    = errors.New("expects a container with integers")
	errPreferenceOutOfRange = errors.New("preference out of range")
)

func getPreferredDish(maxHeap *maxheap.MaxHeap, preference int) (int, error) {
	if preference <= 0 || preference > maxHeap.Len() {
		return 0, errPreferenceOutOfRange
	}

	for range preference - 1 {
		heap.Pop(maxHeap)
	}

	result, ok := heap.Pop(maxHeap).(int)
	if !ok {
		return 0, errNotIntegersInHeap
	}

	return result, nil
}

func scanPreferences(numberOfDishes int) (*maxheap.MaxHeap, error) {
	maxHeap := &maxheap.MaxHeap{}
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
