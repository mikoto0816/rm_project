package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"rimomi.com/project-user/router"
)

func init() {
	log.Println("初始化路由")
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	r.GET("/project/login/getCaptcha", h.getCaptCha)
}
