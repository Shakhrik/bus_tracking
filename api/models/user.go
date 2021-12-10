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
	BusID    *int64 `json:"bus_id" db:"bus_id"`
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
	BusName  string  `json:"bus_name" db:"bus_name"`
}

type UserBus struct {
	BusID           string `json:"bus_id" db:"bus_id"`
	BusName         string `json:"bus_name" db:"bus_name"`
	DestinationName string `json:"destination_name" db:"destination_name"`
	ReservationID   int64  `json:"reservation_id" db:"reservation_id"`
	SeatNumber      int64  `json:"seat_number" db:"seat_number"`
	BookedDate      string `json:"booked_date" db:"booked_date"`
	StartTime       string `json:"start_time" db:"start_time"`
	EndTime         string `json:"end_time" db:"end_time"`
}

type UserBuses struct {
	Buses []UserBus `json:"buses"`
	Count int64     `json:"count"`
}
