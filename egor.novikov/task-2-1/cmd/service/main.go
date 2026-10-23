package main

import "fmt"

func main() {
	var departments int = 0
	_, err := fmt.Scan(&departments)

	if err != nil {
		fmt.Println("Invalid input departments")
		return
	}

	for range departments{
		var staff int = 0
		_, err = fmt.Scan(&staff)

		if err != nil{
			fmt.Println("Invalid input staff")
			return
		}

		for range staff{
			var temperature string = ""
			_, err = fmt.Scan(&temperature)
			
			if err!=nil{
				fmt.Println("Invalid input temperature")
			}
		}
	}


	
	
}
