package controller

import (
	"net/http"

	"github.com/MeytarB/phone-book/types"
	"github.com/gin-gonic/gin"
)

func (pbc *PhoneBookController) EditContact(ctx *gin.Context) {
	var contact types.ContactType
	firstName := ctx.Params.ByName("firstname")
	lastName := ctx.Params.ByName("lastname")

	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := pbc.service.EditContact(firstName, lastName, &contact)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK , gin.H{"message": "OK"})
}
