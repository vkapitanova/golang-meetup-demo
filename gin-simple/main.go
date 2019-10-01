package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Application entry point
func main() {
	// Create router (with default logging and recovery)
	router := gin.Default()
	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World With GIN!")
	})
	// Run web server
	router.Run()
}
