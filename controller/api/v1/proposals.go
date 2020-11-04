package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/model"
)

// AddProposal godoc
// @Tags proposals
// @Accept json
// @Produce json
// @Param id path string true "Offer ID"
// @Param offer body model.Proposal true "Proposal"
// @Success 201 {array} model.Created "Успешное выполнение операции"
// @Failure 500 {object} model.Error "Ошибка сервера"
// @Router /offers/{id}/proposals [post]
func (c *Controller) AddProposal(ctx *gin.Context) {
	var addProposal model.AddProposal

	offerID := ctx.Param("id")
	fmt.Println(offerID)

	if err := ctx.ShouldBindJSON(&addProposal); err != nil {
		ctx.JSON(400, model.Error{Success: false, Error: err.Error()})
		return
	}
	if err := addProposal.Validation(); err != nil {
		ctx.JSON(400, model.Error{Success: false, Error: err.Error()})
		return
	}

	cookie, err := ctx.Cookie("JWT")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(401, model.Error{Success: false, Error: ErrNotAuthorized.Error()})
		return
	}

	err, userID := c.JWT.Decode(cookie)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(401, model.Error{Success: false, Error: ErrNotAuthorized.Error()})
		return
	}

	proposal := model.Proposal{
		UserID: userID,
		OfferID: offerID,
	}

	err, result := c.ProposalRepository.InsertProposal(proposal)

	if err != nil {
		ctx.JSON(400, model.Error{Success: false, Error: ErrSomethingWrong.Error()})
		return
	}

	err = result.Validation()
	if err != nil {
		ctx.JSON(400, model.Error{Success: false, Error: ErrSomethingWrong.Error()})
		return
	}

	ctx.JSON(200, result)
}

func (c *Controller) ShowProposal(ctx *gin.Context) {
	offerID := ctx.Param("id")
	ctx.String(200, "ShowProposal: %s", offerID)
}