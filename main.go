package main

import (
	"CodeHive/database"
	"CodeHive/logger"
	"CodeHive/routes"
	"CodeHive/tools"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		logger.Log.Error(err)
		return
	}
	configs, err := tools.ConfigSet()
	if err != nil {
		logger.Log.Errorf("%s", err)
		return
	}

	logger.Log.Info("Start application " + configs.App.Name)
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "hx-target", "hx-trigger", "hx-current-url", "hx-request", "X-Requested-With"},
		AllowCredentials: true,
	}

	r.Use(cors.New(corsConfig))
	routes.RegisterCodeRoutes(r, db)
	routes.RegisterLanguageRoutes(r, db)

	_ = r.Run(":" + configs.Server.Port)
}
