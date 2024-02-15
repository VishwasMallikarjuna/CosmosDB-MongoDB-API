package service

import (
	"errors"

	"azurepoc/dao"
	"azurepoc/model"
)

const (
//test_shard_id = "SH123"
)

// Create Batch...
func CreateBatch(payload model.BatchCreatePayload) (model.BatchBSON, error) {
	var (
		batch = payload.ConvertToBSON()
	)

	//Create user
	doc, err := dao.CreateBatch(batch)
	if err != nil {
		err = errors.New("failed to create Batch record")
		return doc, err
	}

	return doc, err
}

// Find Batchdetails By batchID ...
func BatchDetailsFindByBatchID(batchId string) (model.BatchBSON, error) {
	doc, err := dao.BatchDetailsFindByBatchID(batchId)
	return doc, err
}

// Find Batchdetails By TenantID ...
func BatchDetailsFindByTenantID(tenantId string) (model.BatchBSON, error) {
	doc, err := dao.BatchDetailsFindByTenantID(tenantId)
	return doc, err
}

//Update Batch

func BatchUpdate(payload model.BatchCreatePayload) (model.BatchBSON, error) {
	//func EmployeeUpdate(payload model.EmployeeCreatePayload) (model.EmployeeBSON, error) {
	var (
		batch = payload.ConvertToBSON()
	)

	//Create user
	doc, err := dao.BatchUpdate(batch)
	if err != nil {
		err = errors.New("failed to update batch details")
		return doc, err
	}

	return doc, err
}

//Update Batch

func UpdateBatchStatus(payload model.BatchCreatePayload, status string) (model.BatchBSON, error) {

	var (
		batch = payload.ConvertToBSON()
	)

	//Create user
	doc, err := dao.BatchStatusUpdate(batch, status)
	if err != nil {
		err = errors.New("failed to update batch status details")
		return doc, err
	}

	return doc, err
}
func BatchDeleteByBatchId(batchId string) error {
	err := dao.BatchDeleteByBatchId(batchId)
	return err
}
