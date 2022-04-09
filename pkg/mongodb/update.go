package pkg

import (
	"log"
)

func Update(collection string, filter interface{}, customUpdateStruct interface{}) (int64, error) {
	var modifiedCount int64
	collectionName := Database.Collection(collection)
	result, insertErr := collectionName.UpdateOne(contextForDb, filter, customUpdateStruct)
	if insertErr != nil {
		log.Println("Mongo Update Failed : ", insertErr)
		return modifiedCount, insertErr
	}
	return result.ModifiedCount, nil
}
