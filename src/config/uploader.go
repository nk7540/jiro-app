package config

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploaderConfig struct {
	*s3manager.Uploader
	AWSProfile  string `mapstructure:"AWS_PROFILE" yaml:"profile" env:"AWS_PROFILE"`
	AWSS3Bucket string `mapstructure:"AWS_S3_BUCKET" yaml:"s3_bucket" env:"AWS_S3_BUCKET"`
	AWSS3Key    string `mapstructure:"AWS_S3_KEY" yaml:"s3_key" env:"AWS_S3_KEY"`
}

func (c *UploaderConfig) Setup() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           c.AWSProfile,
		SharedConfigState: session.SharedConfigEnable,
	}))

	c.Uploader = s3manager.NewUploader(sess)
}

func (c *UploaderConfig) Upload(file io.Reader) (*s3manager.UploadOutput, error) {
	return c.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.AWSS3Bucket),
		Key:    aws.String(c.AWSS3Key),
		Body:   file,
	})
}
