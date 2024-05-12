package controller

import (
	"github.com/MeytarB/phone-book/service"
)

type PhoneBookController struct {
	service service.PhoneBookService
}


func New(service service.PhoneBookService) PhoneBookController {
	return PhoneBookController{
		service: service,
	}
}










