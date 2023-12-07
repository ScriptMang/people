package main

import (
	"github.com/gin-gonic/gin"
)

type person struct {
	ID string    `json:"id"`
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	age int      `json:"age"`
}

func main() {
	//boiler plate
	resp := gin.Default()
	resp.GET("/marco", func(c *gin.Context){
		c.JSON(200, gin.H{
			"msg": "polo",
		})
   })
   resp.Run()
}