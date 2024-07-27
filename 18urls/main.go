package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://localhost:5000/home?user=admin"

func main() {
	fmt.Println(myurl)

	//parsing - extract some info
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparam := result.Query()
	fmt.Printf("%T\n", qparam)

	fmt.Println(qparam["user"])

	for key, val := range qparam {
		fmt.Println(key, val)
	}

	//construct url

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "localhost",
		Path:    "/home",
		RawPath: "user=admin",
	}
	newUrl := partsofUrl.String()
	fmt.Println(newUrl)
}
