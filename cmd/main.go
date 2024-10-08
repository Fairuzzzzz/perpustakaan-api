package main

import (
	"log"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/configs"
	"github.com/Fairuzzzzz/perpustakaan-api/internal/handler/membership"
	membershipsRepo "github.com/Fairuzzzzz/perpustakaan-api/internal/repository/memberships"
	membershipsSvc "github.com/Fairuzzzzz/perpustakaan-api/internal/service/memberships"
	"github.com/Fairuzzzzz/perpustakaan-api/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisialisasi config", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisialisasi database", err)
	}

	membershipsRepo := membershipsRepo.NewRepository(db)

	membershipService := membershipsSvc.NewService(cfg, membershipsRepo)

	membershipHandler := membership.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
