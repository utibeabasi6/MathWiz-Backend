package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	)

type Person struct {
	Name string `json: name`
	Age int `json: age`
}

func helloHandler (w http.ResponseWriter, _ *http.Request){
	person := Person{Name: "Utibeabasi", Age: 17}
	err := json.NewEncoder(w).Encode(person)
	if err != nil{
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/", helloHandler)
	port := "8000"
	fmt.Printf("Starting server on port %v\n", port)
	log.Fatal(http.ListenAndServe(":" + port, r))

}
