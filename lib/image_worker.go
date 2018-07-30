package lib

import (
	"aronimage/storage"
)

type ImageWorker struct {
	ModuleName string
	Storage    storage.StorageInterface
	Image      ImageLabel
	Path       ImagePathInterface
}

func (iw *ImageWorker) processingImage(image ImageProcess) error {
	switch image.Type {
	case "resize":

	}
}

func (iw *ImageWorker) Handle() error {
	path := iw.Path.GetPath(iw.ModuleName)
	for _, val := range path.ImageProcess {
		err := iw.processingImage(val)
		if err != nil {
			return err
		}
	}
	return nil
}
