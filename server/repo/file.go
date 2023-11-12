package repo

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

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

func (r *fileRepo) Upload(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	fileName, err := createFileName(fileHeader)
	if err != nil {
		return "", err
	}

	_, err = r.s3Client.PutObject((&s3.PutObjectInput{
		Bucket: aws.String(r.s3Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	}))
	if err != nil {
		return "", fmt.Errorf("cannot put file into s3: %v", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", r.s3Bucket, fileName)

	return fileURL, nil
}

func createFileName(fileHeader *multipart.FileHeader) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("cannot create uuid: %v", err)
	}

	fileName := fmt.Sprintf("%s%s", id, filepath.Ext(fileHeader.Filename))

	return fileName, nil
}
