package s3

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage struct {
	S3 *s3.S3
}

func (service *S3Storage) GetBucket(name string) (string, error) {
	bucket := aws.String(name)
	cparams := &s3.CreateBucketInput{
		Bucket: bucket,
	}

	_, err := service.S3.CreateBucket(cparams)
	return *bucket, err
}

func (service *S3Storage) Upload(bucket string, filename string, arr []byte, mime string) (string, error) {
	_, err := service.S3.PutObject(&s3.PutObjectInput{
		Body:        bytes.NewReader(arr),
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		ContentType: aws.String(mime),
		Bucket:      &bucket,
		Key:         &filename,
	})

	return filename, err
}
