package ex2

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type person struct {
	FirstName, LastName string
	Age                 int
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Default Page")
}

func dogHandler(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("./ex2/templates/dog.gohtml"))
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}
}

func meHandler(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("./ex2/templates/me.gohtml"))

	me := person{FirstName: "Gabriel Rodrigues", LastName: "Lira", Age: 27}

	err := tpl.Execute(res, me)
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}
}

// Server - Creates a server on port 8082
func Server() {

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/me/", meHandler)

	http.ListenAndServe(":8082", nil)
}
