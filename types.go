package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ContactType struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
	Address      string `json:"address"`
}

type PhonebookOwner struct {
	client *mongo.Client
}

type PhoneBookController struct {
	Service PhoneBookService
}

type PhoneBookService interface {
	AddContact(*ContactType) error
	EditContact(string, string, *ContactType) error
	ShowAllContacts(int64) ([]*ContactType, error)
	FindContactByName(string, string) (*ContactType, error)
	FindContactByNumber(string) (*ContactType, error)
	DeleteContact(string, string) error
	DeleteAll() error
}

type ServiceParams struct {
	coll *mongo.Collection
	ctx   context.Context
}
