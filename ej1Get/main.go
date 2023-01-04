package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	metodoGet()
}

func metodoGet() {

	router := gin.Default() //inicializa enrutamiento

	//http://localhost:8080/ping
	router.GET("/ping", func(c *gin.Context) {

		message := "pong"
		c.String(http.StatusOK, message)

	})

	//escucha go en ese puerto
	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}

}
