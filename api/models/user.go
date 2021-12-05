package models

type User struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
	BusName  string `json:"bus_name" db:"bus_name"`
	UserType string `json:"user_type" db:"alias"`
}

type UserCreateRequest struct {
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
	BusID    int64  `json:"bus_id" db:"bus_id"`
	Alias    string `json:"alias" db:"alias"`
}

type AllUsers struct {
	Users []User `json:"users"`
	Count int64  `json:"count"`
}

type UserLogin struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type UserInfo struct {
	ID       int64   `json:"id" db:"id"`
	UserType string  `json:"user_type" db:"user_type"`
	BusID    *string `json:"bus_id" db:"bus_id"`
}
