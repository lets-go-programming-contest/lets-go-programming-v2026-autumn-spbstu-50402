package main

import "fmt"

func main() {
	var (
		n   int
		str string
	)

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("invalid department amount")

		return
	}

	for range n {
		var k int
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Println("invalid employee amount")

			return
		}

		t1 := 15
		t2 := 30

		for range k {
			_, err = fmt.Scan(&str)
			if err != nil {
				fmt.Println("invalid input")

				return
			}

			var temp int
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

				return
			}

			if t2 < t1 {
				fmt.Println("-1")
			} else {
				fmt.Println(t1)
			}
		}
	}
}
