package models

type User struct {
	ID       string `json:"id,omitempty" gorm:"primary_key"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
