package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println(res.StatusCode)
	fmt.Println(res.ContentLength)

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(content))
	//another way to convert string
	var responseString strings.Builder
	bytecount, _ := responseString.Write(content) //same as contentLength
	fmt.Println(bytecount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	//dummy paload
	requestBody := strings.NewReader(`
		{
			"username" : "John",
			"password": "password"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	data := url.Values{}
	data.Add("username", "John")
	data.Add("password", "password")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	//fmt.Println(res.Body)
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
