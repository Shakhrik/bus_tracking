package models

type User struct {
	ID    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Login string `json:"login" db:"login"`
}

type UserCreateRequest struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
