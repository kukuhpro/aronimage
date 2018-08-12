package tests

import (
	"aronimage/lib"
	"aronimage/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AronImageTestSuite struct {
	aronImage *lib.Aronimage
}

var aronImageTestSuite AronImageTestSuite

func init() {
	aronImageTestSuite.aronImage = &lib.Aronimage{}
}

func TestConstructorFunction(t *testing.T) {
	aronImage := lib.NewAronImage()

	assert.IsType(t, aronImage, aronImageTestSuite.aronImage)
}

func TestSetupStorage(t *testing.T) {
	var config lib.Config

	config.AWSBucket = "xxx"
	config.AWSKey = "xxx"
	config.AWSRegion = "xxx"
	config.AWSSecret = "xxx"

	aronImageTestSuite.aronImage.Config = config

	aronImageTestSuite.aronImage.SetupStorage()

	var storage storage.StorageInterface

	assert.IsType(t, aronImageTestSuite.aronImage.Storage, storage)
}
