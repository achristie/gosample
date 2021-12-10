package main

import (
	"encoding/json"
	"net/http"
)

type PersonStore interface {
	// Add(person Person) error
	GetAll() []Person
}

type PersonServer struct {
	store PersonStore
	http.Handler
}

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Weight int    `json:"weight"`
}

func NewPersonServer(store PersonStore) *PersonServer {
	p := new(PersonServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/data", http.HandlerFunc(p.DataHandler))

	p.Handler = router
	return p
}

func (p *PersonServer) DataHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.addPerson(w, r)
	case http.MethodGet:
		p.getPeople(w)
	}
}

func (p *PersonServer) addPerson(w http.ResponseWriter, r *http.Request) {

}

func (p *PersonServer) getPeople(w http.ResponseWriter) {
	data := p.store.GetAll()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// func (p *PersonServer) AddPerson(w http.ResponseWriter, r *http.Request) {
// 	var newPerson Person
// 	err := json.NewDecoder(r.Body).Decode(&newPerson)

// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("could not decode into a person: %v", err), http.StatusBadRequest)
// 	}
// }
