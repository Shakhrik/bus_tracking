package postgres

import (
	"fmt"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/Shakhrik/inha/bus_tracking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type busStopRepo struct {
	db *sqlx.DB
}

func NewBusStopRepo(db *sqlx.DB) repo.BusStopRepoI {
	return busStopRepo{db: db}
}

func (b busStopRepo) Create(req models.BusStopCreate) (res models.ResponseWithID, err error) {
	query := `INSERT INTO bus_stop(name, destination_id, distance) VALUES($1,$2,$3) RETURNING id`
	err = b.db.Get(&res.ID, query, req.Name, req.DestinationID, req.Distance)
	return
}

func (b busStopRepo) Delete(id int64) (res models.ResponseWithID, err error) {
	query := `DELETE FROM bus_stop WHERE id = $1 RETURNING id`
	err = b.db.Get(&res.ID, query, id)
	return
}

func (b busStopRepo) GetAll(destinationId, limit, page int32) (models.BusStopes, error) {
	var count int64
	res := []models.BusStop{}

	offset := (page - 1) * limit

	query := `SELECT bs.id, bs.name, bs.distance, d.from_place || '-' || d.to_place AS destination_name
			  FROM bus_stop bs
			  JOIN destination d ON bs.destination_id = d.id
			  WHERE bs.destination_id = $1
			  ORDER BY bs.distance
			  LIMIT $2 OFFSET $3`
	err := b.db.Select(&res, query, destinationId, limit, offset)
	if err != nil {
		fmt.Println("error 1")
		return models.BusStopes{BusStopes: res}, err
	}

	queryCount := `SELECT count(1) FROM bus_stop WHERE destination_id = $1 LIMIT $2 OFFSET $3`
	err = b.db.Get(&count, queryCount, destinationId, limit, offset)
	if err != nil {
		fmt.Println("error 2")
		return models.BusStopes{BusStopes: res}, err

	}
	return models.BusStopes{BusStopes: res, Count: count}, nil
}

// func (b busStopRepo) GetAllByBusID (busId, limit, page int32) () {
// 	qury := `SELECT
// 				bs.name,
// 				d.from_place || '-' || d.to_place AS destination_name,
// 				`
// }
