package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var (
		departments int
		employees   int
		temperature int
		operation   string
	)

	if _, err := fmt.Fscan(reader, &departments); err != nil {
		return
	}

	for department := 0; department < departments; department++ {
		if _, err := fmt.Fscan(reader, &employees); err != nil {
			return
		}

		for employee := 0; employee < employees; employee++ {
			if _, err := fmt.Fscan(reader, &operation, &temperature); err != nil {
				return
			}
		}
	}
}
