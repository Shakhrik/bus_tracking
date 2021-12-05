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
