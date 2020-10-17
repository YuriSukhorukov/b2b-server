package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/b2b-server/model"
)

// PingExample godoc
// @Summary Возвращает результат проверки доступности e-mail для регистрации
// @Description Возвращает результат проверки доступности e-mail для регистрации
// @Tags auth
// @Accept json
// @Produce json
// @Param email path string true "Email"
// @Success 200 {object} model.Success "Email свободен для регистрации"
// @Success 400 {object} model.Error "Email занят для регистрации"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /auth/email_free/{email} [get]
func (c *Controller) EmailFree(ctx *gin.Context) {
	email := ctx.Param("email")
	fmt.Printf("email: %s\n", email)
	ctx.JSON(http.StatusOK, model.Success{Success: true})
	// if err != nil {
	// 	ctx.JSON(http.StatusOK, model.Error{Success: false, Error: "Something wrong"})
	// 	return
	// }
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