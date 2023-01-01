package util

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertObjIDsToStrings(objIDs []primitive.ObjectID) []string {
	var IDs []string
	for _, objID := range objIDs {
		IDs = append(IDs, objID.Hex())
	}
	return IDs
}

func ConvertStringsToObjIDs(stringIDs []string) ([]primitive.ObjectID, error) {
	var inID []primitive.ObjectID
	for _, strID := range stringIDs {
		mID, err := primitive.ObjectIDFromHex(strID)
		if err != nil {
			return nil, err
		}
		inID = append(inID, mID)
	}
	return inID, nil
}
