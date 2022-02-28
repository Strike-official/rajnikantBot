package pkg

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(collection string, customInsertStruct interface{}) string {
	var resultID string

	collectionName := Database.Collection(collection)
	result, insertErr := collectionName.InsertOne(contextForDb, customInsertStruct)
	if insertErr != nil {
		log.Println("Mongo Insert Failed : ", insertErr)
		return resultID
	}

	resultID = result.InsertedID.(primitive.ObjectID).Hex()

	return resultID
}

//http://ec2-18-218-96-97.us-east-2.compute.amazonaws.com
