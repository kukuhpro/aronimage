package lib

import "aronimage/storage"

type ImageWorker struct {
	ModuleName string
	Storage    storage.StorageInterface
	Image      ImageLabel
	Path       ImagePathInterface
}

func (iw *ImageWorker) Handle() error {
	return nil
}
