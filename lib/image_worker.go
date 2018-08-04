package lib

import (
	"aronimage/storage"
)

type ImageWorker struct {
	ModuleName        string
	Storage           storage.StorageInterface
	Image             ImageLabel
	Path              ImagePathInterface
	ImageManipulation *ImageManipulation
}

func (iw *ImageWorker) processingImage(image ImageProcess) error {
	// set bytes original image
	iw.ImageManipulation.SetBytes(iw.Image.Bytes)
	switch image.Type {
	case "resize":

	case "thumbnail":

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
