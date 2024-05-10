package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func NewPhonebookOwner(c *mongo.Client) *PhonebookOwner {
	return &PhonebookOwner{
		client: c,
	}
}

func (owner *PhonebookOwner) start(coll *mongo.Collection ) error {

	phoneBookcount, err := coll.CountDocuments(context.TODO(), bson.M{})
    if err != nil {
        return err
    }

	if phoneBookcount == 0 {
	
	myContact := ContactType{
		FirstName: "my",
		LastName: "phone",
		PhoneNumber: "0501234567",
		Address: "my-address",
	}
	_, err := coll.InsertOne(context.TODO(), myContact)
    if err != nil {
        return err
    }
    fmt.Println("Contact inserted successfully.")
	}
    return nil
}


func main(){

	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://meytar:phonebook@localhost:27017/"))
	if err != nil {
		panic(err)
	}
	
	fmt.Println("mongo connection is on")
	
	phoneOwner := NewPhonebookOwner(client)
	
	mdbColl := phoneOwner.client.Database("phonebooks").Collection("myphonebook")
	
	phoneBookService:= NewPhoneBookService(mdbColl,ctx)

	phoneBookController := New(phoneBookService)

	phoneOwner.start(mdbColl)
	server := gin.Default()
	basepath := server.Group("/app")
	phoneBookController.RegisterAllRoutes(basepath)
	server.Run(":3000")


}