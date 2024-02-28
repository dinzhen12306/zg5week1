package routes

import "github.com/gin-gonic/gin"

func NewRoute(engine *gin.Engine) {
	userRoute := UserRoute{}
	userRoute.NewUserRoute(engine)
}
