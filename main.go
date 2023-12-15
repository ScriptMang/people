package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID    string `json:"id"`
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Age   int    `json:"age"`
}

var persons = []person{
	{ID: "1", Fname: "Bob", Lname: "Ross", Age: 50},
	{ID: "2", Fname: "Kongary", Lname: "Kat", Age: 34},
	{ID: "3", Fname: "LimeStone", Lname: "Rose", Age: 45},
}

func getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, persons)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("docs/*.html")

	router.Static("/css", "./docs/css/")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})
	router.GET("/persons", getPersons)
	router.Run()
}
