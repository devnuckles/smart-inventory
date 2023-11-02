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

const (
	userIdxIndex         = "IdxIndex"
	userEmailIndex       = "EmailIndex"
	userPhoneNumberIndex = "PhoneNumberIndex"
	userUsernameIndex    = "UsernameIndex"
	userFullnameIndex    = "FullnameIndex"
	userRoleIndex        = "RoleIndex"
)

type UserRepo interface {
	service.UserRepo
}

type userRepo struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

func NewUserRepo(svc *dynamodb.DynamoDB, tableName string) UserRepo {
	return &userRepo{
		svc:       svc,
		tableName: tableName,
	}
}

func (r *userRepo) Create(ctx context.Context, user *service.User) error {

	usr, err := dynamodbattribute.MarshalMap(user)
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

func (r *userRepo) GetItemByEmail(ctx context.Context, email string) (*service.User, error) {
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

	var user *service.User
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB item: %v", err)
	}
	return user, nil
}
