package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// container hold all templates that you passed
var tpl *template.Template

// init function belong to package block it will call the function once when program executes eg. init function can be used for initialize a container
func init() {
	fmt.Println("Print is called from init function")
	tpl = template.Must(template.ParseGlob("templates/*"))
	fmt.Println("Second print is called")
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
