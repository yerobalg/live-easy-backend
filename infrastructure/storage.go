package infrastructure

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type Storage struct {
	BaseURL    string
	BucketName string
	FolderName string
	client     *storage.Client
}

func InitStorage(baseURL string, bucketName string, folderName string) *Storage {
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile("cloud_storage_credential.json"))
	if err != nil {
		panic(err)
	}

	return &Storage{
		BaseURL:    baseURL,
		BucketName: bucketName,
		FolderName: folderName,
		client:     client,
	}
}

func (s *Storage) GetObjectPlace(objectPath string) (*storage.ObjectHandle, error) {
	return s.client.Bucket(s.BucketName).Object(objectPath), nil
}
