package main

import (
	"log"

	"github.com/gin-gonic/gin"

	config "go-gin-api/config"
	routes "go-gin-api/routes"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
