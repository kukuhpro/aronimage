package tests

import (
	"aronimage/storage"
)

type ConfigS3Suite struct {
	configS3 *storage.ConfigS3
}

var configS3Suite ConfigS3Suite

func init() {
	configS3Suite.configS3 = &storage.ConfigS3{}
}
