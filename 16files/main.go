package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This need to go in files"

	file, err := os.Create("./myfile.txt")
	if err != nil {
		//panic will shut down the program and show the error
		panic(err)
	}

	len, err := io.WriteString(file, content)
	if err != nil {
		//panic will shut down the program and show the error
		panic(err)
	}
	fmt.Println(len)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
