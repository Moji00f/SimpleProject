package main

import (
	"github.com/Moji00f/SimpleProject/api"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/data/cache"
	"github.com/Moji00f/SimpleProject/data/db"
	"log"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatalf("Init redis has a problem ..., err: %v", err)
	}
	_ = cache.GetRedis()

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatalf("Init postgres has a problem ..., err: %v", err)

	}

	api.InitServer(cfg)
}
