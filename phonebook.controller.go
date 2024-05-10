package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhoneBookController struct {
	Service PhoneBookService
}

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
	}
	ctx.JSON(http.StatusOK ,"ok")
}

func (pbc *PhoneBookController) EditContact(ctx *gin.Context) {

	ctx.JSON(200,"ok")
}

func (pbc *PhoneBookController) FindAllContacts(ctx *gin.Context) {

	ctx.JSON(200,"ok")
}

func (pbc *PhoneBookController) FindUserByName(ctx *gin.Context) {

	ctx.JSON(200,"ok")
}

func (pbc *PhoneBookController) DeleteContact(ctx *gin.Context) {

	ctx.JSON(200,"ok")
}

func (pbc *PhoneBookController) RegisterAllRoutes(rg *gin.RouterGroup){
	//base route
	phoneBookRoute := rg.Group("/phonebook")
	//api routes
	phoneBookRoute.POST("/add", pbc.AddContact)
	phoneBookRoute.PATCH("/edit", pbc.EditContact)
	phoneBookRoute.GET("/find-all-contacts", pbc.FindAllContacts)
	phoneBookRoute.GET("/find-by-name/:fullname", pbc.FindUserByName)
	phoneBookRoute.DELETE("/delete/:fullname", pbc.DeleteContact)

}