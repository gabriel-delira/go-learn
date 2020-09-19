package main

import (
	"text/template"

	"./ex1"
	"./ex2"
	"./ex3"
	"./ex4"
	"./ex5"
)

var tpl1, tpl2, tpl3, tpl4, tpl5 *template.Template

func init() {
	tpl1 = template.Must(template.ParseFiles("./ex1/tpl.gohtml"))
	tpl2 = template.Must(template.ParseFiles("./ex2/tpl.gohtml"))
	tpl3 = template.Must(template.ParseFiles("./ex3/tpl.gohtml"))
	tpl4 = template.Must(template.ParseGlob("./ex4/*.gohtml"))
	tpl5 = template.Must(template.ParseFiles("./ex5/tpl.gohtml"))
}

func main() {
	ex1.Test(tpl1)
	ex2.Test(tpl2)
	ex3.Test(tpl3)
	ex4.Test(tpl4)
	ex5.Test(tpl5)
}
