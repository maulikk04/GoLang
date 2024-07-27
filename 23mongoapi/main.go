// maulikkansara04
// ZD9pRMpwbdOuowUk
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maulikk04/mongoapi/router"
)

func main() {
	fmt.Println("MongoAPI")
	r := router.Router()
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000")
}
