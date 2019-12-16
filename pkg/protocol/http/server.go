package main

import (
	// "json"

	"log"

	"github.com/gin-gonic/gin"

	"github.com/atreyagaurav/quotes/pkg/service"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/all", service.All)
	r.GET("/random", service.Random)
	r.GET("/genre/:reqgenre", service.Genre)
	log.Println("Starting Server..")
	r.Run(":8003")
	log.Println("Stoping Server..")
}
