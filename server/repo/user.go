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

func (r *userRepo) GetItemByID(ctx context.Context, id string) (*service.User, error) {
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

	var user *service.User
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB item: %v", err)
	}

	return user, nil
}

func (r *userRepo) DeleteItemByID(ctx context.Context, id string) error {
	u, err := r.GetItemByID(ctx, id)
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

func (r *userRepo) UpdateItemByID(ctx context.Context, user *service.User) error {
	u, err := r.GetItemByID(ctx, user.ID)
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
				S: aws.String(user.ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":e": {
				S: aws.String(user.Email),
			},
			":p": {
				S: aws.String(user.Password),
			},
			":f": {
				S: aws.String(user.Fullname),
			},
			":n": {
				S: aws.String(user.PhoneNumber),
			},
			":u": {
				S: aws.String(user.Username),
			},
			":r": {
				S: aws.String(user.Role),
			},
			":sts": {
				S: aws.String(user.Status),
			},
		},
		UpdateExpression: aws.String("SET #e = :e, #p = :p, #f = :f, #n = :n, #u = :u, #r = :r, #sts = :sts, #pc = :pc"),
		ExpressionAttributeNames: map[string]*string{
			"#e":   aws.String("Email"),
			"#p":   aws.String("Password"),
			"#f":   aws.String("Fullname"),
			"#n":   aws.String("PhoneNumber"),
			"#u":   aws.String("Username"),
			"#r":   aws.String("Role"),
			"#sts": aws.String("Status"),
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
