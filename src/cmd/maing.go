package main

import (
	"github.com/Moji00f/SimpleProject/api"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}
