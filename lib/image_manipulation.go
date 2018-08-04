package lib

import bimg "gopkg.in/h2non/bimg.v1"

type ImageManipulation struct {
	bytesImage   []byte
	imageProcess ImageProcess
}

/**
 * Set bytes on image for processing the manipulation
 * @param bytes []byte
 * @return void
 */
func (ia *ImageManipulation) SetBytes(bytes []byte) {
	ia.bytesImage = bytes
}

/**
 * Set Image Processing for requirements data on manipulation image
 * @param imageProcess ImageProcess
 * @return void
 */
func (ia *ImageManipulation) SetImageProcess(imageProcess ImageProcess) {
	ia.imageProcess = imageProcess
}

/**
 * Processing image resize on original image bytes
 * @return []byte
 */
func (ia *ImageManipulation) Resize() ([]byte, error) {
	var imageResize []byte

	newImage, err := bimg.NewImage(ia.bytesImage).Resize(ia.imageProcess.Width, ia.imageProcess.Height)

	if err != nil {
		return nil, err
	}

	imageResize = newImage

	return imageResize, nil
}

/**
 * Processing Image thumbnail on original image bytes
 * @return []byte
 */
func (ia *ImageManipulation) Thumbnail() ([]byte, error) {
	var imageThumbnail []byte

	newImage, err := bimg.NewImage(ia.bytesImage).Thumbnail(ia.imageProcess.Width)

	if err != nil {
		return nil, err
	}

	imageThumbnail = newImage

	return imageThumbnail, nil
}

func NewImageManipulation() *ImageManipulation {
	var imageManipulation ImageManipulation
	return &imageManipulation
}
