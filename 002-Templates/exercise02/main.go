package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct{
	Name string
	Address string
	City string
	Zip string
	Region string
}

type hotels []hotel

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("hotel.gohtml"))
}


func main() {

	h := []hotel{
		hotel{"Hyatt", "123 N Main", "San Francisco", "94121", "NorCal"},
		hotel{"Le Meridian", "123 N Main", "San Francisco", "94121", "NorCal"},
		hotel{"Filmore", "123 N Main", "San Francisco", "94121", "NorCal"},
		hotel{"Park Regency", "123 N Main", "San Francisco", "94121", "NorCal"},
	}

	err := tpl.Execute(os.Stdout, h)

	if err != nil{
		log.Fatalln(err)
	}

}
