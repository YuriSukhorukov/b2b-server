package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/model"
	"github.com/b2b-server/controller/api/v1"
)

func BasicAuthApi(ctx *gin.Context) {
	user, password, hasAuth := ctx.Request.BasicAuth()
	if hasAuth != true {
		ctx.JSON(401, model.Error{Success: false, Error: controller.ErrNotAuthorized.Error()})
		ctx.Abort()
		return
	}
	if user != "admin" || password != "admin" {
		ctx.JSON(401, model.Error{Success: false, Error: controller.ErrNotAuthorized.Error()})
		ctx.Abort()
		return
	}
}