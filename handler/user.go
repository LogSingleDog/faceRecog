package handler

import (
	"fmt"
	"faceRecog/service"
	"faceRecog/model"
	"math/rand"
	"time"
	"github.com/gin-gonic/gin"
)

// @Summary 注册页面
// @Description GET注册页面
// @Tags user
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Fundmt
// @Router /register [get]
func RegisterGet(c *gin.Context)  {
	c.JSON(200,gin.H{
		"code":200,
		"message":"register index",
	})
}

// @Summary 获得验证码
// @Description 获得邮箱验证码
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Email formData string true "邮箱"
// @Failure 401 {object} model.Fundmt
// @Failure 402 {object} model.Fundmt
// @Success 200 {object} model.Token
// @Router /getcode [post]
func Getcode(c *gin.Context){
	email:=c.PostForm("Email");
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06v", 
		rand.New(rand. NewSource(time. Now().UnixNano())).
		Int31n(1000000))
	err2:=service.SendEmail(email,code)
	if err2!=nil {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "邮件发送失败",
		})
		return
	}
	model.UpdateCode(email,code)
	tokenStr,err:=service.GenerateToken(email)//生成token和一个错误
	if err != nil {
		c.JSON(402,gin.H{
			"code":402,
			"message":"请重试",	
		})
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"message":"send email ok",
			"token":tokenStr,
		})
	}
}

// @Summary 检查验证码
// @Description GET页面
// @Tags user
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Fundmt
// @Router /check [get]
func Check(c *gin.Context)  {
	c.JSON(200,gin.H{
		"code":200,
		"message":"check code page",
	})
}

// @Summary 检查验证码并注册
// @Description 比较验证码
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Code formData string true "验证码"
// @Param Authorization header string true "token"
// @Success 200 {object} model.Fundmt
// @Failure 401 {object} model.Fundmt
// @Failure 403 {object} model.Fundmt
// @Router /register [post]
func Register(c *gin.Context)  {
	//我想干嘛？
	//我现在要验证usertmp所读取的code是否和数据库里面存的一样
	//所以要拿现在的code去和数据库里面比对
	//对的上才进行下一步操作
	//下面的一会儿再改
	//判断有没有搞验证码
	var usertmp model.UserTmp
	code :=c.PostForm("Code")
	email:=service.GetEmail(c)
	resu := model.DB.Where("email = ?", email).First(&usertmp); 
	if resu.Error != nil {//还没搞验证码
		c.JSON(401,gin.H{
			"code":401,
			"message": "you don't get code",
		})
		return;
	}
	if code != usertmp.Code{
		c.JSON(403,gin.H{
			"code":403,
			"message":"code is wrong",
		})
		return;
	}
	c.JSON(200,gin.H{
		"code":200,
		"message":"code is right",
	})
}

// @Summary 输入密码
// @Description GET页面
// @Tags user
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Fundmt
// @Router /inputpassword [get]
func InputPassword(c *gin.Context)  {
	c.JSON(200,gin.H{
		"code":200,
		"message":"please input your password",
	})
}

// @Summary 设置密码
// @Description 输入两个密码并检查
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param Password1 formData string true "输入密码"
// @Param Password2 formData string true "再次输入密码"
// @Success 200 {object} model.Fundmt
// @Failure 400 {object} model.Fundmt
// @Router /setpassword [post]
func SetPassword(c *gin.Context)  {
	pwd1:=c.PostForm("Password1")
	pwd2:=c.PostForm("Password2")
	if pwd1!=pwd2 {
		c.JSON(400,gin.H{
			"code":400,
			"message":"密码不相同",
		})
		return;
	}
	email:=service.GetEmail(c)
	model.UpdateUser(email,pwd1)
	c.JSON(200,gin.H{
		"code":200,
		"message":"set password ok",
	})
}

// @Summary 登录
// @Description 账密登录
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Email formData string true "邮箱"
// @Param Password formData string true "密码"
// @Success 200 {object} model.Token
// @Failure 402 {object} model.Fundmt
// @Failure 401 {object} model.Fundmt
// @Failure 400 {object} model.Fundmt
// @Router /login [post]
func Login(c *gin.Context) {
	var userr model.User
	err1 := c.Bind(&userr)
	if err1 != nil {
		c.JSON(402, gin.H{
			"code":    402,
			"message": "输入信息格式有误",
		})
		return
	}
	if userr.Email == "" {
		c.JSON(402, gin.H{
			"code":    402,
			"message": "邮箱不为空",
		})
		return
	}
	pwd:=userr.Password
	if resu := model.DB.Where("email = ?", userr.Email).First(&userr); resu.Error != nil {
		//这个分支里面是没有注册的情况
		c.JSON(401, gin.H{
			"code":401,
			"message":"you haven't register",
		})
		return
	}else{//用户注册过
		//pwd是用户输进来的密码
		if  pwd!=userr.Password {
			c.JSON(401, gin.H{
				"code":401,
				"message":"Password or account wrong.",
			})
			return//注册过，但是密码错了
		}
	}
	//后续生成token
	tokenStr,err:=service.GenerateToken(userr.Email)//生成token和一个错误
	if err!=nil{
		c.JSON(400,gin.H{
			"code":400,
			"message":"登陆失败，请重试",
		})
		fmt.Printf("token err:%v\n",err)
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"message":"登陆成功",
			"token":tokenStr,
		})
	}
}