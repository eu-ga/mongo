//package mongo implements CRUD system with mongoDB
//NOTE: passing nil as the filter matches all documents in the collection
package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne(c *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {

	res, err := c.InsertOne(context.TODO(), document, opts...)
	if err != nil {
		return nil, errors.New("cannot insert one into mongoDB: " + err.Error())
	}
	return res, nil
}

func InsertMany(c *mongo.Collection, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {

	res, err := c.InsertMany(context.TODO(), documents, opts...)
	if err != nil {
		return nil, errors.New("cannot insert many into mongoDB: " + err.Error())
	}
	return res, nil
}

func UpdateOne(c *mongo.Collection, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {

	updateResult, err := c.UpdateOne(context.TODO(), filter, update, opts...)
	if err != nil {
		return nil, errors.New("cannot update one in mongoDB: " + err.Error())
	}
	return updateResult, nil
}

//FindOne finds single value in the database and returns error
//result is a pointer
func FindOne(c *mongo.Collection, filter, result interface{}, opts ...*options.FindOneOptions) error {

	err := c.FindOne(context.TODO(), filter, opts...).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

//Find finds all documents that match filter and returns error
//result is a pointer
func Find(c *mongo.Collection, filter, results []interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {

	cur, err := c.Find(context.TODO(), filter, opts...)
	if err != nil {
		return nil, errors.New("cannot find all in mongoDB: " + err.Error())
	}

	return cur, nil
}

func DeleteOne(c *mongo.Collection, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {

	f := bson.M{}
	deleteResult, err := c.DeleteMany(context.TODO(), f, opts...)
	if err != nil {
		return nil, errors.New("cannot delete from mongoDB: " + err.Error())
	}
	return deleteResult, nil
}
