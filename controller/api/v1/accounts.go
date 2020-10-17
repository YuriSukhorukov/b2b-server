package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (c *Controller) AddAccount(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}