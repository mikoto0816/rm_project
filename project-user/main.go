package main

import (
	"github.com/gin-gonic/gin"
	srv "rimomi.com/project-common"
)

func main() {

	engine := gin.Default()
	srv.Run(engine, "project-user", ":80")
}
