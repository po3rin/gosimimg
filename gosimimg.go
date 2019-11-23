package gosimimg

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"

	"github.com/disintegration/gift"
)

// Similar has setting for similar image search using averge hash.
type Similar struct {
	Threshold        int
	CompressedWidth  int
	CompressedHeight int
}

// IsSimilar returns whether it is a similar image
func (s *Similar) IsSimilar(img1, img2 image.Image) bool {
	src1 := GetImage(img1, s.CompressedWidth, s.CompressedHeight)
	src2 := GetImage(img2, s.CompressedWidth, s.CompressedHeight)

	hash1 := GetHash(src1)
	hash2 := GetHash(src2)

	d := GetDistance(hash1, hash2)
	return d < s.Threshold
}

// SetThreshold sets threshold for Determining if it is a similar image.
// Recommendations should be made between 0 and 10.
// this should be determined according to the nature of each service.
func SetThreshold(threshold int) func(*Similar) {
	return func(s *Similar) {
		s.Threshold = threshold
	}
}

// SetCompressedWidth sets compressedWidth.
// Recommendations should be made between 8 or 16.
func SetCompressedWidth(compressedWidth int) func(*Similar) {
	return func(s *Similar) {
		s.CompressedWidth = compressedWidth
	}
}

// SetCompressedHeight sets compressedHeight.
// Recommendations should be made between 8 or 16.
func SetCompressedHeight(compressedHeight int) func(*Similar) {
	return func(s *Similar) {
		s.CompressedHeight = compressedHeight
	}
}

// GetImage resizes to width, height and return a grayscaled image.
func GetImage(src image.Image, width, height int) image.Image {
	gi := gift.New(gift.Resize(width, height, gift.LinearResampling), gift.Grayscale())
	dst := image.NewRGBA(gi.Bounds(src.Bounds()))
	gi.Draw(dst, src)
	return dst
}

// NewSimilar inits similar with option func.
func NewSimilar(options ...func(s *Similar)) *Similar {
	// sets default value.
	s := &Similar{
		Threshold:        10,
		CompressedWidth:  8,
		CompressedHeight: 8,
	}
	for _, option := range options {
		option(s)
	}
	return s
}

// GetHash returns the hash value of an Image
// A 64-bit hash value is calculated as "1" if it is larger than the average and "0" if it is smaller.
func GetHash(src image.Image) uint64 {
	b := src.Bounds()
	var (
		pixels    []uint64
		sumPixels uint64
	)
	for i := 0; i < b.Max.Y; i++ {
		for j := 0; j < b.Max.X; j++ {
			r, g, b, _ := src.At(j, i).RGBA()
			// calculates the average value of r, g, b (Because grayscale is used, pixels: = uint64 (r) is also acceptable) .
			pixel := uint64(math.Floor(float64((r + g + b)) / float64(3)))
			pixels = append(pixels, pixel)
			// for calcurate average.
			sumPixels += pixel
		}
	}

	var (
		hash uint64
		one  uint64 = 1
	)

	ave := uint64(math.Floor(float64(sumPixels) / float64((b.Max.Y * b.Max.X))))
	for _, pixel := range pixels {
		if pixel > ave {
			hash |= one
		}
		one = one << 1
	}
	return hash
}

// GetDistance returns the hamming distance between hashes.
func GetDistance(hash1, hash2 uint64) int {
	var d int
	var i, k uint64
	for i = 0; i < 64; i++ {
		k = (1 << i)
		if (hash1 & k) != (hash2 & k) {
			d++
		}
	}
	return d
}
