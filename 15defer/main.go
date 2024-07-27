package main

import "fmt"

func main() {

	//defer statement are executed at the end of the function befor return , executed in LIFO style
	defer fmt.Println("world")
	defer fmt.Println("hi")
	fmt.Println("hello")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
