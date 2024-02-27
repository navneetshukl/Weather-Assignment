package main

import (
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", routes.GetForm)
	router.POST("/", routes.GetTemperature)
	router.Run()
}
