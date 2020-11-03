package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/model"
)

func (c *Controller) AddProposal(ctx *gin.Context) {
	var addProposal model.AddProposal

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
		OfferID: addProposal.OfferID,
	}

	err, result := c.ProposalRepository.InsertProposal(proposal)

	if err := result.Validation(); err != nil {
		ctx.JSON(400, model.Error{Success: false, Error: ErrSomethingWrong.Error()})
		return
	}

	ctx.JSON(200, result)
}

func (c *Controller) ShowProposal(ctx *gin.Context) {
	offerID := ctx.Param("id")
	ctx.String(200, "ShowProposal: %s", offerID)
}