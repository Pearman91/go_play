package main

import (
	"net/http"
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

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

