package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/model"
)

// AddOffer godoc
// @Tags offers
// @Accept json
// @Produce json
// @Param offer body model.Offer true "Offer"
// @Success 201 {array} model.Record "Успешное выполнение операции"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /offers [post]
func (c *Controller) AddOffer(ctx *gin.Context) {
	var addOffer model.AddOffer
	if err := ctx.ShouldBindJSON(&addOffer); err != nil {
		ctx.JSON(400, model.Error{Success: false, Error: err.Error()})
		return
	}
	if err := addOffer.Validation(); err != nil {
		ctx.JSON(400, model.Error{Success: false, Error: err.Error()})
		return
	}

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