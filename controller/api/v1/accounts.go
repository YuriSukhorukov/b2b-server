package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// AddAccount godoc
// @Summary Add an account
// @Description add by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body model.AddAccount true "Add account"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [post]
func (c *Controller) AddAccount(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}