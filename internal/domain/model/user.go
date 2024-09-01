package model

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       uint   `json:"age"`
	Email     string `json:"email"`
	IsActive  bool   `json:"isActive"`
}
