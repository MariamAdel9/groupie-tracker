package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type ErrorPage struct {
	Message string
}

func RenderError(w http.ResponseWriter, message string, status int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Error Error Page", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	tmpl.Execute(w, ErrorPage{Message: message})
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	RenderError(w, message, status)
}

func handler(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound, "Page Not Found")
		return
	}

	//the api urls
	urlname := "https://groupietrackers.herokuapp.com/api/artists"
	//fetching the data using the functions defiend in the functions.go file
	namedata, err := DataName(urlname)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}


	// Define the data to be passed to the HTML template
	pageData := PageData{
		Names:        namedata,
	}

	// Parse and execute the HTML template
	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	if err := tmpl.Execute(w, pageData); err != nil {
		fmt.Fprintf(w, "Error executing template: %v", err)
	}



}

func detailHandler(w http.ResponseWriter, r *http.Request) {
	// The API URLs
	urlname := "https://groupietrackers.herokuapp.com/api/artists"
	urldatelocation := "https://groupietrackers.herokuapp.com/api/relation"

	// Fetching the data using the functions defined in the functions.go file
	namedata, err := DataName(urlname)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}



	Datedatelocation, err := DataRelation(urldatelocation)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	// Extract the ID from the URL (e.g., /detail/1)
	idStr := strings.TrimPrefix(r.URL.Path, "/detail/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, "Invalid ID")

		return
	}

	// Now find the item corresponding to the ID
	var names NameResponse
	for _, d := range namedata {
		if d.ID == id {
			names = d
			break
		}
	}

	// If item is not found, return 404
	if names.ID == 0 {
		ErrorHandler(w, r, http.StatusNotFound, "Artist not found")
		return
	}
	var datewlocation RelationData
	for _, d := range Datedatelocation.Index {
		if d.ID == id {
			datewlocation = d
			break
		}
	}

	if datewlocation.ID == 0 {
		ErrorHandler(w, r, http.StatusNotFound, "Date/Location not found")
	}

	// Define the data to be passed to the HTML template
	details := Details{

		DateWLocation: []RelationData{datewlocation},
		Names:         names,
	}

	// Display the details of the specific item
	tmpl, err := template.ParseFiles("Templates/detail.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Error Loading Template")
		return
	}

	// Pass the detailed item data to the template
	if err := tmpl.Execute(w, details); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Error Executing Template")
	}



}
