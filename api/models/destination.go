package models

type Destination struct {
	ID       int64  `json:"id" db:"id"`
	From     string `json:"from" db:"from_place"`
	To       string `json:"to" db:"to_place"`
	Price    int64  `json:"price" db:"price"`
	Distance int64  `json:"distance" db:"distance"`
}
type Destinations struct {
	Destinations []Destination `json:"destinations"`
	Count        int64         `json:"count"`
}
type DestinationCreate struct {
	From     string `json:"from" db:"from_place"`
	To       string `json:"to" db:"to_place"`
	Price    int64  `json:"price" db:"price"`
	Distance int64  `json:"distance" db:"distance"`
}

type DestinationUpdate struct {
	ID int64 `json:"id" db:"id" swaggerignore:"true"`
	DestinationCreate
}

type DestinationGet struct {
	ID int64 `json:"id" db:"id"`
}
