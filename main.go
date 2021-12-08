package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name   string
	Age    int
	Weight int
}

var people []Person

func main() {
	log.Print("hello!")

	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleGet(w, r)
		return
	}

	handlePost(w, r)

}

func handleGet(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.tmpl")
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template", err)
		return
	}
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	p := Person{
		Name:   r.Form.Get("name"),
		Weight: r.Form.Get("weight"),
		Age:    r.Form.Get("age"),
	}

	log.Printf("%v", p)

}

/*
1) Post JSON
2) Render the JSON back
3) Save JSON to a file
4) Load JSON on INIT
5) Create a form to post data!
*/
