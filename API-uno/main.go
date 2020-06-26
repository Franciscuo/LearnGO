package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person is a person
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address is a address
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func getPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func getPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func createPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func deletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			json.NewEncoder(w).Encode(people)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {

	router := mux.NewRouter()

	people = append(people, Person{
		ID:        "1",
		FirstName: "Diana",
		LastName:  "Delgado",
		Address: &Address{
			City:  "Bogota",
			State: "Kennedy",
		},
	})

	people = append(people, Person{
		ID:        "2",
		FirstName: "Katherine",
		LastName:  "Castelblanco",
		Address: &Address{
			City:  "Bogota",
			State: "20 de julio",
		},
	})

	// endPoints
	router.HandleFunc("/people", getPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", createPersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", deletePersonEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
