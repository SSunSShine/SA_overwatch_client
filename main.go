package main

import (
	"github.com/SSunSShine/SA_overwatch_client/api"
	"github.com/SSunSShine/SA_overwatch_client/middleware"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.POST("/CMD", api.CMD)
	r.GET("/procInfo", api.ProcsInfo)
	r.GET("/memInfo", api.MemInfo)
	r.GET("/diskInfo", api.Disk)

	r.Run(":8000")
}
