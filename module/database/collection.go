package database

import "go.mongodb.org/mongo-driver/mongo"

const (
	testcollection = "test"
)

// PlayerCol ...
func PlayerCol() *mongo.Collection {
	return db.Collection(testcollection)
}

// EmployeeCol ...
func EmployeeCol() *mongo.Collection {
	return db.Collection(testcollection)
}

// EmployeeCol ...
func BatchCol() *mongo.Collection {
	return db.Collection(testcollection)
}
