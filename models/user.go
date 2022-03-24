package models

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateUserInput struct {
	ID       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
