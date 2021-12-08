package models

type BusCreate struct {
	Name          string `json:"name" db:"name"`
	DestinationID int64  `json:"destination_id" db:"destination_id"`
	StartTime     string `json:"start_time" db:"start_time"`
	EndTime       string `json:"end_time" db:"end_time"`
	SeatCount     int64  `json:"seat_count" db:"seat_count"`
}

type Bus struct {
	ID              int64   `json:"id" db:"id"`
	IsFull          bool    `json:"is_full" db:"is_full"`
	RemainingSeats  int64   `json:"remaining_seats" db:"remaining_seats"`
	Name            string  `json:"name" db:"name"`
	StartTime       string  `json:"start_time" db:"start_time"`
	EndTime         string  `json:"end_time" db:"end_time"`
	SeatCount       int64   `json:"seat_count" db:"seat_count"`
	DestinationName string  `json:"destination_name" db:"destination_name"`
	BookedSeats     []int32 `json:"booked_seats" db:"booked_seats"`
}

type Buses struct {
	Buses []Bus `json:"buses" db:"buses"`
	Count int64 `json:"count" db:"count"`
}

type BriefBus struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type BriefBuses struct {
	Buses []BriefBus `json:"buses" db:"buses"`
	Count int64      `json:"count"`
}

type ChangeStatus struct {
	BusID     int64 `json:"bus_id" swaggerignore:"true"`
	BusStopID int64 `json:"bus_stop_id"`
}

type GetBusStop struct {
	BusStopName     string `json:"bus_stop_name" db:"bus_stop_name"`
	BusStopID       int64  `json:"bus_stop_id" db:"bus_stop_id"`
	IsHere          bool   `json:"is_here" db:"is_here"`
	BusStopDistance int64  `json:"bus_stop_distance" db:"bus_stop_distance"`
	DestinationName string `json:"destination_name" db:"destination_name"`
}

type GetAllBusStopes struct {
	BusStops []GetBusStop `json:"bus_stops"`
	Count    int64        `json:"count"`
}

type ReserveBus struct {
	UserID     int64 `json:"user_id" db:"user_id"`
	BusID      int64 `json:"bus_id" db:"bus_id"`
	SeatNumber int64 `json:"seat_number" db:"seat_number"`
}
