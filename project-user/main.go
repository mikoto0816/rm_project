package main

import (
	"github.com/gin-gonic/gin"
	srv "rimomi.com/project-common"
	_ "rimomi.com/project-user/api"
	_ "rimomi.com/project-user/api/user"
	"rimomi.com/project-user/router"
)

func main() {

	engine := gin.Default()

	//初始化路由
	router.InitRouter(engine)
	srv.Run(engine, "project-user", ":80")
}
