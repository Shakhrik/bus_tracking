package repo

import "github.com/Shakhrik/inha/bus_tracking/api/models"

type BusRepoI interface {
	Create(req models.BusCreate) (res models.ResponseWithID, err error)
}
