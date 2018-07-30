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
	Name         string         `json:"name"`
	ImageProcess []ImageProcess `json:"image_process"`
}

type ImageProcess struct {
	Type       string `json:"type"`
	PrefixPath string `json:"prefix_path"`
	Height     int    `json:"height"`
	Width      int    `json:"width"`
}

type ImageLabel struct {
	Name      string
	Extension string
	Bytes     []byte
}
