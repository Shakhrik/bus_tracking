package models

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
