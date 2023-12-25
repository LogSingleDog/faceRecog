package handler

import (
	"faceRecog/service"
	"github.com/gin-gonic/gin"
)

// @Summary  身份验证
// @Description  获取token并验证，成功则将Set Email
// @Tags middleware
// @Accept application/json
// @Produce application/json
// @Param token header string false "token"
// @Failure 401 {object} model.Fundmt
func TokenMiddle(c *gin.Context) {
	email := service.GetToken(c)//得到的就是id，不是结构体
	if email == "0" {
		c.JSON(401, gin.H{
			"code": 401, 
			"messsage": "权限不足",
		})
		c.Abort()
		return
		//到此一游
	} else {
		c.Set("Email", email)
		c.Next()
		//根据 id 找到对应用户信息并返回
		//后期取id的时候使用id,err:=c.Get(UserId)取出id即可
	}
}