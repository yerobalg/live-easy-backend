package infrastructure

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type Storage struct {
	BucketName string
}

func (s *Storage) Init(bucketName string) *Storage {
	return &Storage{
		BucketName: bucketName,
	}
}

func (s *Storage) GetObjectPlace(ctx *gin.Context, objectPath string) (*storage.ObjectHandle, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("cloud_storage_credential.json"))
	if err != nil {
		return nil, err
	}

	return client.Bucket(s.BucketName).Object(objectPath), nil
}
