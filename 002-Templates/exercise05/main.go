package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

type stock struct {
	Open string
	Close string
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("csv.gohtml"))
}


func main() {

	// open the file
	f, err := os.Open("table.csv")
	if err != nil{
		log.Fatalln(err)
	}
	// dont forget to close
	defer f.Chdir()


	// Read file into a var

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil{
		log.Fatalln(err)
	}

	var market []stock
	for i, line := range lines{
		if i == 0{
			continue
		}
		market = append(market, stock{line[1], line[4],})
	}

	err = tpl.Execute(os.Stdout, market)
	if err != nil {
		log.Fatalln(err)
	}
}
