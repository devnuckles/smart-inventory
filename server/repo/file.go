package repo

import "github.com/aws/aws-sdk-go/service/s3"

type FileRepo interface {
	service.FileRepo
}

type fileRepo struct {
	s3Client *s3.S3
	s3Bucket string
}

func NewFileRepo(s3Client *s3.S3, s3Bucket string) FileRepo {
	return &fileRepo{
		s3Client: s3Client,
		s3Bucket: s3Bucket,
	}
}
