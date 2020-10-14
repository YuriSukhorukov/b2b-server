package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "I'm b2b server!")
		})
	}

	r.Run(":8080")

	fmt.Println("b2b server started")
}
