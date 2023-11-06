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

type ProductRepo interface {
	service.ProductRepo
}

type productRepo struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

func NewProductRepo(svc *dynamodb.DynamoDB, tableName string) ProductRepo {
	return &productRepo{
		svc:       svc,
		tableName: tableName,
	}
}

func (r *productRepo) CreateProduct(ctx context.Context, product *service.Product) (*service.Product, error) {
	prod, err := dynamodbattribute.MarshalMap(product)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal product: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      prod,
	}
	_, err = r.svc.PutItemWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return nil, fmt.Errorf("failed to write item: %v - %v", aerr.Code(), aerr.Message())
		}
		return nil, fmt.Errorf("failed to write item: %v", err)
	}

	return product, nil
}
