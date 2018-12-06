package main

import (
	"log"
	"os"
	"text/template"
)

type item struct{ Name string; Price float64}

type meal struct{
	Name string
	Items []item
}

type menu []meal

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("menu.gohtml"))
}

func main() {
	m := menu{
		meal{
			Name: "Breakfast",
			Items: []item{
				item{"Eggs and Bacon", 4.95},
				item{"Steak and eggs", 10.95},
			},

		},
		meal {
			Name: "Lunch",
			Items: []item{
				item{"Burger and Fries", 12.95},
				item{"Rubin", 11.95},
			},
		},

	}

	err := tpl.Execute(os.Stdout, m)

	if err != nil{
		log.Fatalln(err)
	}
}
