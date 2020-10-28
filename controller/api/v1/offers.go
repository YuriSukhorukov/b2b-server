package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/model"
	"github.com/b2b-server/constant"
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
	var offer model.Offer
	if err := ctx.ShouldBindJSON(&offer); err != nil {
		ctx.JSON(400, model.Error{Success: false, Error: err.Error()})
		return
	}
	fmt.Println(offer)
	fmt.Println(constant.MEASURE_UNIT_CODE_TON)


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