package converters

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/nfnt/resize"
)

var (
	UnknownImgTypeErr = errors.New("unknown image type")
)

const defaultWidth = 300

func GetImageTypeFromBytes(content []byte) (string, error) {
	bytesReader := bytes.NewReader(content)
	_, imgType, err := image.Decode(bytesReader)
	if err != nil {
		return "", err
	}
	return imgType, nil
}

func ResizeFromFile(file io.Reader, width uint) ([]byte, error) {
	img, imgType, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	if width == 0 {
		width = defaultWidth
	}
	imgResized := resize.Resize(width, 0, img, resize.Lanczos3)
	return handleImageType(imgType, imgResized)
}

func handleImageType(imgType string, img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	switch imgType {
	case "jpeg", "jpg":
		err := jpeg.Encode(buf, img, nil)
		return buf.Bytes(), err
	case "png":
		err := png.Encode(buf, img)
		return buf.Bytes(), err
	default:
		return nil, UnknownImgTypeErr
	}
}
