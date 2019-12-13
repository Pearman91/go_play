package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nazdar curaku")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")

	// kdyz spustim 'go run *.go, tak tu budu znat AllUsers. NewUser,...'
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Jezis te vidi, zmrde!")

	InitialMigration()

	handleRequests()
}
