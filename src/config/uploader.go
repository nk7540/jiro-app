package config

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploaderConfig struct {
	*s3manager.Uploader
	AWSRegion          string `mapstructure:"AWS_REGION" yaml:"aws_region" env:"AWS_REGION"`
	AWSAccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID" yaml:"aws_access_key_id" env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY" yaml:"aws_secret_access_key" env:"AWS_SECRET_ACCESS_KEY"`
	AWSS3Bucket        string `mapstructure:"AWS_S3_BUCKET" yaml:"aws_s3_bucket" env:"AWS_S3_BUCKET"`
	AWSS3Key           string `mapstructure:"AWS_S3_KEY" yaml:"aws_s3_key" env:"AWS_S3_KEY"`
}

func (c *UploaderConfig) Setup() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(c.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			c.AWSAccessKeyID,
			c.AWSSecretAccessKey,
			"",
		),
	}))

	c.Uploader = s3manager.NewUploader(sess)
}

func (c *UploaderConfig) Upload(file io.Reader, filename string) (*s3manager.UploadOutput, error) {
	return c.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.AWSS3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
}
