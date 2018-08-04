package lib

type ImageManipulation struct {
	bytesImage []byte
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
 * Processing image resize on original image bytes
 * @return []byte
 */
func (ia *ImageManipulation) Resize() []byte {
	var imageResize []byte

	return imageResize
}

/**
 * Processing Image thumbnail on original image bytes
 * @return []byte
 */
func (ia *ImageManipulation) Thumbnail() []byte {
	var imageThumbnail []byte

	return imageThumbnail
}

func NewImageManipulation() *ImageManipulation {
	var imageManipulation ImageManipulation
	return &imageManipulation
}
