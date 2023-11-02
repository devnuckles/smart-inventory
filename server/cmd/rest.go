package cmd

import (
	"github.com/Tonmoy404/Smart-Inventory/cache"
	"github.com/Tonmoy404/Smart-Inventory/config"
	"github.com/Tonmoy404/Smart-Inventory/repo"
	"github.com/Tonmoy404/Smart-Inventory/rest"
	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-redis/redis"
)

func serveRest() {
	appConfig := config.GetApp()
	awsConfig := config.GetAws()
	tableConfig := config.GetTable()
	saltConfig := config.GetSalt()
	tokenConfig := config.GetToken()
	s3Config := config.GetS3()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsConfig.Region),
	})
	if err != nil {
		panic(err)
	}

	ddbClient := dynamodb.New(sess)
	s3Client := s3.New(sess)
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	userRepo := repo.NewUserRepo(ddbClient, tableConfig.UserTableName)
	fileRepo := repo.NewFileRepo(s3Client, s3Config.Bucket)
	errorRepo := repo.NewErrorRepo(tableConfig.ErrorTableName, ddbClient)

	cache := cache.NewCache(redisClient)

	svc := service.NewService(userRepo, fileRepo, errorRepo, cache)
	server, err := rest.NewServer(appConfig, svc, saltConfig, tokenConfig)
	if err != nil {
		panic("Server can not start")
	}

	server.Start()
}
