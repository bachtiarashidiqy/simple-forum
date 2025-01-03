package main

import (
	"github.com/bachtiarashidiqy/simple-forum/internal/configs"
	membershipRepo "github.com/bachtiarashidiqy/simple-forum/internal/repository/memberships"
	"github.com/bachtiarashidiqy/simple-forum/pkg/internalsql"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)
	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"))
	if err != nil {
		log.Fatal("error loading config", err)
	}
	cfg = configs.Get()
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	_ := membershipRepo.NewRepository(db)
	err = r.Run(cfg.Service.Port)
	if err != nil {
		return
	}
}
