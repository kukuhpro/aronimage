package lib

import "aronimage/storage"

type AronImageBuilder struct {
	Base64     string
	FileName   string
	Extension  string
	PrefixPath string
}

func (ab *AronImageBuilder) WithBase64(base64 string) *AronImageBuilder {
	ab.Base64 = base64
	return ab
}

func (ab *AronImageBuilder) WithFileName(fileName string) *AronImageBuilder {
	ab.FileName = fileName
	return ab
}

func (ab *AronImageBuilder) WithExtension(extension string) *AronImageBuilder {
	ab.Extension = extension
	return ab
}

func (ab *AronImageBuilder) setupConfig() Config {
	var config Config

	config.AWSBucket = Env.Get("AWS_BUCKET")
	config.AWSKey = Env.Get("AWS_KEY")
	config.AWSSecret = Env.Get("AWS_SECRET")
	config.AWSRegion = Env.Get("AWS_REGION")

	return config
}

func (ab *AronImageBuilder) Build(modulename string) *Aronimage {
	var imageLabel ImageLabel

	imageLabel.Extension = ab.Extension
	imageLabel.Name = ab.FileName
	imageLabel.Bytes = storage.Base64ToByte(ab.Base64)

	aronImage := NewAronImage()
	aronImage.Image = imageLabel
	aronImage.PrefixPath = ab.PrefixPath
	aronImage.Config = ab.setupConfig()
	aronImage.ModuleName = modulename
	aronImage.setupStorage()

	return aronImage
}
