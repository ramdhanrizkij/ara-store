package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ramdhanrizkij/ara-store/internal/config"
	"github.com/ramdhanrizkij/ara-store/internal/database"
	"github.com/ramdhanrizkij/ara-store/internal/interfaces/http/middleware"
	"github.com/ramdhanrizkij/ara-store/internal/logger"
)

func main() {
	cfg := config.Load()

	// init logger
	log := logger.NewLogger(cfg.App.Env)
	defer log.Sync()

	// init DB
	db := database.NewPostgres(cfg)
	_ = db

	// Run Migration
	// database.RunMigrations(cfg)

	// init GIN
	r := gin.New()

	r.Use(middleware.LoggeMiddleware(log))
	r.Use(gin.Recovery())

	// Health check
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start server
	addr := fmt.Sprintf(":%s", cfg.App.Port)
	log.Info("Server running on " + addr)

	if err := r.Run(addr); err != nil {
		log.Fatal(err.Error())
	}
}
