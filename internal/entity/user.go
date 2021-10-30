package entity

type User struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" db:"name" binding:"required"`
	Email string `json:"email" db:"email" binding:"required"`
}
