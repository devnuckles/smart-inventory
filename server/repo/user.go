package repo

import (
	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
