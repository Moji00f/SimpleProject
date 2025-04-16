package main

import (
	"github.com/Moji00f/SimpleProject/api"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/data/cache"
	"github.com/Moji00f/SimpleProject/data/db"
	"github.com/Moji00f/SimpleProject/data/db/migrations"
	"github.com/Moji00f/SimpleProject/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	_ = cache.GetRedis()

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)

	}

	migrations.Up()
	//migrations.Down()

	api.InitServer(cfg)
}
