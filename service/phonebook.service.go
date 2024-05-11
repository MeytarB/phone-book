package service

import (
	"context"
	"errors"
	"strconv"
    
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/MeytarB/phone-book/controller"
)


func NewPhoneBookService(coll *mongo.Collection, ctx context.Context) PhoneBookService {
	return &ServiceParams{
		coll: coll,
		ctx: ctx,
	}
}

func (sp *ServiceParams) isValidNumber(phoneNumber string) bool {
	_, err := strconv.Atoi(phoneNumber)
    return err == nil
}


func (sp *ServiceParams) checkIfContactExist(newContact *ContactType) error {
	//if the contact name already exists:
	_ , err:= sp.FindContactByName(newContact.FirstName, newContact.LastName)
	if err == nil{
		err = errors.New("error: contact full name already exists")
		return err
	}

	//if the contact number already exists:
	_ , err= sp.FindContactByNumber(newContact.PhoneNumber)
	if err == nil{
		err = errors.New("error: contact number already exists")
		return err
	}
	
	return err
}


func (sp *ServiceParams) AddContact(newContact *ContactType) error {
	if !sp.isValidNumber(newContact.PhoneNumber){
		return errors.New("error: not a valid number")
	}
	err := sp.checkIfContactExist(newContact)
	if err== mongo.ErrNoDocuments {
		_, err = sp.coll.InsertOne(sp.ctx, newContact)	
	}

	return err
}


func (sp *ServiceParams) EditContact(firstName string ,lastName string , updatedContact *ContactType) error {

	filter := bson.D{primitive.E{Key: "firstname", Value: firstName}, 
			  primitive.E{Key: "lastname", Value: lastName}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
				primitive.E{Key: "firstname", Value: updatedContact.FirstName}, 
				primitive.E{Key: "lastname", Value: updatedContact.LastName},
				primitive.E{Key: "phonenumber", Value: updatedContact.PhoneNumber},
				primitive.E{Key: "address", Value: updatedContact.Address}}}}
	
	if !sp.isValidNumber(updatedContact.PhoneNumber){
		return errors.New("error: not a valid number")
	}
	// we will find the original contact
	originalContact, err := sp.FindContactByName(firstName,lastName)
	if err == mongo.ErrNoDocuments{
		return err
	}
	// making sure that if the number is changed - the new one doesnt already exist
	if updatedContact.PhoneNumber != originalContact.PhoneNumber{
		_, err = sp.FindContactByNumber(updatedContact.PhoneNumber)
		if err == nil{
			err = errors.New("error: new phone number already exist, cannot update")
			return err
		}
	}	
	// making sure that if the name is changed - the new full name doesnt already exist
	if updatedContact.FirstName != originalContact.FirstName || 
	   updatedContact.LastName != originalContact.LastName {
		_, err = sp.FindContactByName(updatedContact.FirstName, updatedContact.LastName)
		if err == nil{
			err = errors.New("error: new full name already exist, cannot update")
			return err
		}
	}

	var result *mongo.UpdateResult
	result, err = sp.coll.UpdateOne(sp.ctx, filter, update)
	if result == nil{
		return err
	}

	return nil
}


func (sp *ServiceParams) ShowAllContacts(page int64) ([]*ContactType, error) {
	var results []*ContactType
	paging := 10
	startingPoint := int64(paging) * (page -1)
	filter := bson.M{}
	//find the 10 first results from the page we want
	cursor, err := sp.coll.Find(sp.ctx, filter, options.Find().SetSkip(startingPoint).SetLimit(int64(paging)))
    
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

