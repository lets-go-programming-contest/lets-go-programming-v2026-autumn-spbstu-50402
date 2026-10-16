package main

import "fmt"

func main() {

	var (
		n, k, temp int
		str        string
	)

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("invalid department amount")

		return
	}

	for i := 0; i < n; i++ {
		t1 := 15
		t2 := 30

		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Println("invalid employee amount")
			return
		}

		for j := 0; j < k; j++ {
			_, err = fmt.Scan(&str)
			if err != nil {
				fmt.Println("invalid input")
				return
			}

			_, err = fmt.Scan(&temp)
			if err != nil {
				fmt.Println("invalid input")
				return
			}

			switch str {
			case ">=":
				t1 = max(t1, temp)
			case "<=":
				t2 = min(t2, temp)
			default:
				fmt.Println("invalid input")
			}

			if t2 < t1 {
				fmt.Println("-1")
			} else {
				fmt.Println(t1)
			}
		}
	}
}
