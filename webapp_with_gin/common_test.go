package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

func TestMain(m *testing.M) {
	// set Gin to test mode
	gin.SetMode(gin.TestMode)
	// run other tests
	os.Exit(m.Run())
}

// fce na vytvoreni routeru pro testy
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// fce na zpracovani requestu a otestovani responsu
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// vytvoreni response recorderu
	w := httptest.NewRecorder()
	//vytvoreni servisy a zpracovani requestu
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// uloz main list do tmplistu pro ucely testovani
func saveLists() {
	tmpArticleList = articleList
}

// ziskej main list z tmplistu
func restoreLists() {
	articleList = tmpArticleList
}

