package controller

import (
	"fmt"
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
// @Failure 400 {object} model.Error "Email занят для регистрации"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /email_free/{email} [get]
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
// @Success 201 {object} model.Record "Успешное выполнение операции"
// @Failure 400 {object} model.Error "Email занят для регистрации"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /signup [post]
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
// @Failure 400 {object} model.Error "Неверный Email или Password"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /signin [post]
func (c *Controller) SignIn(ctx *gin.Context) {
	h := model.AuthHeader{}

	if err := ctx.ShouldBindHeader(&h); err != nil {
		ctx.SetCookie("JWT", "", 0, "/", "localhost", true, true)
		ctx.JSON(500, model.Error{Success: false, Error: err.Error()})
		return
	}

	e := h.Email
	p := h.Password
	err, result := c.UserRepository.AuthorizeUser(e, p)

	if err != nil {
		fmt.Printf(err.Error())
		ctx.SetCookie("JWT", "", 0, "/", "localhost", true, true)
		ctx.JSON(500, model.Error{Success: false, Error: "something went wrong"})
		return
	} else if result == nil {
		ctx.SetCookie("JWT", "", 0, "/", "localhost", true, true)
		ctx.JSON(400, model.Error{Success: false, Error: "email or password is not correct"})
		return
	}

	err, token := c.JWT.Encode(result.UserID)

	if err != nil {
		fmt.Printf(err.Error())
		ctx.SetCookie("JWT", "", 0, "/", "localhost", true, true)
		ctx.JSON(500, model.Error{Success: false, Error: "something went wrong"})
		return
	}

	second := 1
	minute := second * 60
	hour := minute * 60
	day := hour * 24
	week := day * 8

	maxAge 		:= week
	path 		:= "/"
	domain 		:= "localhost"
	secure 		:= false
	httpOnly 	:= true

	ctx.SetCookie("JWT", token, maxAge, path, domain, secure, httpOnly)
	ctx.JSON(200, model.Success{Success: true})
}

// Auth godoc
// @Summary Авторизация пользователя проверкой HttpOnly Cookie JWT
// @Description Возвращает результат операции авторизации HttpOnly Cookie JWT пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} model.Success "Успешное выполнение операции"
// @Failure 400 {object} model.Error "Неудачная авторизация HttpOnly Cookie JWT"
// @Failure 401 {object} model.Error "Токен JWT отсутствует"
// @Router /auth [post]
func (c *Controller) Auth(ctx *gin.Context) {
	cookie, err := ctx.Cookie("JWT")

	if err != nil {
		fmt.Println(err)
		ctx.JSON(401, model.Error{Success: false, Error: "no token"})
		return
	}

	err, _ = c.JWT.Decode(cookie)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, model.Error{Success: false, Error: "validation failure"})
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
// @Router /signout [delete]
func (c *Controller) SignOut(ctx *gin.Context) {
	ctx.SetCookie("JWT", "", 0, "/", "localhost", true, true)
	ctx.JSON(200, model.Success{Success: true})
}