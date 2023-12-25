package router

import (
	"faceRecog/handler"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	//handler里面应该有user
	r.GET("/api/v1", handler.Index) //首页
	r.POST("/api/v1/login",handler.Login)//登录
	r.GET("api/v1/register",handler.RegisterGet)
	r.POST("api/v1/getcode",handler.Getcode)
	//这里就要写token的middleware了
	r.Use(handler.TokenMiddle)
	r.GET("api/v1/check",handler.Check)
	r.POST("api/v1/register",handler.Register)
	r.GET("api/v1/setpassword",handler.InputPassword)
	r.POST("api/v1/setpassword",handler.SetPassword)
}
