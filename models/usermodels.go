package models


type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Todos struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Date     string `json:"date"`
}
