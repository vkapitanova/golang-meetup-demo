package items

import (
	"database/sql"
	"errors"
	"github.com/vkapitanova/golang-meetup-demo/gin-db/db"
)

var errItemNotFound = errors.New("item not found")

// get item from db
func getItem(id string) (item, error) {
	var i item
	err := db.GetConn().Get(&i, "SELECT * FROM item WHERE id = $1", id)
	if err != nil {
		// return custom error when item is not found by id
		if err == sql.ErrNoRows {
			return item{}, errItemNotFound
		}
		// propagate error in other cases
		return item{}, err
	}
	return i, nil
}

// get item in db
func createItem(id string, i item) error {
	_, err := db.GetConn().Exec("INSERT INTO item(id, name, email) VALUES ($1, $2, $3)", id, i.Name, i.Email)
	return err
}
