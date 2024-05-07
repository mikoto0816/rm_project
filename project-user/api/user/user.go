package user

import (
	"github.com/gin-gonic/gin"
	common "rimomi.com/project-common"
)

type HandlerUser struct {
}

func (*HandlerUser) getCaptCha(ctx *gin.Context) {
	result := &common.Result{}
	ctx.JSON(200, result.Success("成功"))
}
