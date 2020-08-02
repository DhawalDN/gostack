package main

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// CreateThumbnail - creating thumbnail for image
func main() {
	// func CreateThumbnail(uploadPath string, fileName string) {
	imagePath, _ := os.Open("./server/home/dhost/cdn/data/1596392259067.jpg")
	defer imagePath.Close()
	
	// Thumbnail function of Graphics
	newImage, _ := os.Create("thumbnail.jpg")
	defer newImage.Close()
	// jpeg.Encode(newImage, dstImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	jpeg.Encode(newImage, dstImage, nil)
}
