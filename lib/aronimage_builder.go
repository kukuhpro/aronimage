package lib

import "aronimage/storage"

type AronImageBuilder struct {
	Base64     string
	FileName   string
	Extension  string
	PrefixPath string
}

/**
 * Setup with base64 string
 * @param base64 string
 * @return *AronImageBuilder
 */
func (ab *AronImageBuilder) WithBase64(base64 string) *AronImageBuilder {
	ab.Base64 = base64
	return ab
}

/**
 * Setup with filename image
 * @param fileName string
 * @return *AronImageBuilder
 */
func (ab *AronImageBuilder) WithFileName(fileName string) *AronImageBuilder {
	ab.FileName = fileName
	return ab
}

/**
 * Setup with image file extension
 * @param extension string
 * @return *AronImageBuilder
 */
func (ab *AronImageBuilder) WithExtension(extension string) *AronImageBuilder {
	ab.Extension = extension
	return ab
}

/**
 * Setup configuration for AWS S3
 * @return Config
 */
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
	aronImage.SetupStorage()

	return aronImage
}
