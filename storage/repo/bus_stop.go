package repo

import "github.com/Shakhrik/inha/bus_tracking/api/models"

type BusStopRepoI interface {
	Create(req models.BusStopCreate) (res models.ResponseWithID, err error)
	Delete(id int64) (res models.ResponseWithID, err error)
	GetAll(destinationId, limit, page int32) (models.BusStopes, error)
}
