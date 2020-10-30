package controller

import "github.com/gin-gonic/gin"

func (c *Controller) AddProposal(ctx *gin.Context) {
	ctx.String(200, "AddProposal")
}

func (c *Controller) ShowProposal(ctx *gin.Context) {
	offerID := ctx.Param("id")
	ctx.String(200, "ShowProposal: %s", offerID)
}