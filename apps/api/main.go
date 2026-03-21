package main

import (
	"github.com/bborutenko/LogForge/internal/config"
	"github.com/bborutenko/LogForge/internal/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitLogger()
	config.InitSettings()
	config.InitDatabase()
	defer config.DBPool.Close()

	r := gin.Default()
	api := r.Group("/api")
	{
		logsGroup := api.Group("/logs")
		{
			logsGroup.GET("", logs.ListLogs)
		}
	}

	r.Run()
}
