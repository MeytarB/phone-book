package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func NewPhoneBookService(coll *mongo.Collection, ctx context.Context) PhoneBookService {
	return &ServiceParams{
		coll: coll,
		ctx: ctx,
	}
}


func (sp *ServiceParams) AddContact(newContact *ContactType) error {
	_ , err:= sp.FindContactByName(newContact.FirstName, newContact.LastName)
	
	//if the contact name already exists:
	if err == nil{
		err = errors.New("error: contact already exists")
	}
// if it didnt find the required name in contacts - it will add it 
	if err== mongo.ErrNoDocuments {
	_, err = sp.coll.InsertOne(sp.ctx, newContact)
	fmt.Println("contact was added successfully!")
}

//TODO - add a check for the phone number
	
	return err
}


func (sp *ServiceParams) EditContact(contact *ContactType) error {
	
	return nil
}


func (sp *ServiceParams) ShowAllContacts() ([]*ContactType, error) {
	var results []*ContactType
	filter := bson.M{}
	cursor, err := sp.coll.Find(sp.ctx, filter)
    
	if err != nil {
        return nil, err
    }

	//go over all results and decoding it into the empty array that was declared before
	for cursor.Next(sp.ctx) {
		var contact ContactType
		err := cursor.Decode(&contact)
		if err != nil {
			return nil, err
		}
		results = append(results, &contact)
	}

	if len(results) == 0 {
		return nil, errors.New("there are no contacts")
	}

	cursor.Close(sp.ctx)

	return results, nil
}

func (sp *ServiceParams) FindContactByName(firstName string, lastName string) (*ContactType, error) {
	var result *ContactType	
	filter := bson.M{"firstname": firstName, "lastname": lastName}
	err := sp.coll.FindOne(context.Background(), filter).Decode(&result)
	
	return result, err
}

func (sp *ServiceParams) FindContactByNumber(phoneNumber string) (*ContactType, error) {
	var result *ContactType	
	filter := bson.M{"phonenumber": phoneNumber}
	err := sp.coll.FindOne(context.Background(), filter).Decode(&result)
	
	return result, err
}

func (sp *ServiceParams) DeleteContact(firstName string, lastName string) error {	
	filter := bson.M{"firstname": firstName, "lastname": lastName}
	result, _ := sp.coll.DeleteOne(sp.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("error: contact was not found")
	}
	return nil
}

func (sp *ServiceParams) DeleteAll() error {
	filter := bson.M{}
	_ , err := sp.coll.DeleteMany(sp.ctx,filter)

	if err != nil {
		return errors.New("error: contacts were not deleted")
	}

	return nil
}

