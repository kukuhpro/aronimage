package lib

import (
	"aronimage/storage"
	"log"
)

type ImageWorker struct {
	PrefixPath        string
	ModuleName        string
	Storage           storage.StorageInterface
	Image             ImageLabel
	Path              ImagePathInterface
	ImageManipulation *ImageManipulation
}

/**
 * Processing image handle on worker
 * @param image ImageProcess
 * @return error
 */
func (iw *ImageWorker) processingImage(image ImageProcess) error {
	// set bytes original image
	iw.ImageManipulation.SetBytes(iw.Image.Bytes)

	// set Image Processing variable
	iw.ImageManipulation.SetImageProcess(image)

	var imagesBytes []byte
	var err error

	switch image.Type {
	case "resize":
		imagesBytes, err = iw.ImageManipulation.Resize()
	case "thumbnail":
		imagesBytes, err = iw.ImageManipulation.Thumbnail()
	}

	prefixUploadPath := image.generatePrefixPath(iw.PrefixPath, iw.ModuleName)
	if err == nil {
		// put binnary result image processing to AWS S3
		iw.Storage.Put(prefixUploadPath, iw.Image.Name, imagesBytes)
		return err
	} else {
		log.Println(err)
		return err
	}
}

/**
 * Interface for qasir worker
 * @param void
 * @return error
 */
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
