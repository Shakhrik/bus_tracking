package postgres

import (
	"database/sql"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/Shakhrik/inha/bus_tracking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type destinationRepo struct {
	db *sqlx.DB
}

func NewDestinationRepo(db *sqlx.DB) repo.DestinationRepoI {
	return destinationRepo{db: db}
}

func (d destinationRepo) Create(req models.DestinationCreate) (id int64, err error) {
	query := `INSERT INTO destination(from_place, to_place, price, distance) VALUES($1,$2,$3,$4) RETURNING id`
	err = d.db.Get(&id, query, req.From, req.To, req.Price, req.Distance)
	return
}

func (d destinationRepo) Update(req models.DestinationUpdate) (int64, error) {
	query := `UPDATE destination SET 
					from_place = $1,
					to_place = $2,
					price = $3,
					distance = $4,
					updated_at = now()
			 WHERE id = $5`
	res, err := d.db.Exec(query, req.From, req.To, req.Price, req.Distance, req.ID)
	if err != nil {
		return 0, err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return 0, sql.ErrNoRows
	}

	return req.ID, nil
}
