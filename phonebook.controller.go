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
	ctx.JSON(http.StatusOK , gin.H{"message": "OK"})
}

func (pbc *PhoneBookController) EditContact(ctx *gin.Context) {
	var contact ContactType
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := pbc.Service.EditContact(&contact)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK , gin.H{"message": "OK"})
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

func (pbc *PhoneBookController) FindContactByNumber(ctx *gin.Context) {
	phoneNumber := ctx.Query("phonenumber")
	contact, err := pbc.Service.FindContactByNumber(phoneNumber)
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

	ctx.JSON(http.StatusOK , gin.H{"message": "OK"})
}


func (pbc *PhoneBookController) DeleteAll(ctx *gin.Context) {

	err := pbc.Service.DeleteAll()
	
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK , gin.H{"message": "OK"})
}

func (pbc *PhoneBookController) RegisterAllRoutes(rg *gin.RouterGroup){
	//base route
	phoneBookRoute := rg.Group("/phonebook")
	//api routes
	phoneBookRoute.POST("/add", pbc.AddContact)
	phoneBookRoute.PATCH("/edit", pbc.EditContact)
	phoneBookRoute.GET("/find-all-contacts", pbc.ShowAllContacts)
	phoneBookRoute.GET("/find-by-name/", pbc.FindContactByName)
	phoneBookRoute.GET("/find-by-number/", pbc.FindContactByNumber)
	phoneBookRoute.DELETE("/delete-contact", pbc.DeleteContact)
	phoneBookRoute.DELETE("/delete-all", pbc.DeleteAll)

}