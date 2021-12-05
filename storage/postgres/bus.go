package postgres

import (
	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/Shakhrik/inha/bus_tracking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type busRepo struct {
	db *sqlx.DB
}

func NewBusRepo(db *sqlx.DB) repo.BusRepoI {
	return busRepo{db: db}
}

func (b busRepo) Create(req models.BusCreate) (res models.ResponseWithID, err error) {
	var busStop models.BusStop
	queryBusStop := `SELECT id, min(distance) as distance
					 FROM bus_stop
					 WHERE destination_id = $1
					 GROUP BY destination_id, id`

	query := `INSERT INTO bus(name, seat_count, start_time, end_time, destination_id, bus_stop_id)
			  VALUES($1, $2, $3, $4, $5, $6) RETURNING id`

	err = b.db.Get(&busStop, queryBusStop, req.DestinationID)
	if err != nil {
		return models.ResponseWithID{}, err
	}

	err = b.db.Get(&res.ID, query, req.Name, req.SeatCount, req.StartTime, req.EndTime, req.DestinationID, busStop.ID)
	if err != nil {
		return models.ResponseWithID{}, err
	}

	return
}

func (b busRepo) GetAll(destinationId, limit, page int32) (models.Buses, error) {
	var count int64
	res := []models.Bus{}

	offset := (page - 1) * limit

	query := `WITH temp AS(
		select b.id as id, count(1) as booked_seats_count
		from bus b
		left join bus_seat bs on b.id = bs.bus_id
		group by b.id
		)
		select b.id, b.start_time, b.end_time, b.name, b.seat_count, 
		d.from_place || '-' || d.to_place AS destination_name, 
		is_full, b.seat_count - t.booked_seats_count as remaining_seats
		from temp t
		join bus b on t.id = b.id
		JOIN destination d ON b.destination_id = d.id
		WHERE b.destination_id = $1
		LIMIT $2 OFFSET $3`
	err := b.db.Select(&res, query, destinationId, limit, offset)
	if err != nil {
		return models.Buses{Buses: res}, err
	}

	queryCount := `SELECT count(1) FROM bus_stop WHERE destination_id = $1 LIMIT $2 OFFSET $3`
	err = b.db.Get(&count, queryCount, destinationId, limit, offset)
	if err != nil {
		return models.Buses{Buses: res}, err

	}
	return models.Buses{Buses: res, Count: count}, nil
}

func (b busRepo) ReserveBus(busId int64) (res models.ResponseWithID, err error) {
	query := `INSERT INTO bus_seat (bus_id) VALUES($1) RETURNING id`
	err = b.db.Get(&res.ID, query, busId)
	return
}

func (b busRepo) Delete(id int64) (res models.ResponseWithID, err error) {
	query := `DELETE FROM bus WHERE id = $1 RETURNING id`
	err = b.db.Get(&res.ID, query, id)
	return
}
