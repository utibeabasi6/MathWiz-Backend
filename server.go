package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	)

type Person struct {
	Name string `json: name`
	Age int `json: age`
}

func helloHandler (w http.ResponseWriter, r *http.Request){
	person := Person{Name: "Utibeabasi", Age: 17}
	err := json.NewEncoder(w).Encode(person)
	if err != nil{
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/", helloHandler)
	port := os.Getenv("PORT")
	if port == ""{
		log.Println("Port variable not set")
		log.Println("Defaulting to default port")
		port = "8080"
	}
	log.Printf("Listening on port %v:\n", port)
	log.Fatal(http.ListenAndServe(":" + port, r))

}
