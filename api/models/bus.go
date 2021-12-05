package models

type BusCreate struct {
	Name          string `json:"name" db:"name"`
	DestinationID int64  `json:"destination_id" db:"destination_id"`
	StartTime     string `json:"start_time" db:"start_time"`
	EndTime       string `json:"end_time" db:"end_time"`
	SeatCount     int64  `json:"seat_count" db:"seat_count"`
}

type Bus struct {
	ID              int64  `json:"id" db:"id"`
	IsFull          bool   `json:"is_full" db:"is_full"`
	RemainingSeats  int64  `json:"remaining_seats" db:"remaining_seats"`
	Name            string `json:"name" db:"name"`
	StartTime       string `json:"start_time" db:"start_time"`
	EndTime         string `json:"end_time" db:"end_time"`
	SeatCount       int64  `json:"seat_count" db:"seat_count"`
	DestinationName string `json:"destination_name" db:"destination_name"`
}

type Buses struct {
	Buses []Bus `json:"buses" db:"buses"`
	Count int64 `json:"count" db:"count"`
}

type BriefBus struct {
	ID   int64 `json:"id" db:"id"`
	Name int64 `json:"name" db:"name"`
}
type BusReserve struct {
	BusID int64 `json:"bus_id" db:"bus_id"`
}
type BriefBuses struct {
	Buses []BriefBus `json:"buses" db:"buses"`
	Count int64      `json:"count"`
}
