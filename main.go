package main

import (
	"employe/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	log.Println("App Is Starting")

	routes.SetupRoutes(server)
	server.Run(":9999")

}
