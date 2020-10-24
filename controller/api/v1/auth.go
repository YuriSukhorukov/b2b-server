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
	email 				:= ctx.Param("email")
	err, result 		:= c.UserRepository.IsEmailFree(email)
	
	if err != nil {
		fmt.Printf(err.Error())
		ctx.JSON(500, model.Error{Success: false, Error: "something went wrong"})
		return
	} else if result != true {
		ctx.JSON(400, model.Error{Success: false, Error: "email is not available"})
		return
	}

	ctx.JSON(200, model.Success{Success: true})
}

// SignUp godoc
// @Summary Добавляет нового пользователя
// @Description Возвращает результат операции добавленя нового пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param email header string true "Email"
// @Param password header string true "Password"
// @Success 201 {array} model.Record "Успешное выполнение операции"
// @Success 400 {object} model.Error "Email занят для регистрации"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /auth/signup [post]
func (c *Controller) SignUp(ctx *gin.Context) {
	h 		:= model.AuthHeader{}
	err 	:= ctx.ShouldBindHeader(&h); 
	if err != nil {
		ctx.JSON(200, model.Error{Success: false, Error: err.Error()})
	}

	e := h.Email
	p := h.Password
	err, result := c.UserRepository.InsertUser(e, p)

	if err != nil {
		switch err.Error() {
			case "duplicate email":
				ctx.JSON(400, model.Error{Success: false, Error: err.Error()})
				return
			case "something wrong":
				ctx.JSON(500, model.Error{Success: false, Error: err.Error()})
				fmt.Printf(err.Error())
				return
		}
	}

	ctx.JSON(200, result)
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

	e := h.Email
	p := h.Password
	err, result := c.UserRepository.AuthorizeUser(e, p)

	if err != nil {
		fmt.Printf(err.Error())
		ctx.JSON(500, model.Error{Success: false, Error: "something went wrong"})
		return
	} else if result != true {
		ctx.JSON(400, model.Error{Success: false, Error: "email or password is not correct"})
		return
	}

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

// Verify godoc
// @Summary Валидация JWT пользователя
// @Description Возвращает результат операции валидации HttpOnly Cookie JWT пользователя
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