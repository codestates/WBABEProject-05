package util

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertOBJIDsToStrings(OBJIDs []primitive.ObjectID) []string {
	var IDs []string
	for _, OBJID := range OBJIDs {
		IDs = append(IDs, OBJID.Hex())
	}
	return IDs
}

func ConvertStringsToOBJIDs(stringIDs []string) ([]primitive.ObjectID, error) {
	var inID []primitive.ObjectID
	for _, STRID := range stringIDs {
		mID, err := primitive.ObjectIDFromHex(STRID)
		if err != nil {
			return nil, err
		}
		inID = append(inID, mID)
	}
	return inID, nil
}

func ConvertOBJIDToString(OBJIDs primitive.ObjectID) string {
	return OBJIDs.Hex()
}

func ConvertStringToOBJID(stringID string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(stringID)
}
