package main

import (
	"errors"
	"fmt"

	"arsenii.shirokov/task-2-2/internal/maxheap"
)

var errPreferenceOutOfRange = errors.New("preference out of range")

func getPreferredDish(maxHeap maxheap.IntMaxHeap, preference int) (int, error) {
	if preference <= 0 || preference > maxHeap.Len() {
		return 0, errPreferenceOutOfRange
	}

	for range preference - 1 {
		if _, err := maxHeap.Pop(); err != nil {
			return 0, fmt.Errorf("error when getting the preferred dish: %w", err)
		}
	}

	result, err := maxHeap.Pop()
	if err != nil {
		return 0, fmt.Errorf("error when getting the preferred dish: %w", err)
	}

	return result, nil
}

func scanPreferences(numberOfDishes int) (maxheap.IntMaxHeap, error) {
	maxHeap := maxheap.NewIntMaxHeap()

	for range numberOfDishes {
		var priority int
		if _, err := fmt.Scan(&priority); err != nil {
			return maxHeap, fmt.Errorf("error while reading the priority: %w", err)
		}

		maxHeap.Push(priority)
	}

	return maxHeap, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("unexpected error (panic): %v\n", r)
		}
	}()

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
