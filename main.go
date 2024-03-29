package main

import (
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/middleware"
	"github.com/b2b-server/controller/api/v1"

	_ "github.com/b2b-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/b2b-server/service"
	"github.com/b2b-server/repository"
	"os"
	"fmt"
	"strconv"
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
	signKey 			:= os.Getenv("PRIVATE_KEY")
	verifyKey 			:= os.Getenv("PUBLIC_KEY")

	db_host_name 		:= os.Getenv("DB_HOST")
	db_user_name 		:= os.Getenv("DB_USERNAME")
	db_password 		:= os.Getenv("DB_PASSWORD")
	db_database_name 	:= os.Getenv("DB_DATABASE")
	db_ssl_mode 		:= os.Getenv("DB_SSL_MODE")
	db_host_port, err 	:= strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
	
	jwt := service.NewJWT(signKey, verifyKey)

	psgql 				:= service.NewPostgresql(
		db_host_name,
		db_host_port,
		db_user_name,
		db_password,
		db_database_name,
		db_ssl_mode,
	)
	
	users 		:= repository.NewUserRepository(psgql)
	offers		:= repository.NewOfferRepository(psgql)
	proposals	:= repository.NewProposalRepository(psgql)

	g 			:= gin.Default()
	c 			:= controller.NewController(*jwt, *users, *offers, *proposals)

	v1 := g.Group("/api/v1")
	{
		auth := v1.Group("")
		{
			auth.GET("email_free/:email", c.EmailFree)
			auth.POST("signup", c.AddUser)
			auth.POST("signin", c.Authenticate)
			auth.POST("auth", c.Authorize)
			auth.DELETE("signout", c.SignOut)
		}
		offers := v1.Group("/offers")
		{
			offers.GET(":offerID", c.ShowOffer)
			offers.GET("", c.ListOffers)
			offers.POST("", c.AddOffer)
			offers.PATCH(":offerID", c.UpdateOffer)
			offers.DELETE(":offerID", c.DeleteOffer)


			proposals := offers.Group(":offerID/proposals")
			{
				proposals.GET(":proposalID", c.ShowProposal)
				proposals.GET("", c.ListProposals)
				proposals.POST("", c.AddProposal)
			}
		}
	}
	swagger := g.Group("/swagger")
	swagger.Use(middleware.BasicAuthSwag)
	{
		swagger.GET("*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	g.Run(":8080")
}