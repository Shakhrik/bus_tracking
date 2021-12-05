package repo

import "github.com/Shakhrik/inha/bus_tracking/api/models"

type DestinationRepoI interface {
	Create(req models.DestinationCreate) (res models.DestinationGet, err error)
	Update(req models.DestinationUpdate) (*models.DestinationGet, error)
	Get(id int64) (res models.Destination, err error)
	GetAll(limit, page int32) (res models.Destinations, err error)
	Delete(id int32) (res models.DestinationGet, err error)
}
