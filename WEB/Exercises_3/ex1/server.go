package ex1

import (
	"io"
	"net/http"
)

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Our server is running")
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is our dog handler")
}

func meHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Here goes 'your name'")
}

func Server() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/me/", meHandler)
	http.ListenAndServe(":8081", nil)
}
