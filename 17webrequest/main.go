package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://google.com"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Type : %T\n", res)

	defer res.Body.Close() // caller responsiblity to close the connection

	databyte, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)
}
