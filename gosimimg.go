package gosimimg

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"

	"github.com/disintegration/gift"
)

// GetImage resizes to width, height and return a grayscaled Image
func GetImage(src image.Image, width, height int) image.Image {
	gi := gift.New(gift.Resize(width, height, gift.LinearResampling), gift.Grayscale())
	dst := image.NewRGBA(gi.Bounds(src.Bounds()))
	gi.Draw(dst, src)
	return dst
}

// GetHash returns the hash value of an Image
// A 64-bit hash value is calculated as "1" if it is larger than the average and "0" if it is smaller.
func GetHash(src image.Image) uint64 {
	srcBounds := src.Bounds()
	var (
		pixels    []uint64
		sumPixels uint64
		maxY      = srcBounds.Max.Y
		maxX      = srcBounds.Max.X
	)
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			r, g, b, _ := src.At(j, i).RGBA()
			// calculates the average value of r, g, b (Because grayscale is used, pixels: = uint64 (r) is also acceptable)
			pixel := uint64(math.Floor(float64((r + g + b)) / float64(3)))
			// Total value of pixel
			sumPixels += pixel

			pixels = append(pixels, pixel)
		}
	}
	var (
		hash uint64
		one  uint64 = 1

		// find the average of luminance values
		average = uint64(math.Floor(float64(sumPixels) / float64((maxY * maxX))))
	)
	for _, pixel := range pixels {
		// Check if pixel is greater than average
		if pixel > average {
			hash |= one
		}
		one = one << 1
	}
	return hash
}

// GetDistance returns the hamming distance between hashes
func GetDistance(hash1, hash2 uint64) int {
	distance := 0
	var i, k uint64
	for i = 0; i < 64; i++ {
		k = (1 << i)
		if (hash1 & k) != (hash2 & k) {
			distance++
		}
	}
	return distance
}

const threshold = 10

// IsSimilar returns whether the images are similar
func IsSimilar(src1, src2 image.Image) bool {
	hash1 := GetHash(src1)
	hash2 := GetHash(src2)
	distance := GetDistance(hash1, hash2)
	if distance < threshold {
		return true
	}
	return false
}
