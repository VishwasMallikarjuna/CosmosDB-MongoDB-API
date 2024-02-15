package dao

import (
	"context"
	"fmt"

	"azurepoc/model"
	"azurepoc/module/database"

	"go.mongodb.org/mongo-driver/bson"
)

// EmployeeCreate ...
func EmployeeCreate(doc model.EmployeeBSON) (model.EmployeeBSON, error) {
	var (
		userCol = database.EmployeeCol()
		ctx     = context.Background()
	)

	// Insert one
	_, err := userCol.InsertOne(ctx, doc)
	return doc, err
}

// EmployeeAddressCreate ...
func EmployeeAddressCreate(doc model.EmployeeAddressBSON) (model.EmployeeAddressBSON, error) {
	var (
		addressCol = database.EmployeeAddressCol()
		ctx        = context.Background()
	)

	// Insert one
	_, err := addressCol.InsertOne(ctx, doc)
	return doc, err
}

// EmployeeFindByID ...
func EmployeeFindByID(empId string) (model.EmployeeBSON, error) {
	var (
		userCol = database.EmployeeCol()
		ctx     = context.Background()
		result  model.EmployeeBSON
		filter  = bson.M{"empId": empId}
	)

	// Find
	err := userCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}

// EmployeeFindByID ...
func EmployeeFindByAddressID(addressId string) (model.EmployeeAddressBSON, error) {
	var (
		userCol = database.EmployeeAddressCol()
		ctx     = context.Background()
		result  model.EmployeeAddressBSON
		filter  = bson.M{"addressId": addressId}
	)

	fmt.Println("EmployeeFindByAddressID ", addressId)
	// Find
	err := userCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}

// EmployeeUpdate ...
func EmployeeUpdate(doc model.EmployeeBSON) (model.EmployeeBSON, error) {
	var (
		userCol = database.EmployeeCol()
		ctx     = context.Background()
	)

	//objId, _ := primitive.ObjectIDFromHex(doc.EMPID)
	update := bson.M{"name": doc.Name, "department": doc.DEPARTMENT}
	// update
	result, err := userCol.UpdateOne(ctx, bson.M{"empId": doc.EMPID}, bson.M{"$set": update})
	fmt.Println("Update result count", result.ModifiedCount)
	return doc, err
}

// EmployeeUpdate ...
func EmployeeAddressUpdate(doc model.EmployeeAddressBSON) (model.EmployeeAddressBSON, error) {
	var (
		userCol = database.EmployeeAddressCol()
		ctx     = context.Background()
	)

	//objId, _ := primitive.ObjectIDFromHex(doc.EMPID)
	update := bson.M{"empId": doc.EMPID, "area": doc.AREA, "city": doc.CITY, "state": doc.STATE,
		"country": doc.COUNTRY}
	// update
	result, err := userCol.UpdateOne(ctx, bson.M{"addressId": doc.ADDRESSID}, bson.M{"$set": update})
	fmt.Println("Update result count", result.ModifiedCount)
	return doc, err
}
func EmployeeDelete(empId string) error {

	var (
		userCol = database.EmployeeCol()
		ctx     = context.Background()
		//result  model.EmployeeBSON
		filter = bson.M{"empId": empId}
	)

	// Find
	result, err := userCol.DeleteOne(ctx, filter)
	fmt.Println("deleted result count", result.DeletedCount)

	return err
}

func EmployeeAddressDelete(addressId string) error {

	var (
		userCol = database.EmployeeAddressCol()
		ctx     = context.Background()
		//result  model.EmployeeBSON
		filter = bson.M{"addressId": addressId}
	)

	// Find
	result, err := userCol.DeleteOne(ctx, filter)
	fmt.Println("deleted result count", result.DeletedCount)

	return err
}
