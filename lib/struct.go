package lib

import "core"

var (
	Env = core.Env("./")
)

type ImagePathInterface interface {
	GetPath(key string) Path
}

type Config struct {
	AWSKey    string
	AWSSecret string
	AWSBucket string
	AWSRegion string
}

type Path struct {
	Name         string
	ImageProcess []ImageProcess
}

type ImageProcess struct {
	Type       string
	PrefixPath string
	Height     string
	Width      string
}

type ImageLabel struct {
	Name      string
	Extension string
	Bytes     []byte
}
