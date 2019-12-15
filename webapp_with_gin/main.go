package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
)


var router *gin.Engine

func main() {
	router = gin.Default()
	// nacte sablony "jednou provzdy"
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()

			// route handler
	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"},)
	//})	// set status 200,    use template, pass data that the page uses
	router.Run()
}
