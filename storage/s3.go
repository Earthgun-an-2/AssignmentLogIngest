package storage

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	svc     *s3.S3
	bucket  string
	baseKey string
}

func NewS3Client() *S3Client {
	region := &aws.Config{Region: aws.String("ap-south-1")}
	sess := session.Must(session.NewSession(region))
	svc := s3.New(sess)

	return &S3Client{
		svc:     svc,
		bucket:  os.Getenv("S3_BUCKET"),
		baseKey: "logs/",
	}
}

func (c *S3Client) Store(data []byte) error {
	key := fmt.Sprintf("%slog-%d.json", c.baseKey, time.Now().UnixNano())
	_, err := c.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	return err
}
