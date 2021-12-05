package main

import (
	"fmt"

	"github.com/Shakhrik/bus_tracking/api/api"

	"github.com/Shakhrik/bus_tracking/api/config"
	"github.com/Shakhrik/bus_tracking/api/pkg/logger"
	"github.com/Shakhrik/bus_tracking/api/storage"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	cfg := config.Load()
	logger := logger.New(cfg.LogLevel, "safia_catalogue_service")

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", cfg.PostgresHost,
		cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase, "disable")
	db, err := sqlx.Connect("pgx", conStr)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)
	storageI := storage.NewStoragePG(db)

	apiServer := api.New(&api.RouterOptions{
		Log:     logger,
		Cfg:     &cfg,
		Storage: storageI,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		panic(err)
	}

}
