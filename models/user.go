package models

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
