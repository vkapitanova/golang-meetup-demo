package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Application entry point
func main() {
	router := gin.Default()
	router.GET("/item", getItem)
	router.POST("/items", createItem)
	router.Run()
}

// Struct for parsing and returning JSON
type item struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// Get item handler
func getItem(c *gin.Context) {
	// Return 200 OK and predefined JSON response
	c.JSON(http.StatusOK, item{Name: "Valentina", Email: "valentina.kapitanova@zalando.fi"})
}

// Create item handler
func createItem(c *gin.Context) {
	var req item
	// Parse, validate and map json in request body to item struct
	err := c.BindJSON(&req)
	if err != nil {
		// Return 400 Bad request if validation failed
		c.String(http.StatusBadRequest, "Malformed request: %v", err)
	}

	c.Status(http.StatusAccepted)
}
