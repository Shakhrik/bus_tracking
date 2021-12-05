package repo

import "github.com/Shakhrik/inha/bus_tracking/api/models"

type BusRepoI interface {
	Create(req models.BusCreate) (res models.ResponseWithID, err error)
	GetAll(destinationId, limit, page int32) (models.Buses, error)
	ReserveBus(busId int64) (res models.ResponseWithID, err error)
	Delete(id int64) (res models.ResponseWithID, err error)
}
