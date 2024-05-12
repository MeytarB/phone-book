package mongo

import (
	"context"
	"errors"
	"strconv"

	"github.com/MeytarB/phone-book/service"
	"github.com/MeytarB/phone-book/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	client *mongo.Client
	coll *mongo.Collection
	ctx context.Context
}


func Init() service.PhoneBookService {
	ctx := context.TODO()
//todo - change localhost to db
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://meytar:phonebook@db:27017/"))
	if err != nil {
		panic(err)
	}

	mdbColl := client.Database("phonebooks").Collection("myphonebook")
	phoneOwner := &MongoService{client: client ,coll: mdbColl, ctx: ctx }
	
	phoneOwner.addOwnerDetails()
	
	return phoneOwner
}


func (ms *MongoService) AddContact(newContact *types.Contact) error {
	if !ms.isValidNumber(newContact.PhoneNumber){
		return errors.New(string(types.ValidNumberError))
	}
	err := ms.checkIfContactExist(newContact)
	if err== mongo.ErrNoDocuments {
		_, err = ms.coll.InsertOne(ms.ctx, newContact)	
	}

	return err
}


func (ms *MongoService) EditContact(firstName string ,lastName string , updatedContact *types.Contact) error {

	filter := bson.D{primitive.E{Key: "firstname", Value: firstName}, 
			  primitive.E{Key: "lastname", Value: lastName}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
				primitive.E{Key: "firstname", Value: updatedContact.FirstName}, 
				primitive.E{Key: "lastname", Value: updatedContact.LastName},
				primitive.E{Key: "phonenumber", Value: updatedContact.PhoneNumber},
				primitive.E{Key: "address", Value: updatedContact.Address}}}}
	
	if !ms.isValidNumber(updatedContact.PhoneNumber){
		return errors.New(string(types.ValidNumberError))
	}
	// we will find the original contact
	originalContact, err := ms.FindContactByName(firstName,lastName)
	if err == mongo.ErrNoDocuments{
		return err
	}
	// making sure that if the number is changed - the new one doesnt already exist
	if updatedContact.PhoneNumber != originalContact.PhoneNumber{
		_, err = ms.FindContactByNumber(updatedContact.PhoneNumber)
		if err == nil{
			err = errors.New(string(types.DuplicateNumberError))
			return err
		}
	}	
	// making sure that if the name is changed - the new full name doesnt already exist
	if updatedContact.FirstName != originalContact.FirstName || 
	   updatedContact.LastName != originalContact.LastName {
		_, err = ms.FindContactByName(updatedContact.FirstName, updatedContact.LastName)
		if err == nil{
			err = errors.New(string(types.DuplicateNameError))
			return err
		}
	}

	var result *mongo.UpdateResult
	result, err = ms.coll.UpdateOne(ms.ctx, filter, update)
	if result == nil{
		return err
	}

	return nil
}


func (ms *MongoService) ShowAllContacts(page int64) ([]*types.Contact, error) {
	var results []*types.Contact
	paging := 10
	startingPoint := int64(paging) * (page -1)
	filter := bson.M{}
	//find the 10 first results from the page we want
	cursor, err := ms.coll.Find(ms.ctx, filter, options.Find().SetSkip(startingPoint).SetLimit(int64(paging)))
    
	if err != nil {
        return nil, err
    }

	//go over all results and decoding it into the empty array that was declared before
	for cursor.Next(ms.ctx) {
		var contact types.Contact
		err := cursor.Decode(&contact)
		if err != nil {
			return nil, err
		}
		results = append(results, &contact)
	}

	if len(results) == 0 {
		return nil, errors.New("there are no contacts")
	}

	cursor.Close(ms.ctx)

	return results, nil
}

func (ms *MongoService) FindContactByName(firstName string, lastName string) (*types.Contact, error) {
	var result *types.Contact	
	filter := bson.M{"firstname": firstName, "lastname": lastName}
	err := ms.coll.FindOne(context.Background(), filter).Decode(&result)
	
	return result, err
}

func (ms *MongoService) FindContactByNumber(phoneNumber string) (*types.Contact, error) {
	var result *types.Contact	
	filter := bson.M{"phonenumber": phoneNumber}
	err := ms.coll.FindOne(context.Background(), filter).Decode(&result)
	
	return result, err
}

func (ms *MongoService) DeleteContact(firstName string, lastName string) error {	
	filter := bson.M{"firstname": firstName, "lastname": lastName}
	result, _ := ms.coll.DeleteOne(ms.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New(string(types.NumberNotFoundError))
	}
	return nil
}

func (ms *MongoService) DeleteAll() error {
	filter := bson.M{}
	_ , err := ms.coll.DeleteMany(ms.ctx,filter)

	return err;
}


func (ms *MongoService) isValidNumber(phoneNumber string) bool {
	_, err := strconv.Atoi(phoneNumber)
    return err == nil
}


func (ms *MongoService) checkIfContactExist(newContact *types.Contact) error {
	//if the contact name already exists:
	_ , err:= ms.FindContactByName(newContact.FirstName, newContact.LastName)
	if err == nil{
		err = errors.New("error: contact full name already exists")
		return err
	}

	//if the contact number already exists:
	_ , err= ms.FindContactByNumber(newContact.PhoneNumber)
	if err == nil{
		err = errors.New("error: contact number already exists")
		return err
	}
	
	return err
}


func (ms *MongoService) addOwnerDetails() {
	myContact := &types.Contact{
		FirstName: "my",
		LastName: "phone",
		PhoneNumber: "0501234567",
		Address: "my-address",
	}

	ms.AddContact(myContact)
}