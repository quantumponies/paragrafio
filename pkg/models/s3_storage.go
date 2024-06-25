package models

type S3Storage interface {
	GetBucket(name string) (string, error)
	Upload(bucket string, filename string, arr []byte, mime string) (string, error)
}
