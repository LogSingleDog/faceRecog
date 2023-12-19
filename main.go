package main
import (
	"faceRecog/config"
	"faceRecog/model"
	"faceRecog/router"
	"github.com/gin-gonic/gin"
)

// @title FaceRecognize API
// @version 1.0
// @description 人脸识别API
// @host localhost
// @BasePath /api/v1
func main() {
	err:=config.Init("")//这个是init config.yaml文件
	if err!=nil{
		panic(err)
	}
	//println("200 ok")
	//config.LoadQiniu()
	model.Init()
	defer model.DB.Close()
	model.InitTables()

	r:=gin.Default()
	router.Router(r)
	r.Run(":1145")
}