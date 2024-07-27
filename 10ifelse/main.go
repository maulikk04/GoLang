package main

import "fmt"

func main() {
	logincnt := 23
	var result string
	if logincnt < 10 {
		result = "regular user"
	} else {
		result = "premium user"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("num less than 10")
	} else {
		fmt.Println("num greater than 10")
	}
}
