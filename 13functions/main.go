package main

import "fmt"

func main() {
	greet()

	result := add(3, 4)
	fmt.Println(result)

	res, msg := proadder(2, 5, 6, 8, 9)
	fmt.Println(res, msg)
}
func proadder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hello"
}
func add(a int, b int) int {
	return a + b
}
func greet() {
	fmt.Println("Namaste")
}
