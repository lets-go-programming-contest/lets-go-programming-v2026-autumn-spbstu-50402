package main

import (
	"fmt"
)

func main() {
	var n, k int
	_, err := fmt.Scan(&n);
	if err != nil {
		fmt.Println("Invalid amount of departments")
		return
	}

	_, err = fmt.Scan(&k);
	if err != nil {
		fmt.Println("Invalid amount of staff")
		return
	}
}
