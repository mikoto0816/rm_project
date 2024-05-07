package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	common "rimomi.com/project-common"
	"rimomi.com/project-user/pkg/model"
	"time"
)

type HandlerUser struct {
}

func (*HandlerUser) getCaptCha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	if !common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, result.Fail(model.NoLegalMobile, "手机号不合法"))
		return
	}
	code := "123456"
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("发送短信成功")
	}()
	ctx.JSON(http.StatusOK, result.Success(code))
}
