package main

import (
	"CodeHive/database"
	"CodeHive/logger"
	"CodeHive/routes"
	"CodeHive/tools"
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
	routes.RegisterCodeRoutes(r, db)
	routes.RegisterLanguageRoutes(r, db)

	_ = r.Run(":" + configs.Server.Port)
}
