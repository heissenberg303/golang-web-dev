package main

import (
	"html/template"
	"log"
	"os"
)

type item struct {
	name  string
	price float64
}

type meal struct {
	Meal string
	Item item
}

type menu []meal

type restaurant struct {
	Name     string
	Menu     menu
	employee int
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tml.gohtml"))
}

func main() {
	m1 := menu{
		meal{
			Meal: "Breakfast",
			Item: item{
				name:  "Salad",
				price: 2.14,
			},
		},
		meal{
			Meal: "Lunch",
			Item: item{
				name:  "Steak",
				price: 2.15,
			},
		},
		meal{
			Meal: "Dinner",
			Item: item{
				name:  "Chicker",
				price: 3.00,
			},
		},
	}
	m2 := menu{
		meal{
			Meal: "Breakfast",
			Item: item{
				name:  "Orange",
				price: 2.00,
			},
		},
		meal{
			Meal: "Lunch",
			Item: item{
				name:  "Noodle",
				price: 3.00,
			},
		},
		meal{
			Meal: "Dinner",
			Item: item{
				name:  "Serial",
				price: 1.00,
			},
		},
	}
	m3 := menu{
		meal{
			Meal: "Breakfast",
			Item: item{
				name:  "Apple",
				price: 2.14,
			},
		},
		meal{
			Meal: "Lunch",
			Item: item{
				name:  "Pork",
				price: 2.15,
			},
		},
		meal{
			Meal: "Dinner",
			Item: item{
				name:  "Lamb",
				price: 3.00,
			},
		},
	}

	res := restaurants{
		restaurant{
			Name:     "res1",
			Menu:     m1,
			employee: 10,
		},
		restaurant{
			Name:     "res2",
			Menu:     m2,
			employee: 8,
		},
		restaurant{
			Name:     "res3",
			Menu:     m3,
			employee: 4,
		},
	}

	err := tpl.Execute(os.Stdout, res)
	if err != nil {
		log.Println("Something wrong in execute template")
	}
}
