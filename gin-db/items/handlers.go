package items

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

// Get item handler
func GetItem(c *gin.Context) {
	id := c.Param("id")

	// Get item from db
	i, err := getItem(id)
	if err != nil {
		if err == errItemNotFound {
			// Return 404 Not founf if item nod found in db
			c.Status(http.StatusNotFound)
			return
		}
		// Log error and return 500 for all other errors
		log.Printf("cannot select item from db: %+v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	// Return 200 OK and found item
	c.JSON(http.StatusOK, i)
}

// Create item handler
func CreateItem(c *gin.Context) {
	var req item
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "Malformed request: %v", err)
	}

	// Generate item id
	id := uuid.NewV4().String()
	// Create item in db
	err = createItem(id, req)
	if err != nil {
		// Log error and return 500 in case of db error
		log.Printf("cannot create item: %+v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// Set Location header pointing to a newly created resource
	c.Header("Location", "/items/"+id)
	c.Status(http.StatusCreated)
}
