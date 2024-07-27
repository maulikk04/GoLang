package main

import "fmt"

func main() {
	//no inheritance in golang , no super or parent
	jetha := User{"jetha", "jetha@gokul.com", true, 16}
	fmt.Println(jetha)
	//%+v will give in detail way
	fmt.Printf("detail : %+v\n", jetha)

	jetha.GetStatus()
	jetha.NewMail()
	fmt.Printf("detail : %+v\n", jetha)

}

// fitst letter capital access globally
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "pogo.com"
	fmt.Println(u.Email)
}
