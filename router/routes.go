package router

import (
	v1 "gin-vue-zhtai-server/api/v1"
	"gin-vue-zhtai-server/middleware"
	"gin-vue-zhtai-server/utils"
	"github.com/gin-gonic/gin"
)

func InitRoute() {
	r := gin.New()
	r.Use(middleware.Cors()).Use(gin.Recovery())

	routerV1 := r.Group("api/v1")
	routerV1.POST("login", v1.LoginController)

	r.Run(utils.AppPort)
}
