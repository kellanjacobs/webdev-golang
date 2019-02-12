package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/cat/", jack)
	http.HandleFunc("/me/", about)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("Error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", "World ")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func jack(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("Error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", "Jack")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func about(w http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", "Daniel")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}