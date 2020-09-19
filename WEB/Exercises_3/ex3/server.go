package ex3

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
	tpl := template.Must(template.ParseFiles("./ex3/templates/dog.gohtml"))
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}
}

func meHandler(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("./ex3/templates/me.gohtml"))

	me := person{FirstName: "Gabriel Rodrigues", LastName: "Lira", Age: 27}

	err := tpl.Execute(res, me)
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}
}

type myHandler int

func (mh myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "UHUUUU this is my onw handler !!!")
}

func hardHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is another way to use a Handler")
}

// Server - Creates a server on port 8083
func Server() {
	var mh myHandler

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/me/", meHandler)
	http.Handle("/my/handler/", mh)
	http.Handle("/hardway/", http.HandlerFunc(hardHandler))

	http.ListenAndServe(":8083", nil)
}
