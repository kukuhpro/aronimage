package storage

import (
	"aronimage/storage/s3"
)

type Storage struct {
	driver           string
	config           ConfigS3
	prefixFolderPath string
	awsS3            s3.S3
}

func (st *Storage) SetConfig(configuration ConfigS3) {
	st.config = configuration
	st.awsS3.SettingAllConfiguration(&st.config)
}

func (st *Storage) Put(pathUpload string, filename string, bytes []byte) bool {
	pathUpload = st.prefixFolderPath + "/" + pathUpload + "/" + filename
	st.awsS3.Upload(pathUpload, bytes)
	return true
}

func (st *Storage) Get(path string) ([]byte, error) {
	return st.awsS3.Get(path)
}

func NewStorage(config ConfigS3) *Storage {
	var storages Storage
	storages.SetConfig(config)
	return &storages
}
