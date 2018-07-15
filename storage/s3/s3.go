package s3

import (
	"fmt"
	"net/http"

	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

type ConfigS3interface interface {
	GetAccessKey() string
	GetAccessSecret() string
	GetRegion() string
	GetS3Bucket() string
}

type S3 struct {
	accessKey  string
	secretKey  string
	bucketName string
	awsregion  string
}

func (s *S3) SettingAllConfiguration(config ConfigS3interface) {
	s.SetAccessKey(config.GetAccessKey())
	s.SetSecretKey(config.GetAccessSecret())
	s.SetAwsRegion(config.GetRegion())
	s.SetBucketName(config.GetS3Bucket())
}

func (s *S3) SetAccessKey(accessKey string) {
	s.accessKey = accessKey
}

func (s *S3) SetAwsRegion(awsRegion string) {
	s.awsregion = awsRegion
}

func (s *S3) SetSecretKey(secretkey string) {
	s.secretKey = secretkey
}

func (s *S3) SetBucketName(bucketName string) {
	s.bucketName = bucketName
}

func (s *S3) Get(path string) ([]byte, error) {
	AWSAuth := aws.Auth{
		AccessKey: s.accessKey, // change s to yours
		SecretKey: s.secretKey,
	}

	AWSRegion := aws.Regions[s.awsregion]

	connection := s3.New(AWSAuth, AWSRegion)

	bucket := connection.Bucket(s.bucketName)

	return bucket.Get(path)
}

func (s *S3) Upload(pathUpload string, bytes []byte) {
	AWSAuth := aws.Auth{
		AccessKey: s.accessKey, // change s to yours
		SecretKey: s.secretKey,
	}

	AWSRegion := aws.Regions[s.awsregion]

	connection := s3.New(AWSAuth, AWSRegion)

	bucket := connection.Bucket(s.bucketName)
	fileType := http.DetectContentType(bytes)
	err := bucket.Put(pathUpload, bytes, fileType, s3.ACL("public-read"))

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
