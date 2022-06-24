package main

import (
	"html/template"
	"log"
	"os"
)

type item struct {
	Name  string
	Price float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	m := menu{
		meal{
			Meal: "Breakfast",
			Item: []item{
				item{
					Name:  "Oatmeal",
					Price: 4.05,
				},
				item{
					Name:  "Ommlette",
					Price: 2.0,
				},
				item{
					Name:  "Salad",
					Price: 1.5,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Item: []item{
				item{
					Name:  "Hamburger",
					Price: 4.05,
				},
				item{
					Name:  "Cheese Melted Sandwich",
					Price: 2.0,
				},
				item{
					Name:  "French Fries",
					Price: 1.5,
				},
			},
		},
		meal{
			Meal: "Dinner",
			Item: []item{
				item{
					Name:  "Pasta Bolognese",
					Price: 7.95,
				},
				item{
					Name:  "Steak",
					Price: 2.0,
				},
				item{
					Name:  "French Fries",
					Price: 13.95,
				},
			},
		},
	}
	if err := tpl.Execute(os.Stdout, m); err != nil {
		log.Fatalln(err)
	}

}
