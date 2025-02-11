package main

import (
	"encoding/json"
	"net/http"
)

//1- create struct for each json data that you will import with the following structure
// data type jason name

//json response from the name link

type NameResponse struct {
	Name       string   `json:"name"`
	ID         int      `json:"id"`
	Image      string   `json:"image"`
	Members    []string `json:"members"`
	FirstAlbum string   `json:"firstAlbum"`
	Create     int      `json:"creationDate"`
}

//since the json data a is wrapped in an array called index then define two structures, the first definging the index array

type LocationResponse struct {
	Index []LocationData `json:"index"`
}

// and the second definging	 the elements inside the index array
// 1- the structure of the locaiton data
type LocationData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// 2- the structure of the date data
type DateResponse struct {
	IndexDate []DateData `json:"index"`
}

type DateData struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// 3- the structure of the relationship data
type RelationData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationshipResponse struct {
	Index []RelationData `json:"index"`
}

// combine all the data togather into one structure
type PageData struct {
	Names        []NameResponse
}

type Details struct {
	DateWLocation     []RelationData
	Names     NameResponse
}

//define a get http request functionn to get the json data and store it in

func DataName(url string) ([]NameResponse, error) {

	//get data by the url
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	//create a variable data of type strture
	var data []NameResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	//return the data
	return data, nil
}

func Locations(url string) (*LocationResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data LocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func Datadate(url string) (*DateResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data DateResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func DataRelation(url string) (*RelationshipResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data RelationshipResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
