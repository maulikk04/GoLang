package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Your name : ", input)

	age, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		age = age - 1
		fmt.Println("New age :", age)
	}
}
