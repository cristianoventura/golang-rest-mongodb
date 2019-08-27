package database

import (
	"context"
	"golang-rest-crud/structs"

	"go.mongodb.org/mongo-driver/bson"
)

// AddPerson - Creates a new document in people collection
func (db *DB) AddPerson(person structs.Person) (interface{}, error) {
	collection := db.Collection("people")
	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FindPerson - Returns the document by the specified id
func (db *DB) FindPerson(id string) (interface{}, error) {
	collection := db.Collection("people")
	filter := bson.D{{"_id", id}}
	var res structs.Person
	err := collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}