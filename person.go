package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PersonService interface {
	Add(person Person) error
}

type PersonServer struct {
	service PersonService
}

func (p *PersonServer) AddPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)

	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode into a person: %v", err), http.StatusBadRequest)
	}
}
