package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/po3rin/gosimimg"
)

func main() {
	f1, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	img1, err := jpeg.Decode(f1)
	if err != nil {
		log.Fatal(err)
	}
	img2, err := jpeg.Decode(f2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success to load images")

	s := gosimimg.NewSimilar(
		gosimimg.SetThreshold(10),
		gosimimg.SetCompressedWidth(8),
		gosimimg.SetCompressedHeight(8),
	)
	if !s.IsSimilar(img1, img2) {
		fmt.Println("not simmilar !!")
		return
	}
	fmt.Println("simmilar !!")
}
