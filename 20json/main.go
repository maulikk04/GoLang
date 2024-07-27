package main

import (
	"encoding/json"
	"fmt"
)

// Name will be displayed as name in json

type course struct {
	Name     string `json:"name"` //Name will be displayed as name in json
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` //password will not be reflected whoever is consuming this api
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React", 299, "learn.in", "abc123", []string{"js", "frontend", "web-dev"}},
		{"Node", 399, "learn.in", "bcd123", []string{"js", "backend", "web-dev"}},
		{"Python", 199, "learn.in", "abc123", nil},
	}
	//package this data as json data
	finaljson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finaljson)

}

//decode json

func DecodeJson() {
	jsonDataFromWeb := []byte(`{
                "name": "Node",
                "Price": 399,
                "website": "learn.in",     
                "tags": [
                        "js",
                        "backend",
                        "web-dev"
                ]
        }`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was vaild")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	//some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("%v, %v ,type :%T\n", key, value, value)
	}
}
