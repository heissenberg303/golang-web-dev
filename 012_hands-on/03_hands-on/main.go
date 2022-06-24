package main

import (
	"html/template"
	"log"
	"os"
)

type Hotel struct {
	Name     string
	Location Address
}

type Address struct {
	City   string
	Zip    int
	Region string
}

var tpl *template.Template

// initialize template
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	hotels := []Hotel{
		Hotel{Name: "LA Hotel",
			Location: Address{"LA", 10400, "Southern"},
		},
		Hotel{Name: "Golden Bridge Hotel",
			Location: Address{"San Fransisco", 10600, "Central"},
		},
		Hotel{Name: "North Cali Hotel",
			Location: Address{"San Jose", 10800, "Northen"},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}

}
