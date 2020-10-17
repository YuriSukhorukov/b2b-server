package main

import (
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/controller/api/v1"
)

func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1") 
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET("", c.AddAccount)
		}

		auth := v1.Group("/auth")
		{
			auth.GET("email_free", c.EmailFree)
			auth.GET("signup", c.SignUp)
			auth.GET("signin", c.SignIn)
			auth.GET("signout", c.SignOut)
			auth.GET("verify", c.Verify)
		}

		examples := v1.Group("/examples")
		{
			examples.GET("ping", c.PingExample)
		}
	}

	r.Run(":8080")
}
