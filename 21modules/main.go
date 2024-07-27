package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	//running as server
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("hey mod users")
}

// w for write , r for read
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hey</h1>"))
}
