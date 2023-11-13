package repo

import (
	"context"
	"fmt"

	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type SupplierRepo interface {
	service.SupplierRepo
}

type supplierRepo struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

func NewSupplierRepo(svc *dynamodb.DynamoDB, tableName string) SupplierRepo {
	return &supplierRepo{
		svc:       svc,
		tableName: tableName,
	}
}

func (r *supplierRepo) Create(ctx context.Context, sup *service.Supplier) error {
	usr, err := dynamodbattribute.MarshalMap(sup)
	if err != nil {
		return fmt.Errorf("cannot marshal report: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      usr,
	}
	_, err = r.svc.PutItemWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return fmt.Errorf("failed to write item: %v - %v", aerr.Code(), aerr.Message())
		}

		return fmt.Errorf("failed to write item: %v", err)
	}

	return nil
}

func (r *supplierRepo) GetSupplierByEmail(ctx context.Context, email string) (*service.Supplier, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(r.tableName),
		IndexName:              aws.String(userEmailIndex),
		KeyConditionExpression: aws.String("Email = :email"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {
				S: aws.String(email),
			},
		},
	}

	result, err := r.svc.QueryWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return nil, fmt.Errorf("failed to get item: %v - %v", aerr.Code(), aerr.Message())
		}
		return nil, fmt.Errorf("failed to get item: %v", err)
	}

	if len(result.Items) == 0 {
		return nil, nil
	}

	var supplier *service.Supplier
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &supplier)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB item: %v", err)
	}
	return supplier, nil
}

func (r *supplierRepo) GetSupplierByID(ctx context.Context, id string) (*service.Supplier, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
	}

	result, err := r.svc.GetItemWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return nil, fmt.Errorf("failed to get item: %v - %v", aerr.Code(), aerr.Message())
		}
		return nil, fmt.Errorf("failed to get item: %v", err)
	}

	if result.Item == nil {
		return nil, nil
	}

	var supplier *service.Supplier
	err = dynamodbattribute.UnmarshalMap(result.Item, &supplier)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB item: %v", err)
	}

	return supplier, nil
}

func (r *supplierRepo) DeleteSupplierByID(ctx context.Context, id string) error {
	u, err := r.GetSupplierByID(ctx, id)
	if err != nil {
		return err
	}

	if u == nil {
		return nil
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(u.ID),
			},
		},
	}

	_, err = r.svc.DeleteItemWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return fmt.Errorf("failed to delete item: %v - %v", aerr.Code(), aerr.Message())
		}
		return fmt.Errorf("failed to delete item: %v", err)
	}

	return nil
}

func (r *supplierRepo) UpdateSupplierByID(ctx context.Context, sup *service.Supplier) error {
	u, err := r.GetSupplierByID(ctx, sup.ID)
	if err != nil {
		return err
	}

	if u == nil {
		return nil
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(sup.ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":e": {
				S: aws.String(sup.Email),
			},
			":p": {
				S: aws.String(sup.PhoneNumber),
			},
			":f": {
				S: aws.String(sup.Product),
			},
			":u": {
				S: aws.String(sup.Name),
			},
		},
		UpdateExpression: aws.String("SET #e = :e, #p = :p, #f = :f, #u = :u"),
		ExpressionAttributeNames: map[string]*string{
			"#e": aws.String("Email"),
			"#p": aws.String("PhoneNumber"),
			"#f": aws.String("Product"),
			"#u": aws.String("Name"),
		},
	}

	_, err = r.svc.UpdateItemWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return fmt.Errorf("failed to update item: %v - %v", aerr.Code(), aerr.Message())
		}
		return fmt.Errorf("failed to update item: %v", err)
	}

	return nil
}

func (r *supplierRepo) GetAllSuppliers(ctx context.Context) (*service.SupplierResult, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}

	result, err := r.svc.ScanWithContext(ctx, input) 
	if err != nil {
		return nil, fmt.Errorf("error scanning table: %v", err)
	}

	var suppliers []*service.Supplier
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &suppliers)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal DynamoDB result: %v", err)
	}

	supplyRes := &service.SupplierResult{
		Suppliers: suppliers,
	}

	return supplyRes, nil
}
