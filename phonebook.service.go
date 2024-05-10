package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type PhoneBookService interface {
	AddContact(*ContactType) error
	EditContact(*ContactType) error
	FindAllContacts() ([]*ContactType, error)
	FindUserByName(*string) (*ContactType, error)
	DeleteContact(*string) error
}

type ServiceParams struct {
	coll *mongo.Collection
	ctx            context.Context
}


func NewPhoneBookService(coll *mongo.Collection, ctx context.Context) PhoneBookService {
	return &ServiceParams{
		coll: coll,
		ctx: ctx,
	}
}


func (sp *ServiceParams) AddContact(newContact *ContactType) error {
	_, err := sp.coll.InsertOne(sp.ctx, newContact)
	fmt.Println("Contact was added successfully!")
	return err
}


func (sp *ServiceParams) EditContact(contact *ContactType) error {
	
	return nil
}


func (sp *ServiceParams) FindAllContacts() ([]*ContactType, error) {
	
	return nil , nil
}

func (sp *ServiceParams) FindUserByName(name *string) (*ContactType, error) {
	
	return nil, nil
}

func (sp *ServiceParams) DeleteContact(name *string) error {
	
	return nil
}

// func (owner *PhonebookOwner) findAllContacts() ([]ContactType, error) {
//     coll := owner.client.Database("phonebooks").Collection("myphonebook")

//     // Define an empty filter to retrieve all documents
//     filter := bson.M{}

//     // Declare a slice to store the results
//     var results []ContactType

//     // Find all documents in the collection
//     cursor, err := coll.Find(context.TODO(), filter)
//     if err != nil {
//         return nil, err
//     }

//     // Iterate through the cursor and decode each document into the results slice
//     for cursor.Next(context.Background()) {
//         var result ContactType
//         err := cursor.Decode(&result)
//         if err != nil {
//             return nil, err
//         }
//         results = append(results, result)
//     }

//     // Close the cursor once done
//     cursor.Close(context.Background())

//     return results, nil
// }
