package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func NewPhonebookOwner(c *mongo.Client) *PhonebookOwner {
	return &PhonebookOwner{
		client: c,
	}
}

func (owner *PhonebookOwner) start(sp PhoneBookService) {
	
	myContact := &ContactType{
		FirstName: "my",
		LastName: "phone",
		PhoneNumber: "0501234567",
		Address: "my-address",
	}

	if sp.AddContact(myContact) == nil{
    
    	fmt.Println("my contact details were inserted successfully.")
	} 
}

func main(){
//creating the mongo connection
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://meytar:phonebook@localhost:27017/"))
	if err != nil {
		panic(err)
	}
	
	fmt.Println("mongo connection is on")
	// creating the client and initialize its number	
	phoneOwner := NewPhonebookOwner(client)
	mdbColl := phoneOwner.client.Database("phonebooks").Collection("myphonebook")
	phoneBookService:= NewPhoneBookService(mdbColl,ctx)
	phoneBookController := New(phoneBookService)
	phoneOwner.start(phoneBookService)

	//creating the server and routes
	server := gin.Default()
	basepath := server.Group("/app")
	phoneBookController.RegisterAllRoutes(basepath)
	server.Run(":3000")

	

}