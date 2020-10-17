package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/b2b-server/httputil"

	"github.com/b2b-server/model"
)

// PingExample godoc
// @Summary Возвращает результат проверки доступности e-mail для регистрации
// @Description Возвращает результат проверки доступности e-mail для регистрации
// @Tags auth
// @Accept json
// @Produce json
// @Param email path string true "Email"
// @Success 200 {object} model.Check "Email свободен для регистрации"
// @Success 404 {object} model.Check "Объект не найден"
// @Failure 500 {object} model.Check "Ошибка сервера"
// @Router /auth/email_free/{email} [get]
func (c *Controller) EmailFree(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, model.Check{Success: true, Message: ""})
	// ctx.JSON(http.StatusOK, model.Check{Success: false, Message: ""})
	ctx.JSON(http.StatusOK, model.Check{Success: false, Message: "email занят"})
	// httputil.NewError(ctx, http.StatusBadRequest, err)
	// httputil.NewError(ctx, http.StatusBadRequest, errors.New("WR@R@#R@#"))

}

// PingExample godoc
// @Summary ping auth
// @Description do ping
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /auth/signup [get]
func (c *Controller) SignUp(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}

// PingExample godoc
// @Summary ping auth
// @Description do ping
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /auth/signin [get]
func (c *Controller) SignIn(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}

// PingExample godoc
// @Summary ping auth
// @Description do ping
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /auth/signout [get]
func (c *Controller) SignOut(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}

// PingExample godoc
// @Summary ping auth
// @Description do ping
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /auth/verify [get]
func (c *Controller) Verify(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}