package dao

import (
	"context"
	"fmt"

	"azurepoc/model"
	"azurepoc/module/database"

	"go.mongodb.org/mongo-driver/bson"
)

// EmployeeCreate ...
func CreateBatch(doc model.BatchBSON) (model.BatchBSON, error) {
	var (
		batchCol = database.BatchCol()
		ctx      = context.Background()
	)

	// Insert one
	_, err := batchCol.InsertOne(ctx, doc)
	return doc, err
}

// BatchDetailsFindByID ...
func BatchDetailsFindByBatchID(batchId string) (model.BatchBSON, error) {
	var (
		batchCol = database.BatchCol()
		ctx      = context.Background()
		result   model.BatchBSON
		filter   = bson.M{"batchId": batchId}
	)

	// Find
	err := batchCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}

// BatchDetailsFindBy TenantID ...
func BatchDetailsFindByTenantID(tenantId string) (model.BatchBSON, error) {
	var (
		batchCol = database.BatchCol()
		ctx      = context.Background()
		result   model.BatchBSON
		filter   = bson.M{"tenantId": tenantId}
	)

	// Find
	err := batchCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}

// BatchUpdate ...
func BatchUpdate(doc model.BatchBSON) (model.BatchBSON, error) {
	var (
		batchCol     = database.BatchCol()
		ctx          = context.Background()
		returnResult model.BatchBSON
		filter       = bson.M{"batchId": doc.BATCHID}
	)

	//objId, _ := primitive.ObjectIDFromHex(doc.EMPID)
	update := bson.M{"recordCount": doc.RECORDCOUNT, "status": doc.STATUS}
	// update
	result, err := batchCol.UpdateOne(ctx, bson.M{"batchId": doc.BATCHID}, bson.M{"$set": update})
	fmt.Println("Update result count", result.ModifiedCount)
	if result.ModifiedCount == 1 {
		err := batchCol.FindOne(ctx, filter).Decode(&returnResult)
		return returnResult, err
	}
	return doc, err
}

func BatchStatusUpdate(doc model.BatchBSON, status string) (model.BatchBSON, error) {
	var (
		batchCol     = database.BatchCol()
		ctx          = context.Background()
		returnResult model.BatchBSON
		filter       = bson.M{"batchId": doc.BATCHID}
	)

	//objId, _ := primitive.ObjectIDFromHex(doc.EMPID)
	update := bson.M{"status": status}
	// update
	result, err := batchCol.UpdateOne(ctx, bson.M{"batchId": doc.BATCHID, "tenantId": doc.TENANTID}, bson.M{"$set": update})
	fmt.Println("Update result count", result.ModifiedCount)
	if result.ModifiedCount == 1 {
		err := batchCol.FindOne(ctx, filter).Decode(&returnResult)
		return returnResult, err
	}
	return doc, err
}
func BatchDeleteByBatchId(batchId string) error {

	var (
		userCol = database.BatchCol()
		ctx     = context.Background()
		//result  model.EmployeeBSON
		filter = bson.M{"batchId": batchId}
	)

	// Find
	result, err := userCol.DeleteOne(ctx, filter)
	fmt.Println("deleted result count", result.DeletedCount)

	return err
}
