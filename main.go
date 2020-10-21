package main

import (
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/controller/api/v1"

	_ "github.com/b2b-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/b2b-server/service"
	"github.com/b2b-server/repository"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html	

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	g 			:= gin.Default()
	psgql 		:= service.NewPostgresql()
	
	users 		:= repository.NewUserRepository(psgql)
	offers		:= repository.NewOfferRepository(psgql)

	c 			:= controller.NewController(*users, *offers)

	v1 := g.Group("/api/v1") 
	{
		auth := v1.Group("/auth")
		{
			auth.GET("email_free/:email", c.EmailFree)
			auth.POST("signup", c.SignUp)
			auth.POST("signin", c.SignIn)
			auth.DELETE("signout", c.SignOut)
			auth.POST("verify", c.Verify)
		}
	}

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.Run(":8080")
}
