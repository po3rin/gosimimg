package gosimimg_test

import (
	"image/jpeg"
	"os"
	"testing"

	"github.com/po3rin/gosimimg"
)

func TestIsSimilar(t *testing.T) {
	tests := []struct {
		name         string
		imgFilePath1 string
		imgFilePath2 string
		x            int
		y            int
		similar      bool
	}{
		{
			name:         "similar",
			imgFilePath1: "testdata/same1.jpg",
			imgFilePath2: "testdata/same2.jpg",
			x:            8,
			y:            8,
			similar:      true,
		},
		{
			name:         "similar",
			imgFilePath1: "testdata/sim1_1.jpg",
			imgFilePath2: "testdata/sim1_2.jpg",
			x:            8,
			y:            8,
			similar:      true,
		},
		{
			name:         "similar",
			imgFilePath1: "testdata/sim2_1.jpg",
			imgFilePath2: "testdata/sim2_2.jpg",
			x:            8,
			y:            8,
			similar:      true,
		},
		{
			name:         "similar",
			imgFilePath1: "testdata/sim3_1.jpg",
			imgFilePath2: "testdata/sim3_2.jpg",
			x:            8,
			y:            8,
			similar:      true,
		},
		{
			name:         "similar",
			imgFilePath1: "testdata/sim1_1.jpg",
			imgFilePath2: "testdata/sim2_2.jpg",
			x:            8,
			y:            8,
			similar:      false,
		},
		{
			name:         "similar",
			imgFilePath1: "testdata/sim2_1.jpg",
			imgFilePath2: "testdata/sim1_2.jpg",
			x:            8,
			y:            8,
			similar:      false,
		},
	}
	for _, tt := range tests {
		f1, err := os.Open(tt.imgFilePath1)
		if err != nil {
			t.Fatal(err)
		}
		defer f1.Close()

		f2, err := os.Open(tt.imgFilePath2)
		if err != nil {
			t.Fatal(err)
		}
		defer f2.Close()

		img1, err := jpeg.Decode(f1)
		if err != nil {
			t.Fatal(err)
		}
		img2, err := jpeg.Decode(f2)
		if err != nil {
			t.Fatal(err)
		}

		src1 := gosimimg.GetImage(img1, tt.x, tt.y)
		src2 := gosimimg.GetImage(img2, tt.x, tt.y)

		similar := gosimimg.IsSimilar(src1, src2)

		if tt.similar && !similar {
			t.Error("unexpected similar, but got no similar")
		}

		if !tt.similar && similar {
			t.Error("unexpected no similar, but got imilar")
		}
	}
}
