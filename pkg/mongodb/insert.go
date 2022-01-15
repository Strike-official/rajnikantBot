package pkg

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(collection string, customInsertStruct interface{}) string {
	var resultID string

	collectionName := Database.Collection(collection)
	result, insertErr := collectionName.InsertOne(contextForDb, customInsertStruct)
	if insertErr != nil {
		log.Println("Insert in mongodb failed")
		log.Println(insertErr)
	} else {
		fmt.Println("InsertOne() API result:", result)
		newID := result.InsertedID
		fmt.Println("Inserted new document to mongodb : ", newID)
		resultID = newID.(primitive.ObjectID).Hex()
	}
	return resultID
}
