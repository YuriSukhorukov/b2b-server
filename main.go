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
	}

	r.Run(":8080")
}
