package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

func Init() *s3.S3 {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			viper.GetString("s3.access_key_id"),
			viper.GetString("s3.secret_access_key"),
			"",
		),
		Endpoint:         aws.String(viper.GetString("s3.endpoint")),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		panic(err)
	}
	return s3.New(newSession)
}
