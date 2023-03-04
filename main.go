package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

func SendMSG(c *gin.Context) {
	var person Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, person)
}

func main() {
	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, "SalikhovR")
	})
	g.POST("/msg", SendMSG)
	err := g.Run("localhost:8008")
	if err != nil {
		log.Fatal(err)
		return
	}
}
