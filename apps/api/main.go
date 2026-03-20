package main

import (
	"github.com/bborutenko/LogForge/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitLogger()
	config.InitSettings()
	config.InitDatabase()
	defer config.DBPool.Close()

	router := gin.Default()

	router.Run()
}
