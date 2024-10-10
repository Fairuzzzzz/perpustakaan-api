package main

import (
	"log"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/configs"
	"github.com/Fairuzzzzz/perpustakaan-api/internal/handler/books"
	"github.com/Fairuzzzzz/perpustakaan-api/internal/handler/membership"
	bookRepo "github.com/Fairuzzzzz/perpustakaan-api/internal/repository/books"
	membershipsRepo "github.com/Fairuzzzzz/perpustakaan-api/internal/repository/memberships"
	bookSvc "github.com/Fairuzzzzz/perpustakaan-api/internal/service/books"
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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipsRepo := membershipsRepo.NewRepository(db)
	bookRepo := bookRepo.NewRepository(db)

	membershipService := membershipsSvc.NewService(cfg, membershipsRepo)
	bookService := bookSvc.NewService(cfg, bookRepo)

	membershipHandler := membership.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	bookHandler := books.NewHandler(r, bookService)
	bookHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
