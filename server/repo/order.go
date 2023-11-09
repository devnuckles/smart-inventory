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

type OrderRepo interface {
	service.OrderRepo
}

type orderRepo struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

func NewOrderRepo(svc *dynamodb.DynamoDB, tableName string) OrderRepo {
	return &orderRepo{
		svc:       svc,
		tableName: tableName,
	}
}

func (r *orderRepo) CreateOrder(ctx context.Context, order *service.Order) error {
	newOrder, err := dynamodbattribute.MarshalMap(order)
	if err != nil {
		return fmt.Errorf("cannot marshal product: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      newOrder,
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
