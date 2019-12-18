package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", showIndexPage)

	// vytvoreni requestu na vyse uvedenou routu
	req, _ := http.NewRequest("GET", "/", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// test statusu 200
		statusOK := w.Code == http.StatusOK
		// test nazvu stranky "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0
		return statusOK && pageOK
	})
}

func TestArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("article/view/:article_id", getArticle)

	// vytvoreni requestu na vyse uvedenou routu
	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// test statusu 200
		statusOK := w.Code == http.StatusOK
		// test nazvu stranky "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Prvni clanek</title>") > 0
		return statusOK && pageOK
	})
}

