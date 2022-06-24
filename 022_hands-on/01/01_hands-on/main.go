package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is index page")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is dog page")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("temp.gohtml"))
	err := tpl.ExecuteTemplate(w, "temp.gohtml", "FUSE")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/me/", http.HandlerFunc(me))
	http.Handle("/dog/", http.HandlerFunc(dog))

	http.ListenAndServe(":5000", nil)

}
