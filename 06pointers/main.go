package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println(ptr) //output <nil>

	num := 7
	var ptr *int = &num
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr + 1
	fmt.Println(num)
}
