package routes

import (
	"github.com/gin-gonic/gin"
	"week1/client/controllers"
)

type UserRoute struct {
	controllers.UserController
}

func (r *UserRoute) NewUserRoute(engine *gin.Engine) {
	g := engine.Group("user")
	{
		g.POST("/login", r.Login)       //登录
		g.POST("/register", r.Register) //注册
		g.POST("/upload", r.UpLoad)     //文件上传
	}
}
