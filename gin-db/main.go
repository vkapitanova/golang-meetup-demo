package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vkapitanova/golang-meetup-demo/gin-db/db"
	"github.com/vkapitanova/golang-meetup-demo/gin-db/items"
	"log"
)

// Application entry point
func main() {
	// Init db connection
	err := db.Init("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		// Stop app if cannot connect to db
		log.Fatalf("cannot connect to db: %+v", err)
	}

	router := gin.Default()
	router.GET("/items/:id", items.GetItem)
	router.POST("/items", items.CreateItem)
	router.Run()
}
