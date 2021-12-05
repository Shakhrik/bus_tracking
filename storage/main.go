package storage

import (
	// "github.com/Shakhrik/inha/bus_tracking/api/storage/postgres"
	// "github.com/Shakhrik/inha/bus_tracking/api/storage/repo"
	"github.com/Shakhrik/inha/bus_tracking/storage/postgres"
	"github.com/Shakhrik/inha/bus_tracking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	UserRepo() repo.UserRepoI
	DestinationRepo() repo.DestinationRepoI
	BusStopRepo() repo.BusStopRepoI
	BusRepo() repo.BusRepoI
}

type storagePG struct {
	db              *sqlx.DB
	userRepo        repo.UserRepoI
	destinationRepo repo.DestinationRepoI
	busStopRepo     repo.BusStopRepoI
	busRepo         repo.BusRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return storagePG{
		db:              db,
		userRepo:        postgres.NewUserRepo(db),
		destinationRepo: postgres.NewDestinationRepo(db),
		busStopRepo:     postgres.NewBusStopRepo(db),
		busRepo:         postgres.NewBusRepo(db),
	}
}

func (s storagePG) UserRepo() repo.UserRepoI {
	return s.userRepo
}

func (s storagePG) DestinationRepo() repo.DestinationRepoI {
	return s.destinationRepo
}

func (s storagePG) BusStopRepo() repo.BusStopRepoI {
	return s.busStopRepo
}

func (s storagePG) BusRepo() repo.BusRepoI {
	return s.busRepo
}
