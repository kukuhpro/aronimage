package lib

import (
	"aronimage/storage"
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

func (ai *Aronimage) setupStorage() {
	var configs3 storage.ConfigS3
	configs3.AccessKey = ai.Config.AWSKey
	configs3.AccessSecret = ai.Config.AWSSecret
	configs3.AwsRegion = ai.Config.AWSRegion
	configs3.BucketName = ai.Config.AWSBucket

	ai.Storage = storage.NewStorage(configs3)
}

func (ai *Aronimage) uploadOriginal() {
	ai.Storage.Put(ai.PrefixPath+"/"+ai.ModuleName+"/original", ai.Image.Name, ai.Image.Bytes)
}

func (ai *Aronimage) ProcessImage() {
	// upload first for original image
	ai.uploadOriginal()

	// processing other image send to queue job worker.
	payload := &ImageWorker{
		Storage:    ai.Storage,
		Path:       ai.Path,
		ModuleName: ai.ModuleName,
		ImageManipulation: NewImageManipulation()
	}
	work := qasirworker.Job{Executor: payload}
	qasirworker.JobQueue <- work
}

func NewAronImage() *Aronimage {
	var aron Aronimage

	aron.Path = NewImagePath()

	return &aron
}
