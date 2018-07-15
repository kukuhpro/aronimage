package storage

type StorageInterface interface {
	Put(pathUpload string, filename string, bytes []byte) bool
	Get(path string) ([]byte, error)
}

type ConfigS3 struct {
	AccessKey    string
	AccessSecret string
	BucketName   string
	AwsRegion    string
}

func (ss *ConfigS3) GetAccessKey() string {
	return ss.AccessKey
}
func (ss *ConfigS3) GetAccessSecret() string {
	return ss.AccessSecret
}
func (ss *ConfigS3) GetRegion() string {
	return ss.AwsRegion
}
func (ss *ConfigS3) GetS3Bucket() string {
	return ss.AwsRegion
}
