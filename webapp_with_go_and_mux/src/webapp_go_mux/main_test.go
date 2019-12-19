package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// request bez body
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// target http requestu - "pseudo-browser"
	recorder := httptest.NewRecorder()

	// handler requestu
	hf := http.HandlerFunc(handler)

	// serve http request - provede se tu ten handler
	hf.ServeHTTP(recorder, req)

	// test status kodu
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler vratil napicu status: %v misto: %v", status, http.StatusOK)
	}

	// test responsu
	expected := "Nazdar cype!!"
	// bez muxu by bylo ioutil.ReadAll
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler vratil nahovno body: %v misto %v", actual, expected)
	}
}
