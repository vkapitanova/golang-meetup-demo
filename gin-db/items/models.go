package items

// It's overall not a good idea to combine API and data models, but I'll do it here for the sake of simplicity
type item struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" binding:"required" db:"name"`
	Email string `json:"email" binding:"required,email" db:"email"`
}
