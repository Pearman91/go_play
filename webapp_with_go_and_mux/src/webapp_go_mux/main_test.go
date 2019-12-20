package main

import (
	"io/ioutil"
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

func TestRouter(t *testing.T) {
	r := newRouter()
	// rozbehnuti mock serveru
	mockServer := httptest.NewServer(r)
	// GEt request na /hello
	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}
	
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status na picu: %d misto aby byl vcajku", resp.StatusCode)
	}

	// zavri resp i kdyby se predtim neco posralo
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	
	// tet ze body je prazdne
	respString := string(b)
	expected := "Nazdar cype!!"
	if respString != expected {
		t.Errorf("Response mela byt: %s ale je: %s", expected, respString)
	}
}

func TestRouteButThereIsNoRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL + "/hello", "", nil)

	if err != nil {t.Fatal(err)}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Hele ma to vratit status 405 ale vraci to uplnou picovinu: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {t.Fatal(err)}

	respString := string(b)
	expected := ""
	if respString != expected {
		t.Errorf("Response mela byt: %s ale je: %s", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/assets")

	if err != nil {t.Fatal(err)}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ma to bejt vcjaku 200 ale vraci to uplnou picovinu: %d", resp.StatusCode)
	}

	// otestuj content-type header, at je jisto, ze to vraci html
	contentType := respHandler.Get("Content-Type")
	expectedContentType := "text/html, charset=uft8"

	if expectedContentType != contentType{
		t.Errorf("Content type mel byt: %s ale je: %s", expectedContentType, contentType)
	}
}
