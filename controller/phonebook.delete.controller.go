package controller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func (pbc *PhoneBookController) DeleteContact(ctx *gin.Context) {
	firstName := ctx.Query("firstname")
	lastName := ctx.Query("lastname")
	err := pbc.service.DeleteContact(firstName, lastName)
	
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK , gin.H{"message": "OK"})
}


func (pbc *PhoneBookController) DeleteAll(ctx *gin.Context) {

	err := pbc.service.DeleteAll()
	
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK , gin.H{"message": "OK"})
}