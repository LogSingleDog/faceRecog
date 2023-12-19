package handler

import (
	"github.com/gin-gonic/gin"
)

// @Summary 首页
// @Description 登录前首页
// @Tags handler
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Fundmt
// @Router / [get]
func Index(c *gin.Context) {
	c.JSON(200,gin.H{
		"code":200,
		"message":"welcome!",
	})
}
