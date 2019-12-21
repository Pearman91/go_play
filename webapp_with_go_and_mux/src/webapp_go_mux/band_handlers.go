package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Band struct {
	Name string `json:"bandname"`
	Genre string `json:"genre"`
}

var bands []Band

// get request na bands
func getBandHandler(w http.ResponseWriter, r *http.Request) {
	bandListBytes, err := json.Marshal(bands)
	if err != nil {
		fmt.Print(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)  // vrat 500
		return
	}

	w.Write(bandListBytes)  // write response
}

// post request na band
func createBandHandler(w http.ResponseWriter, r *http.Request) {
	band := Band{}  // nova instance kapely
	err := r.ParseForm()
	if err != nil {
		fmt.Print(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)  // vrat 500
		return
	}

	// vyziskej z formu info o POSTovane kapele a pridej tuto kapelu
	band.Name = r.Form.Get("bandname")
	band.Genre = r.Form.Get("genre")
	bands = append(bands, band)

	// presmeruj na "/assets/"
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
