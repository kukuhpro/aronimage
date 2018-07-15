package core

import (
	"core/storages"
	"fmt"
)

type storage struct {
	driver           string
	config           map[string]interface{}
	prefixFolderPath string
	awsS3            storages.S3
}

func (this *storage) SetDriver(driver string) {
	this.driver = driver
	if this.driver == "s3" {
		var s3 storages.S3
		s3.SettingAllConfiguration(this.config)
		this.awsS3 = s3
	}
}

func (this *storage) SetConfig(configuration map[string]interface{}) {
	this.config = configuration
	this.prefixFolderPath = fmt.Sprint(configuration["prefixFolderPath"])
	this.SetDriver(fmt.Sprint(configuration["driver"]))
}

func (this *storage) Put(pathUpload string, filename string, bytes []byte) bool {
	if this.driver == "s3" {
		pathUpload = this.prefixFolderPath + "/" + pathUpload + "/" + filename
		this.awsS3.Upload(pathUpload, bytes)
	}
	return true
}

func (this *storage) Get(path string) ([]byte, error) {
	return this.awsS3.Get(path)
}

func Storage(config map[string]interface{}) *storage {
	var storages storage
	storages.SetConfig(config)
	return &storages
}
