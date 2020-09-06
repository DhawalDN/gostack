package helpers

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:27:40 pm
 * @copyright Crearosoft
 */

import (
	"image"
	jpeg "image/jpeg"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

type dimensions struct {
	nx int
	ny int
}

// CreateThumbnail - creating thumbnail for image
func CreateThumbnail(filePath, fileName string) {
	thumbnailName := fileName + "_th"
	thumbnailFilePath := strings.Replace(filePath, fileName, thumbnailName, 1)
	imagePath, _ := os.Open(filePath)
	defer imagePath.Close()
	srcImage, _, _ := image.Decode(imagePath)
	width := srcImage.Bounds().Dx()
	height := srcImage.Bounds().Dy()
	// Dimension of new thumbnail 80 X 80
	dstImage := resizer(width, height, srcImage)
	newImage, _ := os.Create(thumbnailFilePath)
	defer newImage.Close()
	jpeg.Encode(newImage, dstImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
}

func resizer(width, height int, srcImage image.Image) image.Image {
	if width > 200 || height > 200 {
		return resize.Resize(200, 0, srcImage, resize.Lanczos3)
	}
	return resize.Resize(uint(width), 0, srcImage, resize.Lanczos3)
}
