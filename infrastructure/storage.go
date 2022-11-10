package infrastructure

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type Storage struct {
	BucketName string
	client     *storage.Client
}

func (s *Storage) Init(bucketName string) *Storage {
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile("credentials.json"))
	if err != nil {
		panic(err)
	}

	return &Storage{
		BucketName: bucketName,
		client:     client,
	}
}

func (s *Storage) GetObjectPlace(objectPath string) (*storage.ObjectHandle, error) {
	return s.client.Bucket(s.BucketName).Object(objectPath), nil
}
