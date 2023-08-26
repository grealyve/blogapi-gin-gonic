package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grealyve/blog-post-api/routes"
)

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.BlogRoute(router)

	router.Run("localhost:" + port)

}
