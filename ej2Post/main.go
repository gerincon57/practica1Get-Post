package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type persona struct {
	Nombre   string `json:name`
	Apellido string `json:apellido`
}

func main() {
	metodoPost()
}

func metodoPost() {

	router := gin.Default()
	// http://localhost:8080/saludo
	router.POST("/saludo", func(c *gin.Context) {
		var p persona

		err := c.BindJSON(&p)

		if err != nil {
			fmt.Println(err)
		}
		respuesta := "Hola " + p.Nombre + " " + p.Apellido
		c.String(200, respuesta)

	})

	router.Run(":8080")
}
