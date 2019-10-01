package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Application entry point
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World With GIN!")
	})
	router.Run()
}
