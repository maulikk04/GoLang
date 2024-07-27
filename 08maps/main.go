package main

import "fmt"

func main() {
	languages := make(map[string]string) //map[key]value

	languages["JS"] = "javascript"
	languages["PHP"] = "php"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	//deletes
	delete(languages, "PHP") // delete(map name, key);
	fmt.Println(languages)

	//loop in map
	for key, value := range languages {
		fmt.Printf("key : %v, value : %v\n", key, value)
		//%v is value value
	}
}
