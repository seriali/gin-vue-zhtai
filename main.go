package main

import (
	"gin-vue-zhtai-server/model"
	"gin-vue-zhtai-server/router"
)

func main() {
	//数据库初始化
	model.InitDb()
	//初始化api
	router.InitRoute()
}
