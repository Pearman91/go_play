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
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nazdar cype!!")
}
