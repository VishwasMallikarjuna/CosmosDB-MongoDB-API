package database

import "go.mongodb.org/mongo-driver/mongo"

const (
	testsharindcollection = "test_sharding"
)

// PlayerCol ...
func EmployeeAddressCol() *mongo.Collection {
	return db.Collection(testsharindcollection)
}
