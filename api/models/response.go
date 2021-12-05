package models

type Empty struct{}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SuccessModel ...
type SuccessModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorModel ...
type ErrorModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type ResponseWithID struct {
	ID int64 `json:"id" db:"id"`
}
