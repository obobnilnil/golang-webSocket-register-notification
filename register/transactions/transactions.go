package transactions

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITransaction interface {
	RegisterLogTransactions(Firstname string, Surname string, RegisterMessage string, JobTitle string, CompanyID string) error
}

type Transaction struct {
	collection *mongo.Collection
}

func NewTransaction(mongoDB *mongo.Database) ITransaction {
	return &Transaction{collection: mongoDB.Collection("logs")}
}

var Broadcast = make(chan bson.D)

func (t *Transaction) RegisterLogTransactions(Firstname string, Surname string, RegisterMessage string, JobTitle string, CompanyID string) error {
	registerLogCollection := t.collection.Database().Collection("registerLog")
	doc := bson.D{
		{Key: "firstname", Value: Firstname},
		{Key: "surname", Value: Surname},
		{Key: "registerMessage", Value: RegisterMessage},
		{Key: "jobTitle", Value: JobTitle},
		{Key: "timestamp", Value: time.Now().Unix()},
		{Key: "time", Value: time.Now().Format("2-Jan-06 03:04PM")},
		{Key: "companyID", Value: CompanyID},
	}
	Broadcast <- doc
	_, err := registerLogCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
