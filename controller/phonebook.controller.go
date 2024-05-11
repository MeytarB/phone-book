package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func New(service PhoneBookService) PhoneBookController {
	return PhoneBookController{
		service: service,
	}
}

func (pbc *PhoneBookController) AddContact(ctx *gin.Context) {
	var contact ContactType
	if err:= ctx.ShouldBindJSON(&contact); err != nil{
		ctx.JSON((http.StatusBadRequest), gin.H{"message": err.Error()})
	}
	err:= pbc.service.AddContact(&contact)
	if err != nil{
		ctx.JSON((http.StatusBadGateway), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK , gin.H{"message": "OK"})
}

func (pbc *PhoneBookController) EditContact(ctx *gin.Context) {
	var contact ContactType
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

func (pbc *PhoneBookController) RegisterAllRoutes(rg *gin.RouterGroup){
	//base route
	phoneBookRoute := rg.Group("/phonebook")
	//api routes
	phoneBookRoute.POST("/add", pbc.AddContact)
	phoneBookRoute.PATCH("/edit/:firstname/:lastname", pbc.EditContact)
	phoneBookRoute.GET("/find-all-contacts/:page", pbc.ShowAllContacts)
	phoneBookRoute.GET("/find-by-name/", pbc.FindContactByName)
	phoneBookRoute.GET("/find-by-number/", pbc.FindContactByNumber)
	phoneBookRoute.DELETE("/delete-contact", pbc.DeleteContact)
	phoneBookRoute.DELETE("/delete-all", pbc.DeleteAll)

}