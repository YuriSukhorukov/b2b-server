package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/b2b-server/model"
)

// EmailFree godoc
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
	ctx.JSON(200, model.Success{Success: true})
	// if err != nil {
	// 	ctx.JSON(200, model.Error{Success: false, Error: err.Error()})
	// 	return
	// }
	c.repository.InsertAccount()
}

// SignUp godoc
// @Summary Добавляет нового пользователя
// @Description Возвращает результат операции добавленя нового пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param email header string true "Email"
// @Param password header string true "Password"
// @Success 201 {object} model.Success "Успешное выполнение операции"
// @Success 400 {object} model.Error "Email занят для регистрации"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /auth/signup [post]
func (c *Controller) SignUp(ctx *gin.Context) {
	h := model.AuthHeader{}

	if err := ctx.ShouldBindHeader(&h); err != nil {
		ctx.JSON(200, model.Error{Success: false, Error: err.Error()})
	}

	fmt.Printf("%#v\n", h)
	ctx.JSON(200, model.Success{Success: true})
	// ctx.JSON(200, gin.H{"email": h.Email, "password": h.Password})
}

// SignIn godoc
// @Summary Добавляет HttpOnly Cookie JWT пользователя
// @Description Возвращает результат операции создания HttpOnly Cookie JWT пользователя при авторизации
// @Tags auth
// @Accept json
// @Produce json
// @Param email header string true "Email"
// @Param password header string true "Password"
// @Success 200 {object} model.Success "Успешное выполнение операции"
// @Success 400 {object} model.Error "Пользователь с таким Email не существует"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /auth/signin [post]
func (c *Controller) SignIn(ctx *gin.Context) {
	h := model.AuthHeader{}

	if err := ctx.ShouldBindHeader(&h); err != nil {
		ctx.JSON(200, model.Error{Success: false, Error: err.Error()})
	}

	fmt.Printf("%#v\n", h)
	ctx.JSON(200, model.Success{Success: true})
}

// SignOut godoc
// @Summary Удаляет HttpOnly Cookie JWT пользователя
// @Description Возвращает результат операции удаления HttpOnly Cookie JWT пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} model.Success "Успешное выполнение операции"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /auth/signout [delete]
func (c *Controller) SignOut(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}

// SignOut godoc
// @Summary Удаляет JWT пользователя
// @Description Возвращает результат операции проверки HttpOnly Cookie JWT пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} model.Success "Успешное выполнение операции"
// @Success 400 {object} model.Error "Неудачная валидация HttpOnly Cookie JWT"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /auth/verify [post]
func (c *Controller) Verify(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}