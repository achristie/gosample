package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Weight int    `json:"weight"`
}

var people []Person

func main() {
	log.Print("hello!")

	http.HandleFunc("/", index)
	http.HandleFunc("/data", getJson)

	people, _ = New("data.json")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func New(pathToFile string) ([]Person, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Println("Could not open file, ", pathToFile)
		return nil, err
	}
	bytes, _ := ioutil.ReadAll(file)
	var p []Person
	json.Unmarshal(bytes, &p)
	log.Println(p)
	return p, nil
}

func getJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
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
	err := tmpl.Execute(w, people)
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

	p := new(Person)

	p.Age, _ = strconv.Atoi(r.Form.Get("age"))
	p.Name = r.Form.Get("name")
	p.Weight, _ = strconv.Atoi(r.Form.Get("weight"))

	people = append(people, *p)
	write()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func write() {
	b, err := json.MarshalIndent(people, "", "")
	if err != nil {
		log.Fatal("unable to convert to json. shutting down")
	}
	err = os.WriteFile("data.json", b, 0644)
	if err != nil {
		log.Fatal("unable to write to disk. shutting down")
	}
}
