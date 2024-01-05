package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPersons(c *gin.Context) {
	c.HTML(http.StatusOK, "request1.tmpl", gin.H{
		"people": "persons.json",
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("docs/templates/*")

	router.Static("docs", "./docs")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"content": "This is an index page...",
		})
	})
	router.GET("/request1", getPersons)
	err := router.Run("127.0.0.1:8080")
	if err == nil {
		fmt.Printf("Route couldn't be ran, %v\n", err)
	}

}
