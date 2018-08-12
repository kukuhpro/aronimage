package lib

import (
	"aronimage/storage"
	"errors"
	"qasirworker"
)

type Aronimage struct {
	ModuleName string
	PathName   string
	PrefixPath string
	Config     Config
	Storage    storage.StorageInterface
	Image      ImageLabel
	Path       ImagePathInterface
}

/**
 * Setup Storage connection to AWS S3
 * @return void
 */
func (ai *Aronimage) SetupStorage() {
	var configs3 storage.ConfigS3
	configs3.AccessKey = ai.Config.AWSKey
	configs3.AccessSecret = ai.Config.AWSSecret
	configs3.AwsRegion = ai.Config.AWSRegion
	configs3.BucketName = ai.Config.AWSBucket

	ai.Storage = storage.NewStorage(configs3)
}

/**
 * Upload Original Image on AWS S3 first
 * @return void
 */
func (ai *Aronimage) UploadOriginal() error {
	flag := ai.Storage.Put(ai.PrefixPath+"/"+ai.ModuleName+"/original", ai.Image.Name, ai.Image.Bytes)

	if !flag {
		return errors.New("Failed Upload original image.")
	}
	return nil
}

/**
 * Processing image for handle manipulation
 * @return void
 */
func (ai *Aronimage) ProcessImage() error {
	// upload first for original image
	err := ai.UploadOriginal()

	if err != nil {
		return err
	}

	// processing other image send to queue job worker.
	payload := &ImageWorker{
		PrefixPath:        ai.PrefixPath,
		Storage:           ai.Storage,
		Path:              ai.Path,
		ModuleName:        ai.ModuleName,
		ImageManipulation: NewImageManipulation(),
		Image:             ai.Image,
	}
	work := qasirworker.Job{Executor: payload}
	qasirworker.JobQueue <- work

	return nil
}

func NewAronImage() *Aronimage {
	var aron Aronimage

	aron.Path = NewImagePath()

	return &aron
}
