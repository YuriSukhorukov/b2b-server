package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (c *Controller) EmailFree(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")	
}

func (c *Controller) SignUp(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}

func (c *Controller) SignIn(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}

func (c *Controller) SignOut(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}

func (c *Controller) Verify(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}