package aws

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Uploader struct {
	Uploader *s3manager.Uploader
	Bucket   string
	Key      string
}

func NewUploader(sess *session.Session, bucket string, key string) *Uploader {
	uploader := s3manager.NewUploader(sess)
	return &Uploader{
		Uploader: uploader,
		Bucket:   bucket,
		Key:      key,
	}
}

func (u *Uploader) Upload(file io.Reader) error {
	_, err := u.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(u.Key),
		Body:   file,
	})

	return err
}
