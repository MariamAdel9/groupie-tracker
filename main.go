package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/detail/", detailHandler)
	http.ListenAndServe(":8080", nil)
}
