package contacts

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sachin120/contact_app/datasources/mongodb/contacts_db"
	"github.com/sachin120/contact_app/utils/errors"
)

// Get ...
func (contact *Contact) Get() *errors.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := contacts_db.Collection.FindOne(ctx, bson.M{"id": contact.ID}).Decode(&contact)
	if err != nil {
		log.Println(err)
		return &errors.RestErr{
			Message: fmt.Sprintf("Result not found for id %v", contact.ID),
			Status:  404,
			Error:   "DB error",
		}
	}
	return nil
}

// Save ...
func (contact *Contact) Save() *errors.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := contacts_db.Collection.InsertOne(ctx, contact)
	if err != nil {
		return &errors.RestErr{} // appropriate msg
	}
	log.Printf("Inserted doc %v with id %v\n", contact, result.InsertedID)
	return nil
}

// Delete ...
func (contact *Contact) Delete() *errors.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := contacts_db.Collection.DeleteOne(ctx, bson.M{"id": contact.ID})
	if err != nil {
		return &errors.RestErr{} // appropriate msg
	}
	log.Printf("Deleted doc count %v\n", result.DeletedCount)
	return nil
}
