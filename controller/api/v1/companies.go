package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/b2b-server/model"
)

// AddCompany godoc
// @Tags companies
// @Router /offers/{offerID}/proposals/{proposalID} [get]
func (c *Controller) AddCompany(ctx *gin.Context) {
	fmt.Println("AddCompany")
	ctx.JSON(200, model.Success{Success: true})
}

// ListCompanies godoc
// @Tags companies
// @Router /companies [get]
func (c *Controller) ListCompanies(ctx *gin.Context) {
	fmt.Println("ListCompanies")
	ctx.JSON(200, model.Success{Success: true})
}

// ShowCompany godoc
// @Tags companies
// @Router /companies/{companyID} [get]
func (c *Controller) ShowCompany(ctx *gin.Context) {
	fmt.Println("ShowCompany")
	ctx.JSON(200, model.Success{Success: true})
}

// UpdateCompany godoc
// @Tags companies
// @Router /companies/{companyID} [patch]
func (c *Controller) UpdateCompany(ctx *gin.Context) {
	fmt.Println("UpdateCompany")
	ctx.JSON(200, model.Success{Success: true})
}