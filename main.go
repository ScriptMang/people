package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPersons(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/docs/persons.json")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("docs/*.html")

	router.Static("/docs", "./docs")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})
	router.GET("/persons", getPersons)
	router.Run()
}
