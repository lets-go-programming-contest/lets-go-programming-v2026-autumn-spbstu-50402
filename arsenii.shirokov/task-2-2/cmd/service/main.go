package main

import (
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
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

func getPreferredDish(h *MaxHeap, preference int) int

func scanPreferences(numberOfDishes int) (*MaxHeap, error)

func main() {
	var numberOfDishes int
	if _, err := fmt.Scanln(&numberOfDishes); err != nil {
		fmt.Printf("error when reading the number of dishes: %v\n", err)

		return
	}

	h, err := scanPreferences(numberOfDishes)
	if err != nil {
		fmt.Printf("error when reading the all preferences: %v\n", err)

		return
	}

	var preference int
	if _, err := fmt.Scanln(&preference); err != nil {
		fmt.Printf("error when reading the preference: %v\n", err)

		return
	}

	result := getPreferredDish(h, preference)

	fmt.Println(result)
}
