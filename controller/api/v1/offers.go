package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/model"
)

// AddOffer godoc
// @Tags offers
// @Router /offers [post]
func (c *Controller) AddOffer(ctx *gin.Context) {
	fmt.Println("CreateOffer")
	ctx.JSON(200, model.Success{Success: true})
}

// UpdateOffer godoc
// @Tags offers
// @Router /offers/{id} [patch]
func (c *Controller) UpdateOffer(ctx *gin.Context) {
	fmt.Println("CreateOffer")
	ctx.JSON(200, model.Success{Success: true})
}

// DeleteOffer godoc
// @Tags offers
// @Router /offers/{id} [delete]
func (c *Controller) DeleteOffer(ctx *gin.Context) {
	fmt.Println("CreateOffer")
	ctx.JSON(200, model.Success{Success: true})
}

// ShowOffer godoc
// @Tags offers
// @Router /offers/{id} [get]
func (c *Controller) ShowOffer(ctx *gin.Context) {
	fmt.Println("CreateOffer")
	ctx.JSON(200, model.Success{Success: true})
}

// ListOffers godoc
// @Tags offers
// @Router /offers [get]
func (c *Controller) ListOffers(ctx *gin.Context) {
	fmt.Println("CreateOffer")
	ctx.JSON(200, model.Success{Success: true})
}