package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func New(service PhoneBookService) PhoneBookController {
	return PhoneBookController{
		Service: service,
	}
}

func (pbc *PhoneBookController) AddContact(ctx *gin.Context) {
	var contact ContactType
	if err:= ctx.ShouldBindJSON(&contact); err != nil{
		ctx.JSON((http.StatusBadRequest), gin.H{"message": err.Error()})
	}
	err:= pbc.Service.AddContact(&contact)
	if err != nil{
		ctx.JSON((http.StatusBadGateway), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK ,"ok")
}

func (pbc *PhoneBookController) EditContact(ctx *gin.Context) {

	ctx.JSON(http.StatusOK ,"ok")
}

func (pbc *PhoneBookController) ShowAllContacts(ctx *gin.Context) {
	contacts, err := pbc.Service.ShowAllContacts()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK , contacts)
}

func (pbc *PhoneBookController) FindContactByName(ctx *gin.Context) {
	firstName := ctx.Query("firstname")
	lastName := ctx.Query("lastname")
	contact, err := pbc.Service.FindContactByName(firstName, lastName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, contact)
}

func (pbc *PhoneBookController) DeleteContact(ctx *gin.Context) {
	firstName := ctx.Query("firstname")
	lastName := ctx.Query("lastname")
	err := pbc.Service.DeleteContact(firstName, lastName)
	
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK ,"ok")
}

func (pbc *PhoneBookController) RegisterAllRoutes(rg *gin.RouterGroup){
	//base route
	phoneBookRoute := rg.Group("/phonebook")
	//api routes
	phoneBookRoute.POST("/add", pbc.AddContact)
	phoneBookRoute.PATCH("/edit", pbc.EditContact)
	phoneBookRoute.GET("/find-all-contacts", pbc.ShowAllContacts)
	phoneBookRoute.GET("/find-by-name/", pbc.FindContactByName)
	phoneBookRoute.DELETE("/delete", pbc.DeleteContact)

}