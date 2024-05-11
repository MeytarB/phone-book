package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (pbc *PhoneBookController) ShowAllContacts(ctx *gin.Context) {
	page := ctx.Params.ByName("page")
	pageNumber, _ := strconv.Atoi(page)
	contacts, err := pbc.service.ShowAllContacts(int64(pageNumber))

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK , contacts)
}

func (pbc *PhoneBookController) FindContactByName(ctx *gin.Context) {
	firstName := ctx.Query("firstname")
	lastName := ctx.Query("lastname")
	contact, err := pbc.service.FindContactByName(firstName, lastName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, contact)
}

func (pbc *PhoneBookController) FindContactByNumber(ctx *gin.Context) {
	phoneNumber := ctx.Query("phonenumber")
	contact, err := pbc.service.FindContactByNumber(phoneNumber)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, contact)
}