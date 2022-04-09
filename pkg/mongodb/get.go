package pkg

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetById(collection string, docID string) interface{} {

	var result interface{}

	collectionName := Database.Collection(collection)

	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	err := collectionName.FindOne(contextForDb, filter).Decode(&result)
	if err != nil {
		log.Println("Error getting document by Id")
		log.Println(err)
	}

	return result
}

func GetByField(collection string, fieldName, fieldValue string) []bson.M {

	var result []bson.M

	collectionName := Database.Collection(collection)
	cursor, err := collectionName.Find(contextForDb, bson.M{fieldName: fieldValue})
	if err != nil {
		log.Println("Error getting document by field name")
		log.Println(err)
	}
	if err = cursor.All(contextForDb, &result); err != nil {
		log.Println("Error getting all cursor by field in mongodb")
		log.Println(err)
	}

	return result
}
