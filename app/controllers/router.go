package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Ping)

	//r.GET("/home", GetDataForHomepage)
	r.Static("/static", "./static")
	r.StaticFile("/privacy", "./static/privacy_policy.html")

	return r
}
