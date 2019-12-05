package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shikanon/IoTOrbHub/pkg/api"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.POST("/api/v1/login", api.DeviceManager)
	router.Run("0.0.0.0:9898")
}
