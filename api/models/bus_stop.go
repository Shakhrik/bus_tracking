package models

type BusStop struct {
	ID              int64  `json:"id" db:"id"`
	Name            string `json:"bus_stop" db:"name"`
	Distance        int64  `json:"distance" db:"distance"`
	DestinationName string `json:"destination_name" db:"destination_name"`
}
type BusStopCreate struct {
	Name          string `json:"name" db:"name"`
	DestinationID int64  `json:"destination_id" db:"destination_id"`
	Distance      int64  `json:"distance" db:"distance"`
}

type BusStopes struct {
	BusStopes []BusStop `json:"bus_stopes"`
	Count     int64     `json:"count" db:"count"`
}
