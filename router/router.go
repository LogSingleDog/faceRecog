package router
import (
	"faceRecog/handler"
	"github.com/gin-gonic/gin"
	
)

func Router(r *gin.Engine) {
	//handler里面应该有user
	r.GET("/api/v1", handler.Index) //首页
	r.POST("api/v1/register",handler.Register)
	r.POST("api/v1/getcode",handler.Getcode)
	r.POST("/api/v1/login",handler.Login)//登录
}
