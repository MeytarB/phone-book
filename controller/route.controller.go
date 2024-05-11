package controller

import "github.com/gin-gonic/gin"

func (pbc *PhoneBookController) RegisterAllRoutes(rg *gin.RouterGroup) {
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