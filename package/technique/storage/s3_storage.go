package storage

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

type S3FileStorage struct {
	bucket string
	client *s3.Client
}

func NewS3FileStorage(client *s3.Client, bucket string) (*S3FileStorage, error) {
	return &S3FileStorage{
		bucket: bucket,
		client: client,
	}, nil
}

func (s *S3FileStorage) SaveFile(filename string, data []byte) error {
	// Upload the data to S3
	uploader := manager.NewUploader(s.client)
	_, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &filename,
		Body:   bytes.NewReader(data),
	})
	return err
}

func (s *S3FileStorage) GetFile(filename string) ([]byte, error) {
	result, err := s.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	return io.ReadAll(result.Body)

}

func (s *S3FileStorage) DeleteFile(filename string) error {
	// Delete the object from S3
	_, err := s.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    &filename,
	})

	return err
}
