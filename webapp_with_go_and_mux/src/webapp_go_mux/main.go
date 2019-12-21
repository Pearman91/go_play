// vse s timhle nazvem balicku uvidi vse z ostatnich skriptu
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
	// bez muxu: http.ListenAndServe(":8080", nil)
	// nil na miste handleru -> je pouzit defaultni DefaultServeMux
}

// abychom mohli testovat router vne mainu
func newRouter() *mux.Router {
	r := mux.NewRouter()

	// pres net/http: http.HandleFunc("/", handler)
	// ale ux umoznuje definovat metody
	r.HandleFunc("/hello", handler).Methods("GET")
	
	// definice cesty ke static souborum
	staticFileDirectory := http.Dir("./assets/")

	// routovani reuestu k jejich handlerum
	// StripPrefix proto, ze nechceme resit /assets/ - PathPrefix slouzi k matchovani
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")	

	r.HandleFunc("/band", getBandHandler).Methods("GET")
	r.HandleFunc("/band", createBandHandler).Methods("POST")

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nazdar cype!!")
}
