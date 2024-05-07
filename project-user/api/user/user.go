package user

import "github.com/gin-gonic/gin"

type HandlerUser struct {
}

func (*HandlerUser) getCaptCha(ctx *gin.Context) {
	ctx.JSON(200, "sucess")
}
