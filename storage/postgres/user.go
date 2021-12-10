package postgres

import (
	"time"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/Shakhrik/inha/bus_tracking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserRepoI {
	return userRepo{db: db}
}

func (u userRepo) Create(req models.UserCreateRequest) (res models.ResponseWithID, err error) {
	query := `INSERT INTO users (name, login, password, bus_id, user_type_id) 
			  VALUES($1, $2, $3, $4, (SELECT id FROM user_type WHERE alias = $5))
			  RETURNING id`
	err = u.db.Get(&res.ID, query, req.Name, req.Login, req.Password, req.BusID, req.Alias)
	return
}

func (u userRepo) Delete(id int32) (res models.ResponseWithID, err error) {
	query := `DELETE FROM users WHERE id = $1 RETURNING id`
	err = u.db.Get(&res.ID, query, id)
	return
}

func (u userRepo) GetAll(alias string, limit, page int32) (models.AllUsers, error) {
	var count int64
	res := []models.User{}

	offset := (page - 1) * limit

	query := `SELECT u.id, u.name, u.login, b.name AS bus_name, ut.alias
			  FROM users u
			  JOIN bus b ON u.bus_id = b.id
			  JOIN user_type ut ON u.user_type_id = ut.id
			  WHERE u.user_type_id = (SELECT id FROM user_type WHERE alias = $1)
			  LIMIT $2 OFFSET $3`
	err := u.db.Select(&res, query, alias, limit, offset)
	if err != nil {
		return models.AllUsers{Users: res}, err
	}

	queryCount := `SELECT count(1)
				   FROM users u
				   JOIN bus b ON u.bus_id = b.id
				   JOIN user_type ut ON u.user_type_id = ut.id
				   WHERE u.user_type_id = (SELECT id FROM user_type WHERE alias = $1)
				   LIMIT $2 OFFSET $3`
	err = u.db.Get(&count, queryCount, alias, limit, offset)
	if err != nil {
		return models.AllUsers{Users: res}, err

	}
	return models.AllUsers{Users: res, Count: count}, nil
}

func (u userRepo) Login(req models.UserLogin) (res models.UserInfo, err error) {
	query := `SELECT u.id, ut.alias as user_type, bus_id, b.name bus_name
			  FROM users u
			  JOIN user_type ut ON ut.id = u.user_type_id
			  JOIN bus b ON u.bus_id = b.id
			  WHERE login = $1 and password = $2
			  `
	err = u.db.Get(&res, query, req.Login, req.Password)
	return
}

func (u userRepo) GetUserBuses(userID int64) (models.UserBuses, error) {
	var (
		bookedDate time.Time
		res        []models.UserBus
	)

	query := `SELECT b.id bus_id, b.name bus_name, d.from_place || '-' || d.to_place AS destination_name,
			   b.start_time, b.end_time, bs.id reservation_id, bs.created_at booked_date, bs.seat_number
			  FROM bus_seat bs
			  JOIN bus b ON b.id = bs.bus_id
			  JOIN destination d ON b.destination_id = d.id
			  WHERE bs.user_id = $1`
	result, err := u.db.Query(query, userID)
	if err != nil {
		return models.UserBuses{Buses: res}, err
	}

	for result.Next() {
		var b models.UserBus
		err = result.Scan(&b.BusID, &b.BusName, &b.DestinationName,
			&b.StartTime, &b.EndTime, &b.ReservationID, &bookedDate, &b.SeatNumber)
		if err != nil {
			return models.UserBuses{Buses: res}, err
		}
		b.BookedDate = bookedDate.Format("02-01-2006 15:04:05")
		res = append(res, b)
	}
	return models.UserBuses{Buses: res}, nil
}
