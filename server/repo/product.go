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

func (r *productRepo) UpdateProduct(ctx context.Context, product *service.Product) error {
	// Marshal the product into a DynamoDB attribute map
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(product.ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(product.Name),
			},
			":bp": {
				N: aws.String(fmt.Sprintf("%f", product.BuyingPrice)),
			},
			":q": {
				N: aws.String(fmt.Sprintf("%d", product.Quantity)),
			},
			"ed": {
				N: aws.String(fmt.Sprintf("%d", product.ExpiryDate)),
			},
		},
		UpdateExpression: aws.String("SET #n = :n, #d = :d, #p = :p, #q = :q"),
		ExpressionAttributeNames: map[string]*string{
			"#n": aws.String("Name"),
			"#d": aws.String("Description"),
			"#p": aws.String("Price"),
			"#q": aws.String("Quantity"),
		},
	}

	_, err := r.svc.UpdateItemWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return fmt.Errorf("failed to update item: %v - %v", aerr.Code(), aerr.Message())
		}
		return fmt.Errorf("failed to update item: %v", err)
	}

	return nil
}

func (r *productRepo) DeleteProductById(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
	}

	_, err := r.svc.DeleteItemWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return fmt.Errorf("failed to delete item: %v - %v", aerr.Code(), aerr.Message())
		}
		return fmt.Errorf("failed to delete item: %v", err)
	}

	return nil
}

func (r *productRepo) GetAllProducts(ctx context.Context) (*service.ProductsResult, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}

	result, err := r.svc.ScanWithContext(ctx, input) // Assuming you have the DynamoDB service client in r.svc
	if err != nil {
		return nil, fmt.Errorf("error scanning table: %v", err)
	}

	var products []*service.Product
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &products)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal DynamoDB result: %v", err)
	}

	productResult := &service.ProductsResult{
		Products: products,
	}

	return productResult, nil
}
