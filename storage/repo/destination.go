package repo

import "github.com/Shakhrik/inha/bus_tracking/api/models"

type DestinationRepoI interface {
	Create(req models.DestinationCreate) (int64, error)
	Update(req models.DestinationUpdate) (int64, error)
	Get(id int64) (res models.Destination, err error)
	GetAll(limit, page int32) (res models.Destinations, err error)
}
