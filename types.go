package main

import (
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


