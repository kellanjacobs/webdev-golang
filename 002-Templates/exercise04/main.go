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

type restaruant struct{
	Name string
	RMenu menu
}

type restaruants []restaruant

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("menu.gohtml"))
}

func main() {
	r := restaruants{
		{
			Name: "Squat and Gobble",
			RMenu: menu{
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

			},
		},
		{
			Name: "Ampawa",
			RMenu: menu{
				meal{
					Name: "Lunch",
					Items: []item{
						item{"Pad Thai", 14.95},
						item{"Beef Noodle Soup", 10.95},
						item{"Pad Kee Mao", 14.95},
						item{"Cucumber Salad", 4.95},
					},
				},
			},
		},
	}


	err := tpl.Execute(os.Stdout, r)

	if err != nil{
		log.Fatalln(err)
	}
}
