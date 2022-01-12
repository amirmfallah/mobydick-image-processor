package lib

import (
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	uuid "github.com/google/uuid"
)

const PRE_SIGNED_URI_EXPIRY = 10

var (
	onceS3Downloader sync.Once
	S3Downloader     *s3manager.Downloader
	S3Uploader       *s3manager.Uploader
	onceS3Instance   sync.Once
	onceS3Uploader   sync.Once
	S3Instance       *s3.S3
)

func GetS3Instance() (*s3.S3, error) {
	var err error = nil
	onceS3Instance.Do(func() {
		awsConf := &aws.Config{
			Region: aws.String("default"),
			Credentials: credentials.NewStaticCredentials(
				S3_ACCESS_KEY,
				S3_SECRET_KEY,
				"", // a token will be created when the session it's used.
			),
			Endpoint: aws.String(S3_ENDPOINT),
		}
		session, errSession := session.NewSession(awsConf)
		if errSession != nil {
			log.Println("GetS3Instance error:", errSession)
			err = errSession
			return
		}
		S3Instance = s3.New(session)
	})
	return S3Instance, err
}

func GetS3Uploader() (*s3manager.Uploader, error) {
	var err error = nil
	onceS3Uploader.Do(func() {
		awsConf := &aws.Config{
			Region: aws.String("default"),
			Credentials: credentials.NewStaticCredentials(
				S3_ACCESS_KEY,
				S3_SECRET_KEY,
				"", // a token will be created when the session it's used.
			),
			Endpoint: aws.String(S3_ENDPOINT),
		}
		session, errSession := session.NewSession(awsConf)
		if errSession != nil {
			log.Println("GetS3Downloader error:", errSession)
			err = errSession
			return
		}
		S3Uploader = s3manager.NewUploader(session)
	})
	return S3Uploader, err
}

func UploadImage(user string, types string, file io.Reader) (*UploadedFile, error) {
	uploadedFile := new(UploadedFile)
	newUuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Print("UploadImage-Error creating a new UUID:", err.Error())
		return nil, err
	}
	fileId := newUuid.String()
	objectKey := strings.Join([]string{"/users", user, types, fileId + ".jpg"}, "/")
	s3Uploader, err := GetS3Uploader()
	if err != nil {
		fmt.Print("UploadImage-Error GetS3Uploader:", err.Error())
		return nil, err
	}

	up, err := s3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3_BUCKET_NAME),
		ACL:    aws.String("public-read"),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		fmt.Print("UploadImage-Error s3Uploader.Upload", err.Error())
		return nil, err
	}
	uploadedFile.Key = objectKey
	uploadedFile.Location = up.Location

	return uploadedFile, nil
}
