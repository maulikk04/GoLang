package main

import "fmt"

func main() {

	//days := []string{"Sun", "Mon", "Tue", "Wed", "Thu"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for ind, day := range days {
	// 	fmt.Printf("ind : %v ,val : %v\n", ind, day)
	// }

	value := 1
	for value < 10 {

		if value == 2 {
			goto hello
		}
		if value == 5 {
			break
		}
		fmt.Println(value)
		value++
	}

hello:
	fmt.Println("hello")
}
