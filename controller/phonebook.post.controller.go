package controller

import (
	"net/http"

	"github.com/MeytarB/phone-book/service"
	"github.com/MeytarB/phone-book/types"
	"github.com/gin-gonic/gin"
)

func (pbc *PhoneBookController) AddContact(ctx *gin.Context) {
	var contact types.Contact
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON((http.StatusBadRequest), gin.H{"message": err.Error()})
	}
	err := pbc.service.AddContact(&contact)
	
	if err != nil {
		if service.IsUserError(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON((http.StatusBadGateway), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}